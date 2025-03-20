package pepay

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

const defaultTestBaseURL = "https://psp.stg.01123581.com"
const clientKey = "0ee4453c-d8ef-4e39-9deb-a897acc74713"
const clientSecret = "YzNpUzVKLyNMeiZaRWJgdyRcWVNdRn5lOmt8RWptUSZ9YFZYTS1AOHxsTG9GcXh9Ryp+J1R5O28xMX1TZEpBbz9MdWhnZHNZIXMrKFI3LEZocSll"

func NewFakeClient(t *testing.T, tr RoundTripFunc) *Client {
	t.Helper()
	a := AuthTransport{
		ClientKey: clientKey, ClientSecret: clientSecret,
	}
	c := NewClient(a.Client(), defaultTestBaseURL)
	c.BaseURL, _ = url.Parse(defaultTestBaseURL)
	return c
}

func TestNewRequest(t *testing.T) {
	const path = "/some/path"
	fn := func(r *http.Request) *http.Response {
		if got, want := r.URL.String(), fmt.Sprintf("%s%s", defaultTestBaseURL, path); got != want {
			t.Errorf("wrong request url got = %v, want =%v", got, want)
		}
		if got, want := r.Header.Get("key"), clientKey; got != want {
			t.Errorf("Invalid Key header: got = %v, want = %v", got, want)
		}
		defer r.Body.Close()

		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(`{"status": "SUCCESS"}`)),
		}
	}
	client := NewFakeClient(t, fn)
	req, err := client.NewRequest("GET", path, "")
	if err != nil {
		t.Errorf("Unexpected error when creating the request: %v", err)
	}
	var buf bytes.Buffer
	res, err := client.Do(req, &buf)
	if err != nil {
		t.Errorf("Unexpected error when receiving the server response: %v", err)
	}
	if got, want := buf.String(), `{"status": "SUCCESS"}`; got != want {
		t.Errorf("Unexpected response body: got = %v, want = %v", got, want)
	}
	if got, want := res.StatusCode, http.StatusOK; got != want {
		t.Errorf("Unexpected status code: got = %v, want = %v", got, want)
	}
}
