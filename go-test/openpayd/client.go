package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const defaultBaseURL = "https://secure.openpayd.com"

type OpenpayService struct {
	baseURL *url.URL
	client  *http.Client
}

func NewOpenpayService(c *http.Client, baseURL *url.URL) *OpenpayService {
	if baseURL == nil {
		var err error
		baseURL, err = url.Parse(defaultBaseURL)
		if err != nil {
			panic(err)
		}
	}
	if c == nil {
		c = http.DefaultClient
	}

	return &OpenpayService{
		baseURL: baseURL,
		client:  c,
	}
}

func (s OpenpayService) FetchTxByID(ctx context.Context, id TxID) (Tx, error) {
	u, err := s.baseURL.Parse("/api/transactions/" + string(id))
	if err != nil {
		return Tx{}, err
	}
	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)
	if err != nil {
		return Tx{}, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		return Tx{}, err
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return Tx{}, err
	}
	var v struct {
		ID                   TxID            `json:"transactionId,omitempty"`
		ShortID              string          `json:"shortId,omitempty"`
		AccountID            string          `json:"accountId,omitempty"`
		CreatedDate          time.Time       `json:"createdDate,omitempty"`
		UpdatedDate          time.Time       `json:"updatedDate,omitempty"`
		PaymentDate          time.Time       `json:"paymentDate,omitempty"`
		TransactionID        string          `json:"id,omitempty"`
		TransactionCategory  string          `json:"transactionCategory,omitempty"`
		PaymentType          string          `json:"paymentType,omitempty"`
		Type                 TxType          `json:"type,omitempty"`
		SourceInfo           SourceInfo      `json:"sourceInfo,omitempty"`
		DestinationInfo      DestinationInfo `json:"destinationInfo,omitempty"`
		Source               string          `json:"source,omitempty"`
		Destination          string          `json:"destination,omitempty"`
		TotalAmount          *RequestAmount  `json:"totalAmount,omitempty"`
		Amount               *RequestAmount  `json:"amount,omitempty"`
		Fee                  *RequestAmount  `json:"fee,omitempty"`
		RunningBalance       *RequestAmount  `json:"runningBalance,omitempty"`
		BuyAmount            *RequestAmount  `json:"buyAmount,omitempty"`
		FxRate               float64         `json:"fxRate,omitempty"`
		MidMarketRate        float64         `json:"midMarketRate,omitempty"`
		FixedSide            string          `json:"fixedSide,omitempty"`
		Status               TxStatus        `json:"status,omitempty"`
		FailureReason        string          `json:"failureReason,omitempty"`
		Comment              string          `json:"comment,omitempty"`
		TransactionReference string          `json:"transactionReference,omitempty"`
		ReferenceAmount      *RequestAmount  `json:"referenceAmount,omitempty"`
		AccountHolderID      string          `json:"accountHolderId,omitempty"`
	}
	err = json.NewDecoder(res.Body).Decode(&v)
	if err == io.EOF {
		err = nil // ignore EOF errors caused by empty response body
	}
	return Tx{
		ID:                   v.ID,
		ShortID:              v.ShortID,
		AccountID:            v.AccountID,
		CreatedDate:          v.CreatedDate,
		UpdatedDate:          v.UpdatedDate,
		PaymentDate:          v.PaymentDate,
		EncID:                v.TransactionID,
		TransactionCategory:  v.TransactionCategory,
		PaymentType:          v.PaymentType,
		Type:                 v.Type,
		SourceInfo:           v.SourceInfo,
		DestinationInfo:      v.DestinationInfo,
		Source:               v.Source,
		Destination:          v.Destination,
		TotalAmount:          requestAmount(v.TotalAmount),
		Amount:               requestAmount(v.Amount),
		Fee:                  requestAmount(v.Fee),
		RunningBalance:       requestAmount(v.RunningBalance),
		BuyAmount:            requestAmount(v.BuyAmount),
		FxRate:               v.FxRate,
		MidMarketRate:        v.MidMarketRate,
		FixedSide:            v.FixedSide,
		Status:               v.Status,
		FailureReason:        v.FailureReason,
		Comment:              v.Comment,
		TransactionReference: v.TransactionReference,
		ReferenceAmount:      requestAmount(v.ReferenceAmount),
		AccountHolderID:      v.AccountHolderID,
	}, err
}

type AuthTransport struct {
	ClientID     string
	ClientSecret string
	// Transport is the underlying HTTP transport to use when making requests.
	// It will default to http.DefaultTransport if nil.
	Transport http.RoundTripper
}

