package main

import (
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalid         = errors.New("invalid argument")           // validation failed
	ErrPermission      = errors.New("permission denied")          // permission error action cannot be perform.
	ErrExist           = errors.New("already exists")             // entity does exist
	ErrNotExist        = errors.New("does not exist")             // entity does not exist
	ErrConflict        = errors.New("action cannot be performed") // action cannot be performed
	ErrInvalidAuthKey  = errors.New("invalid auth key")
	ErrSessionPaid     = errors.New("already paid")
	ErrSessionExpired  = errors.New("session expired")
	ErrPaymentDeclined = errors.New("payment declined")
	ErrReattempt       = errors.New("short time for reattempt")
)

type TxID string

type TxType int

const (
	PayIn TxType = iota + 1
	PayOut
	Transfer
	Adjustment
	ReturnIn
	ReturnOut
	Fee
)

func (t *TxType) UnmarshalText(b []byte) error {
	s := string(b)
	switch s {
	case "PAYIN":
		*t = PayIn
	case "PAYOUT":
		*t = PayOut
	case "TRANSFER":
		*t = Transfer
	case "ADJUSTMENT":
		*t = Adjustment
	case "RETURNIN":
		*t = ReturnIn
	case "RETURNOUT":
		*t = ReturnOut
	case "FEE":
		*t = Fee
	}
	return nil
}

func (t TxType) MarshalText() ([]byte, error) {
	var s string
	switch t {
	case PayIn:
		s = "PAYIN"
	case PayOut:
		s = "PAYOUT"
	case Transfer:
		s = "TRANSFER"
	case Adjustment:
		s = "ADJUSTMENT"
	case ReturnIn:
		s = "RETURNIN"
	case ReturnOut:
		s = "RETURNOUT"
	case Fee:
		s = "FEE"
	}
	return []byte(s), nil
}

type TxStatus int

const (
	TxStatusIni TxStatus = iota
	TxStatusPending
	TxStatusReleased
	TxStatusCompleted
	TxStatusFailed
	TxStatusCancelled
)

func (t *TxStatus) UnmarshalText(b []byte) error {
	s := string(b)
	switch s {
	case "PROCESSING":
		*t = TxStatusPending
	case "RELEASED":
		*t = TxStatusReleased
	case "COMPLETED":
		*t = TxStatusCompleted
	case "FAILED":
		*t = TxStatusFailed
	case "CANCELLED":
		*t = TxStatusCancelled
	}
	return nil
}

func (t TxStatus) MarshalText() ([]byte, error) {
	var s string
	switch t {
	case TxStatusIni:
		s = "INITIATE"
	case TxStatusPending:
		s = "PROCESSING"
	case TxStatusReleased:
		s = "RELEASED"
	case TxStatusCompleted:
		s = "COMPLETED"
	case TxStatusFailed:
		s = "FAILED"
	case TxStatusCancelled:
		s = "CANCELLED"
	}
	return []byte(s), nil
}

