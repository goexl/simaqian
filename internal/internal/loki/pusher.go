package loki

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/goexl/exc"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/simaqian/internal/internal/loki/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Pusher struct {
	config  *Config
	ctx     context.Context
	client  *http.Client
	quit    chan struct{}
	entries chan Entry
	group   sync.WaitGroup
}

func New(ctx context.Context, config Config) (pusher *Pusher) {
	client := new(http.Client)
	config.Url = strings.TrimSuffix(config.Url, "/")
	config.Url = fmt.Sprintf("%s/loki/api/v1/push", config.Url)
	pusher = &Pusher{
		config:  &config,
		ctx:     ctx,
		client:  client,
		quit:    make(chan struct{}),
		entries: make(chan Entry),
	}
	pusher.group.Add(1)
	go pusher.run()

	return
}

func (p *Pusher) Hook(entry zapcore.Entry) error {
	p.entries <- Entry{
		Level:     entry.Level.String(),
		Timestamp: float64(entry.Time.UnixMilli()),
		Message:   entry.Message,
		Caller:    entry.Caller.TrimmedPath(),
	}

	return nil
}

func (p *Pusher) Sink(_ *url.URL) (zap.Sink, error) {
	return NewSink(p), nil
}

func (p *Pusher) Stop() {
	close(p.quit)
	p.group.Wait()
}

func (p *Pusher) Build(config zap.Config, opts ...zap.Option) (logger *zap.Logger, err error) {
	if rse := zap.RegisterSink(internal.KeyLokiSink, p.Sink); nil != rse {
		err = rse
	} else {
		key := fmt.Sprintf("%s://", internal.KeyLokiSink)
		config.OutputPaths = gox.Ift(nil == config.OutputPaths, []string{key}, append(config.OutputPaths, key))
		logger, err = config.Build(opts...)
	}

	return
}

func (p *Pusher) run() {
	var entries []Entry
	ticker := time.NewTimer(p.config.Batch.Wait)
	defer p.cleanup(&entries)

	for {
		select {
		case <-p.ctx.Done():
			return
		case <-p.quit:
			return
		case entry := <-p.entries:
			entries = append(entries, entry)
			if len(entries) >= p.config.Batch.Size {
				_ = p.send(&entries)
				entries = make([]Entry, 0)
				ticker.Reset(p.config.Batch.Wait)
			}
		case <-ticker.C:
			if len(entries) > 0 {
				_ = p.send(&entries)
				entries = make([]Entry, 0)
			}
			ticker.Reset(p.config.Batch.Wait)
		}
	}
}

func (p *Pusher) send(entries *[]Entry) (err error) {
	req := p.assemble(entries)
	if data, me := json.Marshal(req); nil != me {
		err = me
	} else if writer, ge := p.gzip(data); nil != ge {
		err = ge
	} else {
		err = p.post(writer)
	}

	return
}

func (p *Pusher) assemble(entries *[]Entry) (request *Request) {
	request = new(Request)

	var values [][2]string
	for _, entry := range *entries {
		timestamp := time.Unix(int64(entry.Timestamp), 0)
		value := [2]string{strconv.FormatInt(timestamp.UnixNano(), 10), entry.raw}
		values = append(values, value)
	}
	request.Streams = append(request.Streams, Stream{
		Labels: p.config.Labels,
		Values: values,
	})

	return
}

func (p *Pusher) post(writer *bytes.Buffer) (err error) {
	if req, nre := http.NewRequest(internal.Post, p.config.Url, writer); nil != nre {
		err = nre
	} else {
		err = p.http(req)
	}

	return
}

func (p *Pusher) http(request *http.Request) (err error) {
	request.Header.Set(internal.ContentType, internal.Json)
	request.Header.Set(internal.ContentEncoding, internal.Gzip)
	if "" != p.config.Username && "" != p.config.Password {
		request.SetBasicAuth(p.config.Username, p.config.Password)
	}

	rsp, de := p.client.Do(request)
	defer p.httpClose(rsp)

	if nil != de {
		err = de
	} else if rsp.StatusCode != http.StatusNoContent {
		err = exc.NewField("Loki服务器返回错误", field.New("status", rsp.Status))
	}

	return
}

func (p *Pusher) gzip(data []byte) (buffer *bytes.Buffer, err error) {
	buffer = new(bytes.Buffer)
	writer := gzip.NewWriter(buffer)
	if _, we := writer.Write(data); nil != we {
		err = we
	} else if ce := writer.Close(); nil != ce {
		err = ce
	}

	return
}

func (p *Pusher) httpClose(rsp *http.Response) {
	if nil != rsp {
		_ = rsp.Body.Close()
	}
}

func (p *Pusher) cleanup(entries *[]Entry) {
	if len(*entries) > 0 {
		_ = p.send(entries)
	}
	p.group.Done()

	return
}
