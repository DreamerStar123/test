package pepay

import (
	"testing"
)

func TestCreatePayout(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		c := NewFakeClient(t, nil)
		in := CreatePayoutRequest{
			Username:    "david.ramirez",
			AccountId:   "f8620c81-d75d-4eff-b2ab-1cb785fbbaaf",
			Reference:   "ref",
			CallbackUrl: "https://psp.stg.01123581.com",
			RequestedAmount: Money{
				Currency: "USDT",
				Amount:   1,
			},
			DestinationAddress: "0xB5A2e8C0F3c7D6E9b1C4D5F7A8eF0A6C9F4C2E3D",
			Network:            "ethereum",
			VerifyBalance:      true,
			FeePolicy:          "SLOW",
			Nonce:              1090,
		}
		_, err := c.Payout.Create(false, in)
		if err != nil {
			t.Errorf("got = %v", err)
		}
	})
}

func TestRetrievePayout(t *testing.T) {
	t.Run("retrieve", func(t *testing.T) {
		c := NewFakeClient(t, nil)
		_, err := c.Payout.Retrieve("")
		if err != nil {
			t.Errorf("got = %v", err)
		}
	})
}
