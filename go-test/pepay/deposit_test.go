package pepay

import (
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
			t.Errorf("got = %v", err)
		}
	})
}

func TestRetrieveDeposit(t *testing.T) {
	t.Run("retrieve", func(t *testing.T) {
		c := NewFakeClient(t, nil)
		depo, err := c.Deposit.Retrieve("3a5375a0-a4ee-4064-b3eb-afe7779e5909")
		fmt.Printf("%v", depo)
		if err != nil {
			t.Errorf("got = %v", err)
		}
	})
}