type Tx struct {
	ID                   TxID
	ShortID              string
	AccountID            string
	CreatedDate          time.Time
	UpdatedDate          time.Time
	PaymentDate          time.Time
	EncID                string
	TransactionCategory  string
	PaymentType          string
	Type                 TxType
	SourceInfo           SourceInfo
	DestinationInfo      DestinationInfo
	Source               string
	Destination          string
	TotalAmount          Amount
	Amount               Amount
	Fee                  Amount
	RunningBalance       Amount
	BuyAmount            Amount
	ReferenceAmount      Amount
	FxRate               float64
	MidMarketRate        float64
	FixedSide            string
	Status               TxStatus
	FailureReason        string
	Comment              string
	TransactionReference string
	AccountHolderID      string
	SenderBic            string
	SenderAccountNumber  string
	SenderRoutingCodes   []SenderRoutingCodes
	SenderIban           string
	SenderName           string
	SenderAddress        string
	SenderInformation    string
	OriginalTxID         string

	TenantID   string
	MerchantID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (tx Tx) IsPending() bool {
	return tx.Status == TxStatusIni || tx.Status == TxStatusPending || tx.Status == TxStatusReleased
}

type SourceInfo struct {
	Type              string `json:"type,omitempty"`
	Identifier        string `json:"identifier,omitempty"`
	InternalAccountID string `json:"internalAccountId,omitempty"`
}

func (s SourceInfo) IsZero() bool {
	return s == (SourceInfo{})
}

type DestinationInfo struct {
	Type              string `json:"type,omitempty"`
	Identifier        string `json:"identifier,omitempty"`
	InternalAccountID string `json:"internalAccountId,omitempty"`
}

func (d DestinationInfo) IsZero() bool {
	return d == (DestinationInfo{})
}

type Amount struct {
	Value    int64    `json:"value"`
	Currency Currency `json:"currency"`
}

func (a Amount) IsZero() bool {
	return a.Value == 0 && a.Currency.String() == ""
}

type Merchant struct {
	ID           string
	Name         string
	TenantID     string
	Enabled      bool
	AuthKey      string
	WebhookURL   string
	BaseURL      string
	ClientID     string
	ClientSecret string
	PublicKey    string
}

type EventID string

type EventAttempt struct {
	ID         int64     // Unique ID for the attempt
	EventID    EventID   // Foreign key to the event
	CreatedAt  time.Time // Timestamp of the attempt
	ErrorMsg   string    // Error message if the attempt failed
	StatusCode int
	WebhookURL string
}

type EventFilter struct {
	ID            *EventID
	TxID          *TxID
	Delivered     *bool
	Retry         *bool
	NextAttemptAt *time.Time
	Limit         int
	Offset        int
}

// Event represents an event that occurs in the system. These events are
// eventually propagated out to connected users via Webhook URL whenever changes
// occur.
type Event struct {
	ID                           EventID
	Type                         TxType
	TransactionID                TxID
	Status                       TxStatus
	InternalAccountID            string
	PaymentType                  string
	SenderBic                    string
	SenderAccountNumber          string
	SenderRoutingCodes           []SenderRoutingCodes
	SenderIban                   string
	SenderName                   string
	SenderAddress                string
	SenderInformation            string
	SourceInfo                   SourceInfo
	DestinationInfo              DestinationInfo
	Source                       string
	Destination                  string
	TotalAmount                  Amount
	Amount                       Amount
	Fee                          Amount
	RunningBalance               Amount
	BuyAmount                    Amount
	FxRate                       float64
	MidMarketRate                float64
	FixedSide                    string
	FailureReason                string
	Comment                      string
	TransactionReference         string
	ReferenceAmount              Amount
	AccountHolderID              string
	TransactionDateTime          time.Time
	BeneficiaryAccountHolderName string
	CreatedBy                    string
	MandateID                    string
	TransactionCategory          string
	CreatedDate                  time.Time
	UpdatedDate                  time.Time
	OriginalTxID                 string

	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeliveredAt   time.Time
	WebhookURL    string
	FailCount     int
	NextAttemptAt time.Time
	Retry         bool
	MerchantID    string
}

func (e *Event) IsDelivered() bool {
	return !e.DeliveredAt.IsZero()
}

func (e *Event) MarkForRetry() {
	if e.IsDelivered() {
		e.Retry = false
		return
	}
	if e.NextAttemptAt.IsZero() {
		e.NextAttemptAt = time.Now()
	}
	if e.FailCount >= 16 {
		e.Retry = false
		e.NextAttemptAt = time.Time{}
	}
	e.FailCount++
	e.Retry = true
	failCount := e.FailCount
	baseInterval := 10 * time.Second       // Base interval for retries
	maxInitialInterval := 40 * time.Second // Max interval before exponential backoff kicks in
	growthFactor := 2                      // Exponential growth factor

	// Calculate the exponential backoff

	if failCount > 0 {
		// If the fail count is such that it would exceed the maxInitialInterval
		if time.Duration(failCount)*baseInterval <= maxInitialInterval {
			e.NextAttemptAt = e.NextAttemptAt.Add(baseInterval)
		} else {
			// Exponential backoff after reaching maxInitialInterval
			e.NextAttemptAt = e.NextAttemptAt.Add(time.Duration(math.Pow(float64(growthFactor), float64(failCount-1))) * baseInterval)
		}
	}
}

func (e *Event) MarkAsDelivered() {
	if e.IsDelivered() {
		return
	}
	now := time.Now()
	e.DeliveredAt = now
	e.UpdatedAt = now
	e.Retry = false
	e.NextAttemptAt = time.Time{}
}

func NewEvent(id EventID, typ TxType, webhookURL string) (Event, error) {
	if webhookURL != "" {
		_, err := url.ParseRequestURI(webhookURL)
		if err != nil {
			return Event{}, errors.Join(err, ErrInvalid)
		}
	}
	now := time.Now()
	return Event{
		ID:         id,
		Type:       typ,
		WebhookURL: webhookURL,
		CreatedAt:  now,
		UpdatedAt:  now,
	}, nil
}

type SenderRoutingCodes struct {
	RoutingCodeKey   string `json:"routingCodeKey,omitempty"`
	RoutingCodeValue string `json:"routingCodeValue,omitempty"`
}

type TransactionService interface {
	FetchTxByID(ctx context.Context, tx TxID) (Tx, error)
}

type EventStore interface {
	PublishEvent(context.Context, Event) error
}

type JSONTime struct {
	time.Time
}

const RFC3339Milli = "2006-01-02T15:04:05.999"

func (ct *JSONTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}
	ct.Time, err = time.Parse(RFC3339Milli, s)
	return
}

