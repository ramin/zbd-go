//go:build e2e

package e2e

import (
	"fmt"
	"strings"
	"testing"
	"time"

	zebedee "github.com/zebedeeio/go-sdk"
)

func TestCharges(t *testing.T) {

	var internalID = "testx"
	var callbackURL = "https://example.com/callback"
	var amount = "123000"
	var description = "a test invoice"

	client := NewClient()
	charge, err := client.Charge(&zebedee.Charge{
		ExpiresIn:   30 * 60,
		Amount:      amount,
		Description: description,
		InternalID:  internalID,
		CallbackURL: callbackURL,
	})
	if err != nil {
		t.Errorf("got error from .Charge(): %s", err)
	} else {
		if charge.ExpiresAt.After(time.Now().Add(time.Minute * 30)) {
			t.Error("charge expires after we wanted")
		}
		if !strings.HasPrefix(charge.Invoice.Request, "lntbs123") {
			t.Error("charge has wrong bolt11 invoice")
		}
	}

	// fetch this same charge
	charge, err = client.GetCharge(charge.ID)
	if err != nil {
		t.Errorf("got error from .GetCharge(): %s", err)
	} else {
		if charge.Amount != amount {
			t.Error("charge amount is different than specified")
		}
		if charge.Description != description {
			t.Error("charge description is different than specified")
		}
		if charge.InternalID != internalID {
			t.Error("charge internal id is different than specified")
		}
		if charge.CallbackURL != callbackURL {
			t.Error("charge callback url is different than specified")
		}
	}
}

func TestChargesBad(t *testing.T) {
	client := NewClient()
	_, err := client.Charge(&zebedee.Charge{
		Amount:      "5000000000",
		Description: "a test invoice",
		InternalID:  "testb",
	})

	if err == nil {
		t.Errorf(".Charge() should have returned an error")
	} else if err.Error() != "The maximum Charge amount supported is 500,000 satoshis." {
		t.Errorf(".Charge() returned the wrong error %s", err.Error())
	}

	_, err = client.Charge(&zebedee.Charge{
		Amount:      "-120",
		Description: "a test invoice",
		InternalID:  "testb",
	})

	if err == nil {
		fmt.Println(err)
		t.Errorf(".Charge() should have returned an error")
	}
}
