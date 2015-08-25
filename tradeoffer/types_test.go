package tradeoffer

import (
	"testing"
)

func TestSteamTradeOfferID(t *testing.T) {
	var steamTradeOfferID SteamTradeOfferID
	steamTradeOfferID = 123
	expectedValue := "123"
	gotValue := steamTradeOfferID.String()
	if expectedValue != gotValue {
		t.Errorf(
			"SteamTradeOfferID.String(%v), expected %v, got %v",
			steamTradeOfferID.String(), expectedValue, gotValue,
		)

	}
}
