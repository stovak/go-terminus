package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v3"
)

type HTTPRequest struct {
	Method  string            `yaml:"method"`
	URL     string            `yaml:"url"`
	Body    string            `yaml:"body,omitempty"`
	Headers map[string]string `yaml:"headers,omitempty"`
}

func (r *HTTPRequest) Send() (*http.Response, error) {
	var req *http.Request
	var err error

	if r.Body != "" {
		req, err = http.NewRequest(r.Method, r.URL, strings.NewReader(r.Body))
	} else {
		req, err = http.NewRequest(r.Method, r.URL, nil)
	}
	if err != nil {
		return nil, err
	}

	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func FromYAML(yamlData []byte) (*HTTPRequest, error) {
	var req HTTPRequest
	err := yaml.Unmarshal(yamlData, &req)
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func main() {
	yamlData := `
    method: GET
    url: http://example.com
    headers:
      Content-Type: application/json
    `

	req, err := FromYAML([]byte(yamlData))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := req.Send()
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
