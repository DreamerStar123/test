package pepay

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL   = "https://api.fortris.com"
	defaultMediaType = "application/json"
)

// A Client manages communication with the PEPay API.
type Client struct {
	client  *http.Client
	BaseURL *url.URL

	common service

	Payout  *PayoutService
	Deposit *DepositService
}

type service struct {
	client *Client
}

// NewClient returns a new PEPay API client. If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client, base_url string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)
	if base_url != "" {
		baseURL, _ = url.Parse(base_url)
	}

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.common.client = c
	c.Payout = (*PayoutService)(&c.common)
	c.Deposit = (*DepositService)(&c.common)
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body any) (*http.Request, error) {
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	return c.newRequest(method, u, body)
}

func (c *Client) newRequest(method string, u *url.URL, body any) (*http.Request, error) {
	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), &buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", defaultMediaType)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v any) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return resp, err
	}
	defer resp.Body.Close()
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil // ignore EOF errors caused by empty response body
			}
		}
	}
	return resp, err
}

type ResponseStatus int8

const (
	ResponseStatusOK ResponseStatus = iota + 1
	ResponseStatusFailed
)

func (s *ResponseStatus) UnmarshalText(text []byte) error {
	switch text := string(text); text {
	case "SUCCESS":
		*s = ResponseStatusOK
	case "FAIL":
		*s = ResponseStatusFailed
	}
	return nil
}

type ValidationError struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

type Error struct {
	Message   string `json:"error"`
	Code      string `json:"code"`
	Timestamp string `json:"timestamp"`
	HttpCode  int    `json:"httpcode"`
}

func (e *Error) UnWrap() error {
	return nil
}

func (e *Error) Error() string {
	out := fmt.Sprintf("%s: %s", e.Code, e.Message)
	return out
}

func CheckResponse(r *http.Response) error {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	r.Body.Close()
	if data == nil {
		return nil // Empty body
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))

	var response Error
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	response.HttpCode = r.StatusCode
	return &response
}

// AuthTransport is a http.RoundTripper that authenticates all requests
// sent to pbpay.
type AuthTransport struct {
	ClientKey, ClientSecret string

	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func calcSign(dataToSign, secret string) string {
	// Decode the base64-encoded secret
	decodedSecret, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		fmt.Println("Error decoding secret:", err)
		return "error"
	}

	// Compute HMAC-SHA512 signature
	h := hmac.New(sha512.New, decodedSecret)
	h.Write([]byte(dataToSign))
	signature := hex.EncodeToString(h.Sum(nil))
	// println(signature)

	return signature
}

// RoundTrip implements the RoundTripper interface.
func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := new(http.Request)
	*req2 = *req
	req2.Header = make(http.Header, len(req.Header))
	for k, s := range req.Header {
		req2.Header[k] = append([]string(nil), s...)
	}
	var b bytes.Buffer
	if req.Body != nil && req.Body != http.NoBody && req.GetBody != nil {
		body, err := req.GetBody()
		if err != nil {
			return nil, err
		}
		if _, err := b.ReadFrom(body); err != nil {
			return nil, err
		}
	}

	dataToSign := req.URL.Path
	if req.URL.RawQuery != "" {
		dataToSign += "?" + req.URL.RawQuery
	}

	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	payload := string(bodyBytes)

	if payload != "" {
		// Compute SHA-256 hash of the payload
		hash := sha256.Sum256([]byte(payload))
		dataToSign += hex.EncodeToString(hash[:])
	}

	signature := calcSign(dataToSign, t.ClientSecret)

	// Authentication Headers
	req2.Header.Add("key", t.ClientKey)
	req2.Header.Add("signature", signature)
	req2.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return t.transport().RoundTrip(req2)
}

// Client returns an *http.Client that makes requests that are authenticated
// using BitBay HTTP Authentication headers.
func (t *AuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}
