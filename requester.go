package requester

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Poster interface {
	Post(ra RequestArguments) (*http.Response, error)
}

type Getter interface {
	Get(ra RequestArguments) (*http.Response, error)
}

type Request struct {
	Timeout int
	Headers []map[string]interface{}
}

type RequestArguments struct {
	Endpoint string
	Payload  string
}

// Get simply send get http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Get(ra RequestArguments) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, ra.Endpoint, strings.NewReader(ra.Payload))

	if err != nil {
		return nil, err
	}

	r.applyHeadersToRequest(req)

	req.Close = true

	return (&http.Client{Timeout: r.applyTimeout()}).Do(req)
}

// Post simply send post http request to the given endpoint and return *http.Response and error if have it
func (r *Request) Post(ra RequestArguments) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, ra.Endpoint, strings.NewReader(ra.Payload))

	if err != nil {
		return nil, err
	}

	r.applyHeadersToRequest(req)
	req.Close = true

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
