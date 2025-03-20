package pepay

import (
	"fmt"
	"time"
)

// DepositService is used to invoke APIs related to transactions.
type DepositService service

type DepositID string

type DepositFiltersDto struct {
	QueryDate  string   `json:"queryDate"`
	DepositIds []string `json:"depositIds"`
	References []string `json:"references"`
	States     []string `json:"states"`
}

type ReceiveCryptoTransactionDto struct {
	TxHash        string          `json:"txHash"`
	Amount        Money           `json:"amount"`
	CreatedDate   string          `json:"createdDate"`
	ConfirmedDate string          `json:"confirmedDate"`
	State         string          `json:"state"`
	Rate          ExchangeRateDto `json:"rate"`
	TransactionId string          `json:"transactionId"`
	Confirmations int             `json:"confirmations"`
}

type DepositDtoV3 struct {
	DepositId                  PayoutID                      `json:"depositId"`
	AccountId                  string                        `json:"accountId"`
	CallbackUrl                string                        `json:"callbackUrl"`
	Reference                  string                        `json:"reference"`
	CustomerId                 string                        `json:"customerId"`
	RequestedAmountInFiat      Money                         `json:"requestedAmountInFiat"`
	RequestedAmountInCrypto    Money                         `json:"requestedAmountInCrypto"`
	ReceiverAddress            string                        `json:"receiverAddress"`
	VerificationSignature      string                        `json:"verificationSignature"`
	Network                    string                        `json:"network"`
	CryptoCurrency             string                        `json:"cryptoCurrency"`
	DepositState               string                        `json:"depositState"`
	CreatedDate                string                        `json:"createdDate"`
	ExpiryDate                 string                        `json:"expiryDate"`
	ReceiveOrderId             string                        `json:"receiveOrderId"`
	ReceiveOrderAlphanumericId string                        `json:"receiveOrderAlphanumericId"`
	FixedExchangeRate          ExchangeRateDto               `json:"fixedExchangeRate"`
	ReceivedFunds              []ReceiveCryptoTransactionDto `json:"receivedFunds"`
	DuplicatedFromDepositId    string                        `json:"duplicatedFromDepositId"`
}

type PageDepositDtoV3 struct {
	TotalPages       int            `json:"totalPages"`
	TotalElements    int64          `json:"totalElements"`
	First            bool           `json:"first"`
	Size             int            `json:"size"`
	Content          []DepositDtoV3 `json:"content"`
	Number           int            `json:"number"`
	Sort             SortObject     `json:"sort"`
	Pageable         PageableObject `json:"pageable"`
	NumberOfElements int            `json:"numberOfElements"`
	Last             bool           `json:"last"`
	Empty            bool           `json:"empty"`
}

type Deposit struct {
	Filters DepositFiltersDto `json:"filters"`
	Results PageDepositDtoV3  `json:"results"`
}

type CreateDepositRequest struct {
	AccountId       string `json:"accountId"`
	Nonce           int64  `json:"nonce"`
	Reference       string `json:"reference"`
	CustomerId      string `json:"customerId,omitempty"`
	CallbackUrl     string `json:"callbackUrl,omitempty"`
	ExpiryDate      string `json:"expiryDate,omitempty"`
	Network         string `json:"network"`
	RequestedAmount Money  `json:"requestedAmount"`
}

func (s *DepositService) Create(r CreateDepositRequest) (DepositID, error) {
	var v struct {
		DepositId DepositID `json:"depositId"`
	}

	req, err := s.client.NewRequest("POST", "/v3/deposits", r)
	if err != nil {
		return "", fmt.Errorf("create deposit http request has failed: %w", err)
	}
	_, err = s.client.Do(req, &v)
	if err != nil {
		return "", err
	}
	return v.DepositId, err
}

func (s *DepositService) Retrieve(id string) (*Deposit, error) {
	var v Deposit

	queryDate := time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	url := fmt.Sprintf("/v3/deposits?queryDate=%s&depositIds=%s", queryDate, id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("exchange rate http request has failed: %w", err)
	}
	_, err = s.client.Do(req, &v)
	if err != nil {
		return nil, err
	}
	return &v, err
}
