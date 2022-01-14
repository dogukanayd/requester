package requester

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Requester represent the package structure, with creating exactly the same interface your own codebase you can
// easily mock the functions inside this package while writing unit tests.
type Requester interface {
	Get(re RequestEntity) (*http.Response, error)
	Post(re RequestEntity) (*http.Response, error)
	Put(re RequestEntity) (*http.Response, error)
	Delete(re RequestEntity) (*http.Response, error)
}

// RequestEntity contains required information for sending http request
type RequestEntity struct {
	Timeout  int
	Headers  []map[string]interface{}
	Endpoint string
	Body     []byte
}

type Request struct{}

// Get simply send get http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Get(re RequestEntity) (*http.Response, error) {
	return r.sendRequest(http.MethodGet, re)
}

// Post simply send post http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Post(re RequestEntity) (*http.Response, error) {
	return r.sendRequest(http.MethodPost, re)
}

// Put simply execute put http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Put(re RequestEntity) (*http.Response, error) {
	return r.sendRequest(http.MethodPut, re)
}

// Delete send delete method
func (r *Request) Delete(re RequestEntity) (*http.Response, error) {
	return r.sendRequest(http.MethodDelete, re)
}

func (r *Request) sendRequest(httpMethod string, re RequestEntity) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, re.Endpoint, re.readBody())

	if err != nil {
		return nil, err
	}

	req.Close = true
	re.applyHeadersToRequest(req)

	return (&http.Client{Timeout: re.applyTimeout()}).Do(req)
}

func (r *RequestEntity) applyHeadersToRequest(request *http.Request) {
	for _, header := range r.Headers {
		for key, value := range header {
			if key == "Host" {
				request.Host = fmt.Sprintf("%v", value)
			} else {
				request.Header.Set(key, fmt.Sprintf("%v", value))
			}
		}
	}
}

func (r *RequestEntity) applyTimeout() time.Duration {
	if r.Timeout <= 0 {
		r.Timeout = 30

		return time.Duration(r.Timeout) * time.Second
	}

	return time.Duration(r.Timeout) * time.Second
}

func (r *RequestEntity) readBody() *strings.Reader {
	return strings.NewReader(string(r.Body))
}
