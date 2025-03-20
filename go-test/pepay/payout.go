package pepay

import (
	"fmt"
	"time"
)

// PayoutService is used to invoke APIs related to transactions.
type PayoutService service

type PayoutID string
type PayoutStatus int8

func (s *PayoutStatus) UnmarshalText(text []byte) error {
	switch text := string(text); text {
	case "REQUESTED":
		*s = PayoutStatusRequested
	case "CREATE_FAILED":
		*s = PayoutStatusCreateFailed
	case "CREATE_TIMEOUT":
		*s = PayoutStatusCreateTimeout
	case "CREATED":
		*s = PayoutStatusCreated
	case "COSIGNER_AUTHORIZATIONS_PENDING":
		*s = PayoutStatusCosignerAuthPending
	case "COSIGNER_AUTHORIZED":
		*s = PayoutStatusCosignerAuthorized
	case "CANCELLED":
		*s = PayoutStatusCancelled
	case "AUTHORIZED":
		*s = PayoutStatusAuthorized
	case "SIGNING":
		*s = PayoutStatusSigning
	case "SENDING":
		*s = PayoutStatusSending
	case "SENT":
		*s = PayoutStatusSent
	case "COMPLETED":
		*s = PayoutStatusCompleted
	default:
		if len(text) > 0 {
			return fmt.Errorf("invalid withdrawal status value: %s", text)
		}
	}
	return nil
}

const (
	PayoutStatusRequested PayoutStatus = iota + 1
	PayoutStatusCreateFailed
	PayoutStatusCreateTimeout
	PayoutStatusCreated
	PayoutStatusCosignerAuthPending
	PayoutStatusCosignerAuthorized
	PayoutStatusCancelled
	PayoutStatusAuthorized
	PayoutStatusSigning
	PayoutStatusSending
	PayoutStatusSent
	PayoutStatusCompleted
)

type SourceType int8

func (s *SourceType) UnmarshalText(text []byte) error {
	switch text := string(text); text {
	case "ACCOUNT":
		*s = SourceTypeAccount
	case "WALLET":
		*s = SourceTypeWallet
	default:
		if len(text) > 0 {
			return fmt.Errorf("invalid source type value: %s", text)
		}
	}
	return nil
}

func (s SourceType) MarshalText() (text []byte, err error) {
	switch s {
	case SourceTypeAccount:
		return []byte("ACCOUNT"), nil
	case SourceTypeWallet:
		return []byte("WALLET"), nil
	default:
		return nil, fmt.Errorf("invalid source type: %d", s)
	}
}

const (
	SourceTypeAccount SourceType = iota + 1
	SourceTypeWallet
)

type PayoutFiltersDto struct {
	QueryDate  string   `json:"queryDate"`
	PayoutIds  []string `json:"payoutIds"`
	References []string `json:"references"`
	States     []string `json:"states"`
}

type FortrisCurrency struct {
	CurrencyCode          string `json:"currencyCode"`
	DefaultFractionDigits int    `json:"defaultFractionDigits"`
	NumericCode           int    `json:"numericCode"`
	NumericCodeAsString   string `json:"numericCodeAsString"`
	DisplayName           string `json:"displayName"`
	Symbol                string `json:"symbol"`
}

type ExchangeRateDto struct {
	Rate         float64 `json:"rate"`
	FromCurrency string  `json:"fromCurrency"`
	ToCurrency   string  `json:"toCurrency"`
	MeasuredDate string  `json:"measuredDate"`
}

type SentCryptoTransactionDto struct {
	TxHash        string          `json:"txHash"`
	Amount        Money           `json:"amount"`
	Fee           Money           `json:"fee"`
	CreatedDate   string          `json:"createdDate"`
	ConfirmedDate string          `json:"confirmedDate"`
	State         string          `json:"state"`
	Rate          ExchangeRateDto `json:"rate"`
	Confirmations int             `json:"confirmations"`
	TransactionId string          `json:"transactionId"`
}

type CosignerAuthorizationDto struct {
	State         string `json:"state"`
	UserGroupName string `json:"userGroupName"`
	Username      string `json:"username"`
	Expire        string `json:"expire"`
}

type AuthorizationDto struct {
	Id                     string                     `json:"id"`
	State                  string                     `json:"state"`
	CosignerAuthorizations []CosignerAuthorizationDto `json:"cosignerAuthorizations"`
}

