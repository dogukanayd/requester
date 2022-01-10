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
	Get() (*http.Response, error)
	Post() (*http.Response, error)
	Put() (*http.Response, error)
	Delete() (*http.Response, error)
}

type Request struct {
	Timeout  int
	Headers  []map[string]interface{}
	Endpoint string
	Body     []byte
}

// Get simply send get http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Get() (*http.Response, error) {
	return r.sendRequest(http.MethodGet)
}

// Post simply send post http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Post() (*http.Response, error) {
	return r.sendRequest(http.MethodPost)
}

// Put simply execute put http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Put() (*http.Response, error) {
	return r.sendRequest(http.MethodPut)
}

// Delete send delete method
func (r *Request) Delete() (*http.Response, error) {
	return r.sendRequest(http.MethodDelete)
}

func (r *Request) sendRequest(httpMethod string) (*http.Response, error) {
	req, err := http.NewRequest(httpMethod, r.Endpoint, r.readBody())

	if err != nil {
		return nil, err
	}

	req.Close = true
	r.applyHeadersToRequest(req)

	return (&http.Client{Timeout: r.applyTimeout()}).Do(req)
}

func (r *Request) applyHeadersToRequest(request *http.Request) {
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

func (r *Request) applyTimeout() time.Duration {
	if r.Timeout <= 0 {
		r.Timeout = 30

		return time.Duration(r.Timeout) * time.Second
	}

	return time.Duration(r.Timeout) * time.Second
}

func (r *Request) readBody() *strings.Reader {
	return strings.NewReader(string(r.Body))
}
