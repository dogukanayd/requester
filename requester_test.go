package requester

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newMockServer(res []byte, statusCode int, headers map[string]string) *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		for key, value := range headers {
			w.Header().Set(key, value)
		}

		w.WriteHeader(statusCode)
		_, _ = w.Write(res)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

func TestRequest_Get(t *testing.T) {
	t.Run("it_should_return_error_when_try_to_create_new_request", func(t *testing.T) {
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: "://///////\\***",
			Body:     []byte(""),
		}

		_, err := request.Get()

		if err == nil {
			t.Errorf("expected error but return nil")
		}
	})
	t.Run("it_should_not_return_error_when_sending_get_request", func(t *testing.T) {
		mockResponse := []byte(`{"status": "accepted"}`)
		mockServer := newMockServer(mockResponse, http.StatusAccepted, nil)
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: mockServer.URL,
			Body:     []byte(""),
		}
		response, err := request.Get()

		if response.StatusCode != http.StatusAccepted {
			t.Error(fmt.Sprintf("expected status code is: %d, got: %d", http.StatusAccepted, response.StatusCode))
		}

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("it_should_return_error_when_can_not_parse_the_url_and_payload", func(t *testing.T) {
		corruptedURL := "```"
		corruptedPayload := []byte("```")
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
			},
			Endpoint: corruptedURL,
			Body:     corruptedPayload,
		}
		_, err := request.Get()

		if err == nil {
			t.Error("expected error but return success")
		}
	})
}

func TestRequest_Post(t *testing.T) {
	t.Run("it_should_return_error_when_try_to_create_new_request", func(t *testing.T) {
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: "://///////\\***",
			Body:     []byte(""),
		}

		_, err := request.Post()

		if err == nil {
			t.Errorf("expected error but return nil")
		}
	})

	t.Run("it_should_not_return_error_when_sending_post_request", func(t *testing.T) {
		mockResponse := []byte(`{"status": "accepted"}`)
		mockServer := newMockServer(mockResponse, http.StatusAccepted, nil)
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: mockServer.URL,
			Body:     []byte(""),
		}
		response, err := request.Post()

		if response.StatusCode != http.StatusAccepted {
			t.Error(fmt.Sprintf("expected status code is: %d, got: %d", http.StatusAccepted, response.StatusCode))
		}

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("it_should_return_error_when_can_not_parse_the_url_and_payload", func(t *testing.T) {
		corruptedURL := "```"
		corruptedPayload := []byte("```")
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
			},
			Endpoint: corruptedURL,
			Body:     corruptedPayload,
		}
		_, err := request.Post()

		if err == nil {
			t.Error("expected error but return success")
		}
	})
}

func TestRequest_Put(t *testing.T) {
	t.Run("it_should_return_error_when_try_to_create_new_request", func(t *testing.T) {
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: "://///////\\***",
			Body:     []byte(""),
		}

		_, err := request.Put()

		if err == nil {
			t.Errorf("expected error but return nil")
		}
	})

	t.Run("it_should_not_return_error_when_sending_put_request", func(t *testing.T) {
		mockResponse := []byte(`{"status": "accepted"}`)
		mockServer := newMockServer(mockResponse, http.StatusAccepted, nil)
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: mockServer.URL,
			Body:     []byte(""),
		}
		response, err := request.Put()

		if response.StatusCode != http.StatusAccepted {
			t.Error(fmt.Sprintf("expected status code is: %d, got: %d", http.StatusAccepted, response.StatusCode))
		}

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("it_should_return_error_when_can_not_parse_the_url_and_payload", func(t *testing.T) {
		corruptedURL := "```"
		corruptedPayload := []byte("```")

		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
			},
			Endpoint: corruptedURL,
			Body:     corruptedPayload,
		}
		_, err := request.Put()

		if err == nil {
			t.Error("expected error but return success")
		}
	})
}

func TestRequest_Delete(t *testing.T) {
	t.Run("it_should_return_error_when_try_to_create_new_request", func(t *testing.T) {
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: "://///////\\***",
			Body:     []byte(""),
		}

		_, err := request.Delete()

		if err == nil {
			t.Errorf("expected error but return nil")
		}
	})
	t.Run("it_should_not_return_error_when_sending_delete_request", func(t *testing.T) {
		mockResponse := []byte(`{"status": "accepted"}`)
		mockServer := newMockServer(mockResponse, http.StatusAccepted, nil)
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
				{
					"Host": "test.test.com",
				},
			},
			Endpoint: mockServer.URL,
			Body:     []byte(""),
		}
		response, err := request.Delete()

		if response.StatusCode != http.StatusAccepted {
			t.Error(fmt.Sprintf("expected status code is: %d, got: %d", http.StatusAccepted, response.StatusCode))
		}

		if err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("it_should_return_error_when_can_not_parse_the_url_and_payload", func(t *testing.T) {
		corruptedURL := "```"
		corruptedPayload := []byte("```")
		request := Request{
			Timeout: 60,
			Headers: []map[string]interface{}{
				{
					"Content-Type": "application/json",
				},
			},
			Endpoint: corruptedURL,
			Body:     corruptedPayload,
		}
		_, err := request.Delete()

		if err == nil {
			t.Error("expected error but return success")
		}
	})
}

func TestRequest_applyTimeout(t *testing.T) {
	t.Run("it_should_apply_time_out_as_30_seconds_if_there_is_no_defined_timeout", func(t *testing.T) {
		request := Request{}
		expectedTimeout := time.Duration(30) * time.Second

		if request.applyTimeout() != expectedTimeout {
			t.Error(fmt.Sprintf("expected time out %d, got: %d", expectedTimeout, request.Timeout))
		}
	})
}
