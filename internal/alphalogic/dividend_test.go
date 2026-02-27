package alphalogic_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/SicklesScript/cl-investments/internal/alphalogic"
)

/*
TestGetDividendData provides a baseline unit test for the JSON decoding logic
using a hardcoded mock response to avoid external API calls.
*/
func TestGetDividendData(t *testing.T) {
	// Initalize dividned data struct to test filling mechanism
	data := alphalogic.DividendData{}

	// Sample json response
	reader := strings.NewReader(`{
		"symbol": "IBM",
		"data": [
			{
				"ex_dividend_date": "2026-02-10",
				"declaration_date": "2026-01-28",
				"record_date": "2026-02-10",
				"payment_date": "2026-03-10",
				"amount": "1.68"
			},
			{
				"ex_dividend_date": "2025-11-10",
				"declaration_date": "2025-10-22",
				"record_date": "2025-11-10",
				"payment_date": "2025-12-10",
				"amount": "1.68"
			}
		]
	}`)
	// Decode sample json data into dd struct
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		t.Fatalf("Error during decoding: %v", err)
	}

	// Check if symbol filled correctly
	expectedSymbol := "IBM"
	if data.Symbol != expectedSymbol {
		t.Errorf("Symbol mismatch: expected %s, got %s", expectedSymbol, data.Symbol)
	}
	// Check if DivData inner slice filled correctly
	expectedCount := 2
	if len(data.DivData) != expectedCount {
		t.Errorf("Data count mismatch: expected %d, got %d", expectedCount, len(data.DivData))
	}
}
