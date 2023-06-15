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

func (b *Builder) CollectionRequest(path string) *http.Request {
	req := b.CreateRequest("GET", path, nil)
	resp := b.SendRequest(req)

}

func (b *Builder) ModelRequest(path string) *http.Request {
	req := b.CreateRequest("GET", path, nil)
	resp := b.SendRequest(req)

}

func (b *Builder) CreateRequest(m string, u string, body io.ReadCloser) *http.Request {
	r := http.Request{
		Method: m,
		URL:    b.GetUrl(u),
		Header: http.Header{},
		Body:   body,
	}
	// Dereference and turn pinnocio into a real boy
	req := r.WithContext(*b.Tc.GetContext())
	b.Tc.Session.AddSessionHeader(req)
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
		req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	}
	return req
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
