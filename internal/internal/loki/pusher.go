package loki

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian/internal/internal/internal"
	"github.com/goexl/simaqian/internal/internal/loki/internal/key"
	"github.com/golang/snappy"
	"github.com/grafana/loki/pkg/logproto"
	"go.uber.org/zap"
)

type Pusher struct {
	config *Config
	ctx    context.Context
	client *http.Client
	quit   chan gox.Empty
	logs   chan *internal.Log
	group  sync.WaitGroup
}

func New(ctx context.Context, config Config) (pusher *Pusher) {
	client := new(http.Client)
	config.Url = strings.TrimSuffix(config.Url, "/")
	config.Url = fmt.Sprintf("%s/loki/api/v1/push", config.Url)
	pusher = &Pusher{
		config: &config,
		ctx:    ctx,
		client: client,
		quit:   make(chan gox.Empty),
		logs:   make(chan *internal.Log),
	}
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
		logger, err = config.Build(opts...)
	}

	return
}

func (p *Pusher) run() {
	logs := make([]*internal.Log, 0, p.config.Batch.Size)
	ticker := time.NewTimer(p.config.Batch.Wait)
	defer p.cleanup(&logs)

	for {
		select {
		case <-p.ctx.Done():
			break
		case <-p.quit:
			break
		case log := <-p.logs:
			logs = append(logs, log)
			if len(logs) >= p.config.Batch.Size {
				_ = p.send(&logs)
				logs = make([]*internal.Log, 0)
				ticker.Reset(p.config.Batch.Wait)
			}
		case <-ticker.C:
			if len(logs) > 0 {
				_ = p.send(&logs)
				logs = make([]*internal.Log, 0)
			}
			ticker.Reset(p.config.Batch.Wait)
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
		Labels:  p.config.Labels.String(),
		Entries: entries,
	})

	return
}

func (p *Pusher) post(data []byte) (err error) {
	data = snappy.Encode(nil, data)
	buffer := bytes.NewBuffer(data)
	if req, nre := http.NewRequest(key.Post, p.config.Url, buffer); nil != nre {
		err = nre
	} else {
		err = p.do(req)
	}

	return
}

func (p *Pusher) do(request *http.Request) (err error) {
	request.Header.Set(key.ContentType, key.Protobuf)
	if "" != p.config.Username && "" != p.config.Password {
		request.SetBasicAuth(p.config.Username, p.config.Password)
	}

	rsp, de := p.client.Do(request)
	defer p.close(rsp)

	if nil != de {
		err = de
	} else if rsp.StatusCode != http.StatusNoContent {
		err = exc.NewField("Loki服务器返回错误", field.New("status", rsp.Status))
	}

	return
}

func (p *Pusher) close(rsp *http.Response) {
	if nil != rsp {
		_ = rsp.Body.Close()
	}
}

func (p *Pusher) cleanup(logs *[]*internal.Log) {
	if len(*logs) > 0 {
		_ = p.send(logs)
	}
	p.group.Done()

	return
}
