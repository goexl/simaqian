package loki

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/simaqian/internal/config"
	"github.com/goexl/simaqian/internal/internal/internal"
	"github.com/goexl/simaqian/internal/internal/loki/internal/api"
	"github.com/goexl/simaqian/internal/internal/loki/internal/key"
	"go.uber.org/zap"
)

type Pusher struct {
	ctx   context.Context
	quit  chan gox.Empty
	logs  chan *internal.Log
	group sync.WaitGroup

	http   *http.Client
	batch  *config.Batch
	labels gox.Labels
	url    string
	logger *zap.Logger
}

func New(ctx context.Context, config *Config) (pusher *Pusher) {
	pusher = new(Pusher)
	pusher.ctx = ctx
	pusher.http = config.Http
	pusher.labels = config.Labels
	pusher.batch = config.Batch
	pusher.quit = make(chan gox.Empty)
	pusher.logs = make(chan *internal.Log)

	pusher.url = fmt.Sprintf("%s/loki/api/v1/push", config.Url)
	pusher.http.SetBasicAuth(config.Username, config.Password).SetHeader(key.ContentType, key.Json)
	pusher.group.Add(1)
	go pusher.run()

	return
}

func (p *Pusher) sink(_ *url.URL) (zap.Sink, error) {
	return internal.NewSink(p), nil
}

func (p *Pusher) Stop() {
	close(p.quit)
	p.group.Wait()
}

func (p *Pusher) Push(log *internal.Log) {
	p.logs <- log
}

func (p *Pusher) Build(config zap.Config, opts ...zap.Option) (logger *zap.Logger, err error) {
	if rse := zap.RegisterSink(key.KeyLokiSink, p.sink); nil != rse {
		err = rse
	} else {
		_key := fmt.Sprintf("%s://", key.KeyLokiSink)
		config.OutputPaths = gox.Ift(nil == config.OutputPaths, []string{_key}, append(config.OutputPaths, _key))
		p.logger, err = config.Build(opts...)
		logger = p.logger
	}

	return
}

func (p *Pusher) run() {
	logs := make([]*internal.Log, 0, p.batch.Size)
	ticker := time.NewTimer(p.batch.Wait)
	defer p.cleanup(&logs)

	for {
		select {
		case <-p.ctx.Done():
			break
		case <-p.quit:
			break
		case log := <-p.logs:
			logs = append(logs, log)
			if len(logs) >= p.batch.Size {
				_ = p.send(&logs)
				logs = make([]*internal.Log, 0)
				ticker.Reset(p.batch.Wait)
			}
		case <-ticker.C:
			if len(logs) > 0 {
				_ = p.send(&logs)
				logs = make([]*internal.Log, 0)
			}
			ticker.Reset(p.batch.Wait)
		}
	}
}

func (p *Pusher) send(logs *[]*internal.Log) (err error) {
	if request, ae := p.make(logs); nil != ae {
		err = ae
	} else if data, me := json.Marshal(request); nil != me {
		err = me
	} else {
		err = p.post(data)
	}
	if nil != err {
		p.logger.Warn("推送日志出错", zap.Error(err))
	}

	return
}

func (p *Pusher) make(logs *[]*internal.Log) (request *api.Request, err error) {
	request = new(api.Request)
	values := make([]*api.Value, 0, len(*logs))
	for _, log := range *logs {
		values = append(values, &api.Value{
			strconv.FormatInt(log.Timestamp.UnixNano(), 10),
			log.Raw(),
		})
	}
	request.Streams = append(request.Streams, &api.Stream{
		Stream: p.labels,
		Values: values,
	})

	return
}

func (p *Pusher) post(data []byte) (err error) {
	request := p.http.R()
	if buffer, ge := p.gzip(request, data); nil != ge {
		err = ge
	} else if rsp, pe := request.SetBody(buffer).Post(p.url); nil != pe {
		err = pe
	} else if rsp.IsError() {
		err = exc.NewFields("Loki服务器返回错误", field.New("status", rsp.Status()), field.New("body", string(rsp.Body())))
	}

	return
}

func (p *Pusher) gzip(request *resty.Request, data []byte) (buffer *bytes.Buffer, err error) {
	buffer = new(bytes.Buffer)
	writer := gzip.NewWriter(buffer)
	if _, we := writer.Write(data); nil != we {
		err = exc.NewField("压缩数据出错", field.Error(we))
	}
	if nil == err {
		err = writer.Close()
	}
	if nil == err {
		request.SetHeader(key.ContentEncoding, key.Gzip)
	}

	return
}

func (p *Pusher) cleanup(logs *[]*internal.Log) {
	if len(*logs) > 0 {
		_ = p.send(logs)
	}
	p.group.Done()

	return
}
