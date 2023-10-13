package loki

import (
	"context"
	"fmt"
	"net/url"
	"sync"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/http"
	"github.com/goexl/simaqian/internal/config"
	"github.com/goexl/simaqian/internal/internal/internal"
	"github.com/goexl/simaqian/internal/internal/loki/internal/key"
	"github.com/golang/snappy"
	"github.com/grafana/loki/pkg/logproto"
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
	pusher.http.SetBasicAuth(config.Username, config.Password).SetHeader(key.ContentType, key.Protobuf)
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
	if push, ae := p.make(logs); nil != ae {
		err = ae
	} else if data, me := push.Marshal(); nil != me {
		err = me
	} else {
		err = p.post(data)
	}
	if nil != err {
		p.logger.Warn("推送日志出错", zap.Error(err))
	}

	return
}

func (p *Pusher) make(logs *[]*internal.Log) (request *logproto.PushRequest, err error) {
	request = new(logproto.PushRequest)
	entries := make([]logproto.Entry, 0, len(*logs))
	for _, log := range *logs {
		entries = append(entries, logproto.Entry{
			Timestamp: log.Timestamp,
			Line:      log.Raw(),
		})
	}
	request.Streams = append(request.Streams, logproto.Stream{
		Labels:  p.labels.String(),
		Entries: entries,
	})

	return
}

func (p *Pusher) post(data []byte) (err error) {
	data = snappy.Encode(nil, data)
	if rsp, pe := p.http.R().SetBody(data).Post(p.url); nil != pe {
		err = pe
	} else if rsp.IsError() {
		err = exc.NewFields("Loki服务器返回错误", field.New("status", rsp.Status()), field.New("body", string(rsp.Body())))
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
