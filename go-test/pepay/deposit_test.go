package pepay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateDeposit(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		c := NewFakeClient(t, nil)
		in := CreateDepositRequest{
			AccountId:   "f8620c81-d75d-4eff-b2ab-1cb785fbbaaf",
			Reference:   "ref",
			CallbackUrl: "https://psp.stg.01123581.com",
			RequestedAmount: Money{
				Currency: "USDT",
				Amount:   1,
			},
			Nonce: 1096,
		}
		depoId, err := c.Deposit.Create(in)
		println(depoId)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestRetrieveDeposit(t *testing.T) {
	t.Run("retrieve", func(t *testing.T) {
		c := NewFakeClient(t, nil)
		depo, err := c.Deposit.Retrieve("b5f4d9b0-6f1d-4c1c-849e-1d7058751635")
		content, err := json.Marshal(depo)
		fmt.Printf("%s", content)
		if err != nil {
			t.Errorf("got = %v", err)
		}
	})
}