type RequestEvent struct {
	ID                           EventID              `json:"id,omitempty"`
	Type                         TxType               `json:"type,omitempty"`
	ShortID                      string               `json:"shortId,omitempty"`
	TransactionID                TxID                 `json:"transactionId,omitempty"`
	Status                       TxStatus             `json:"status,omitempty"`
	InternalAccountID            string               `json:"internalAccountId,omitempty"`
	PaymentType                  string               `json:"paymentType,omitempty"`
	SenderBic                    string               `json:"senderBic,omitempty"`
	SenderAccountNumber          string               `json:"senderAccountNumber,omitempty"`
	SenderRoutingCodes           []SenderRoutingCodes `json:"senderRoutingCodes,omitempty"`
	SenderIban                   string               `json:"senderIban,omitempty"`
	SenderName                   string               `json:"senderName,omitempty"`
	SenderAddress                string               `json:"senderAddress,omitempty"`
	SenderInformation            string               `json:"senderInformation,omitempty"`
	Amount                       RequestAmount        `json:"amount,omitempty"`
	TransactionDateTime          *JSONTime            `json:"transactionDateTime,omitempty"`
	TransactionReference         string               `json:"transactionReference,omitempty"`
	FailureReason                string               `json:"failureReason,omitempty"`
	BeneficiaryAccountHolderName string               `json:"beneficiaryAccountHolderName,omitempty"`
	AccountHolderID              string               `json:"accountHolderId,omitempty"`
	Source                       string               `json:"source,omitempty"`
	CreatedBy                    string               `json:"createdBy,omitempty"`
	Comment                      string               `json:"comment,omitempty"`
	FxRate                       json.Number          `json:"fxRate,omitempty"`
	MandateID                    string               `json:"mandateId,omitempty"`
	MidMarketRate                json.Number          `json:"midMarketRate,omitempty"`
	FixedSide                    string               `json:"fixedSide,omitempty"`
	TransactionCategory          string               `json:"transactionCategory,omitempty"`
	OriginalTxID                 string               `json:"originalTransactionId,omitempty"`
	SourceInfo                   *SourceInfo          `json:"sourceInfo,omitempty"`
	DestinationInfo              *DestinationInfo     `json:"destinationInfo,omitempty"`
	CreatedDate                  *JSONTime            `json:"createdDate,omitempty"`
	UpdatedDate                  *JSONTime            `json:"updatedDate,omitempty"`
	TotalAmount                  *RequestAmount       `json:"totalAmount,omitempty"`
	Fee                          *RequestAmount       `json:"fee,omitempty"`
	RunningBalance               *RequestAmount       `json:"runningBalance,omitempty"`
	BuyAmount                    *RequestAmount       `json:"buyAmount,omitempty"`
	ReferenceAmount              *RequestAmount       `json:"referenceAmount,omitempty"`
}

type TxResponse struct {
	ID                   TxID            `json:"id,omitempty"`
	ShortID              string          `json:"shortId,omitempty"`
	AccountID            string          `json:"accountId,omitempty"`
	CreatedDate          int64           `json:"createdDate,omitempty"`
	UpdatedDate          int64           `json:"updatedDate,omitempty"`
	PaymentDate          int64           `json:"paymentDate,omitempty"`
	EncID                string          `json:"encId,omitempty"`
	TransactionCategory  string          `json:"transactionCategory,omitempty"`
	PaymentType          string          `json:"paymentType,omitempty"`
	Type                 TxType          `json:"type,omitempty"`
	SourceInfo           SourceInfo      `json:"sourceInfo,omitempty"`
	DestinationInfo      DestinationInfo `json:"destinationInfo,omitempty"`
	Source               string          `json:"source,omitempty"`
	Destination          string          `json:"destination,omitempty"`
	Amount               Amount          `json:"amount,omitempty"`
	TotalAmount          Amount          `json:"totalAmount,omitempty"`
	Fee                  Amount          `json:"fee,omitempty"`
	RunningBalance       Amount          `json:"runningBalance,omitempty"`
	BuyAmount            Amount          `json:"buyAmount,omitempty"`
	FxRate               float64         `json:"fxRate,omitempty"`
	MidMarketRate        float64         `json:"midMarketRate,omitempty"`
	FixedSide            string          `json:"fixedSide,omitempty"`
	Status               TxStatus        `json:"status,omitempty"`
	FailureReason        string          `json:"failureReason,omitempty"`
	Comment              string          `json:"comment,omitempty"`
	TransactionReference string          `json:"transactionReference,omitempty"`
	ReferenceAmount      Amount          `json:"referenceAmount,omitempty"`
	AccountHolderID      string          `json:"accountHolderId,omitempty"`
}

type RequestAmount struct {
	Value    json.Number `json:"value,omitempty"`
	Currency Currency    `json:"currency,omitempty"`
}

func requestAmount(a *RequestAmount) Amount {
	if a == nil {
		return Amount{}
	}
	fa, _ := a.Value.Float64()
	amount := NewMoneyFromFloat(fa, a.Currency)
	return Amount{Value: amount.Amount(), Currency: amount.Currency()}
}