func (t *AuthTransport) tokenSource(r *http.Request) TokenSource {
	ts := tokensCache.Load(t.ClientID, t.ClientSecret)
	if ts != nil {
		return ts
	}
	u := *r.URL
	u.Path = ""
	u.RawQuery = ""
	u.Fragment = ""
	src := &tokenSource{
		ctx:      r.Context(),
		ClientID: t.ClientID,
		Secret:   t.ClientSecret,
		TokenURL: u.String(),
	}
	ts = &reuseTokenSource{
		new: src,
	}
	tokensCache.Store(t.ClientID, t.ClientSecret, ts)
	return ts
}

func (t *AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	ts := t.tokenSource(req)
	token, err := ts.Token()
	if err != nil {
		tokensCache.Delete(t.ClientID, t.ClientSecret)
		return nil, err
	}
	req2 := cloneRequest(req)
	req2.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.Token))
	req2.Header.Set("X-ACCOUNT-HOLDER-ID", token.AccountHolderID)
	return t.transport().RoundTrip(req2)
}

func (t *AuthTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
	}
}

func (t *AuthTransport) SetTransport(rt http.RoundTripper) {
	t.Transport = rt
}

func (t *AuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// TokenSource is anything that can return a token.
type TokenSource interface {
	Token() (*Token, error)
}

type Token struct {
	Token           string
	AccountHolderID string
	Type            string
	Expires         time.Time
}

func (t *Token) Valid() bool {
	expiryDelta := 10 * time.Second
	return t != nil && t.Expires.Add(-expiryDelta).After(time.Now())
}

var tokensCache tokenCache

type tokenCache struct {
	m sync.Map
}

func (tc *tokenCache) Len() int {
	count := 0
	tc.m.Range(func(key, value any) bool {
		count++
		return true
	})
	return count
}

func (tc *tokenCache) key(id, secret string) [32]byte {
	b := []byte(id + secret)
	return sha256.Sum256(b)
}

func (tc *tokenCache) Delete(id, secret string) {
	tc.m.Delete(tc.key(id, secret))
}

func (tc *tokenCache) Load(id, secret string) TokenSource {
	v, ok := tc.m.Load(tc.key(id, secret))
	if !ok {
		return nil
	}
	return v.(TokenSource)
}

func (tc *tokenCache) Store(id, secret string, ts TokenSource) {
	tc.m.Store(tc.key(id, secret), ts)
}

func (tc *tokenCache) Clear() {
	tc.m = sync.Map{}
}

type reuseTokenSource struct {
	new TokenSource

	mu  sync.Mutex
	tok *Token
}

func (ts *reuseTokenSource) Token() (*Token, error) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	if ts.tok.Valid() {
		return ts.tok, nil
	}
	t, err := ts.new.Token()
	if err != nil {
		return nil, err
	}
	ts.tok = t
	return t, nil
}

type tokenSource struct {
	ctx      context.Context
	ClientID string
	Secret   string
	TokenURL string
}

func (ts *tokenSource) Token() (*Token, error) {
	tk, err := fetchToken(ts.TokenURL, ts.ClientID, ts.Secret)
	if err != nil {
		return nil, fmt.Errorf("fetch auth token: %w", err)
	}
	return tk, err
}

func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}

func fetchToken(baseURL, clientID, clientSecret string) (*Token, error) {
	if baseURL == "" {
		baseURL = defaultBaseURL
	}
	req, err := http.NewRequest("POST", baseURL+"/api/oauth/token?grant_type=client_credentials", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()
	if err := CheckResponse(rsp); err != nil {
		return nil, err
	}

	var v AuthResponse
	err = json.NewDecoder(rsp.Body).Decode(&v)
	if err == io.EOF {
		err = nil // ignore EOF errors caused by empty response body
	}
	if err != nil {
		return nil, err
	}
	exp := time.Unix(time.Now().Unix()+v.ExpiresIn, 0)
	if err != nil {
		return nil, err
	}

	return &Token{
		Token:           v.AccessToken,
		Expires:         exp,
		AccountHolderID: v.AccountHolderID,
	}, nil
}

type AuthResponse struct {
	AccessToken         string   `json:"access_token"`
	TokenType           string   `json:"token_type"`
	ExpiresIn           int64    `json:"expires_in"`
	Scope               string   `json:"scope"`
	AccountHolderID     string   `json:"accountHolderId"`
	ClientID            string   `json:"clientId"`
	ReferralID          string   `json:"referralId"`
	AccountHolderStatus string   `json:"accountHolderStatus"`
	ClientTenantID      string   `json:"clientTenantId"`
	Authorities         []string `json:"authorities"`
	Jti                 string   `json:"jti"`
	AccountHolderType   string   `json:"accountHolderType"`
}

func CheckResponse(r *http.Response) (err error) {
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}
	b := make([]byte, 256)
	if _, err := r.Body.Read(b); err != nil {
		return fmt.Errorf("read body: %s", err)
	}
	defer r.Body.Close()

	return fmt.Errorf("http request status: %d: %s", r.StatusCode, b)
}