type PayoutDtoV3 struct {
	PayoutId                  PayoutID                   `json:"payoutId"`
	AccountId                 string                     `json:"accountId"`
	ClientId                  string                     `json:"clientId"`
	CallbackUrl               string                     `json:"callbackUrl"`
	Reference                 string                     `json:"reference"`
	CustomerId                string                     `json:"customerId"`
	RequestedAmount           Money                      `json:"requestedAmount"`
	Network                   string                     `json:"network"`
	CryptoCurrency            string                     `json:"cryptoCurrency"`
	CreatedDate               string                     `json:"createdDate"`
	DestinationAddress        string                     `json:"destinationAddress"`
	DestinationAccountId      string                     `json:"destinationAccountId"`
	PayoutState               PayoutStatus               `json:"payoutState"`
	SendOrderId               string                     `json:"sendOrderId"`
	SendOrderAlphanumericId   string                     `json:"sendOrderAlphanumericId"`
	SentFunds                 []SentCryptoTransactionDto `json:"sentFunds"`
	Authorization             AuthorizationDto           `json:"authorization"`
	FeePolicy                 string                     `json:"feePolicy"`
	CustomFeeRate             int                        `json:"customFeeRate"`
	SubtractFee               bool                       `json:"subtractFee"`
	UseCoinConsolidation      bool                       `json:"useCoinConsolidation"`
	OrderCodeName             string                     `json:"orderCodeName"`
	GeneralLedgerName         string                     `json:"generalLedgerName"`
	GeneralLedgerProject      string                     `json:"generalLedgerProject"`
	GeneralLedgerDistribution string                     `json:"generalLedgerDistribution"`
}

type SortObject struct {
	// Direction    string `json:"direction"`
	// NullHandling string `json:"nullHandling"`
	// Ascending    bool   `json:"ascending"`
	// Property     string `json:"property"`
	// IgnoreCase   bool   `json:"ignoreCase"`
	Empty    bool `json:"empty"`
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}

type PageableObject struct {
	Offset     int64      `json:"offset"`
	Sort       SortObject `json:"sort"`
	PageNumber int        `json:"pageNumber"`
	PageSize   int        `json:"pageSize"`
	Paged      bool       `json:"paged"`
	Unpaged    bool       `json:"unpaged"`
}

type PagePayoutDtoV3 struct {
	TotalPages       int            `json:"totalPages"`
	TotalElements    int64          `json:"totalElements"`
	First            bool           `json:"first"`
	Size             int            `json:"size"`
	Content          []PayoutDtoV3  `json:"content"`
	Number           int            `json:"number"`
	Sort             SortObject     `json:"sort"`
	Pageable         PageableObject `json:"pageable"`
	NumberOfElements int            `json:"numberOfElements"`
	Last             bool           `json:"last"`
	Empty            bool           `json:"empty"`
}

type Payout struct {
	Filters PayoutFiltersDto `json:"filters"`
	Results PagePayoutDtoV3  `json:"results"`
}

type Money struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
}

type CreatePayoutRequest struct {
	Username                  string `json:"username"`
	OtpCode                   string `json:"otpCode,omitempty"`
	AccountId                 string `json:"accountId"`
	Reference                 string `json:"reference"`
	CustomerId                string `json:"customerId,omitempty"`
	CallbackUrl               string `json:"callbackUrl"`
	RequestedAmount           Money  `json:"requestedAmount"`
	DestinationAddress        string `json:"destinationAddress"`
	DestinationAccountId      string `json:"destinationAccountId,omitempty"`
	Network                   string `json:"network"`
	VerifyBalance             bool   `json:"verifyBalance,omitempty"`
	FeePolicy                 string `json:"feePolicy,omitempty"`
	CustomFeeRate             int    `json:"customFeeRate,omitempty"`
	SubtractFee               bool   `json:"subtractFee,omitempty"`
	UseCoinConsolidation      bool   `json:"useCoinConsolidation,omitempty"`
	OrderCodeName             string `json:"orderCodeName,omitempty"`
	GeneralLedgerName         string `json:"generalLedgerName,omitempty"`
	GeneralLedgerProject      string `json:"generalLedgerProject,omitempty"`
	GeneralLedgerDistribution string `json:"generalLedgerDistribution,omitempty"`
	Nonce                     int64  `json:"nonce"`
}

func (s *PayoutService) Create(authorize bool, r CreatePayoutRequest) (PayoutID, error) {
	var v struct {
		PayoutId PayoutID `json:"payoutId"`
	}

	req, err := s.client.NewRequest("POST", fmt.Sprintf("/v3/payouts?authorize=%v", authorize), r)
	if err != nil {
		return "", fmt.Errorf("create payout http request has failed: %w", err)
	}
	_, err = s.client.Do(req, &v)
	if err != nil {
		return "", err
	}
	w := v.PayoutId
	return w, err
}

func (s *PayoutService) Retrieve(id PayoutID) (*Payout, error) {
	var v Payout

	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/payouts?queryDate=%s&accountIds=%s", queryDate, id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if _, err := s.client.Do(req, &v); err != nil {
		return nil, err
	}
	return &v, nil
}
