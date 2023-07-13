package request

import (
	"fmt"
	"github.com/stovak/go-terminus/config"
	"io"
	"net/http"
	"net/http/httptrace"
	"net/url"
	"os"
)

type Builder struct {
	Tc *config.TerminusConfig
}

func NewBuilder(tc *config.TerminusConfig) *Builder {
	return &Builder{
		Tc: tc,
	}
}

func (b *Builder) getClient() *http.Client {
	return &http.Client{
		Timeout: b.Tc.Timeout,
	}
}

func (b *Builder) CollectionRequest(m string, path string, body *io.ReadCloser) *http.Request {
	req := b.CreateRequest(m, path, body)
	return req
}

func (b *Builder) ModelRequest(m string, path string, body *io.ReadCloser) *http.Request {
	req := b.CreateRequest(m, path, body)
	return req
}

func (b *Builder) CreateRequest(m string, u string, body *io.ReadCloser) *http.Request {
	r := http.Request{
		Method: m,
		URL:    b.GetUrl(u),
		Header: http.Header{},
		Body:   *body,
	}
	err := b.Tc.Session.AddAuthHeaderToRequest(&r)
	if err != nil {
		return nil
	}
	if b.Tc.Verbose {
		trace := &httptrace.ClientTrace{
			DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
				fmt.Printf("DNS Info: %+v\n", dnsInfo)
			},
			GotConn: func(connInfo httptrace.GotConnInfo) {
				fmt.Printf("Got Conn: %+v\n", connInfo)
			},
			WroteRequest: func(wroteInfo httptrace.WroteRequestInfo) {
				fmt.Printf("Wrote Request: %+v\n", wroteInfo)
			},
			GotFirstResponseByte: func() {
				fmt.Println("Got First Response Byte")
			},
		}
		return r.WithContext(httptrace.WithClientTrace(*b.Tc.GetContext(), trace))
	}
	return r.WithContext(*b.Tc.GetContext())
}

func (b *Builder) SendRequest(req *http.Request) *http.Response {
	resp, err := b.getClient().Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		os.Exit(1)
	}
	return resp
}

func (b *Builder) GetUrl(path string) *url.URL {
	if path == "" {
		panic("Path cannot be empty")
	}
	return &url.URL{
		Scheme: "https",
		Host:   b.Tc.Host,
		Path:   path,
	}
}
