package alphalogic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// Helpter function for accumulating ttm and prev dividend
func (dd *DividendData) CalculateDiv(shares string) (float64, float64) {
	var ttmDiv, prevDiv float64

	// Get TTM Div total
	for i := range 4 {
		// Parse div and shares to get weighted div
		div, _ := strconv.ParseFloat(dd.DivData[i].Amount, 64)
		s, _ := strconv.ParseFloat(shares, 64)

		// Total div factoring in shares
		weightedDiv := div * s
		ttmDiv += weightedDiv
	}

	// Get 1 year prior to TTM div total
	for i := 4; i <= 7; i++ {
		// Parse div and shares to get total weighted div
		div, _ := strconv.ParseFloat(dd.DivData[i].Amount, 64)
		s, _ := strconv.ParseFloat(shares, 64)

		// Total div factoring in shares
		weightedDiv := div * s
		prevDiv += weightedDiv
	}
	return ttmDiv, prevDiv
}

// Helper function for getting current stock price
func (gqr *GlobalQuoteResponse) GetCurrentPrice(ticker, apiKey string) (float64, error) {
	// Generate URL for making api req for price data
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", ticker, apiKey)

	// Create custom client with 10 second timeout feature
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// Get response from alphavantage
	resp, err := client.Get(url)
	if err != nil {
		return 0.0, err
	}
	defer resp.Body.Close()

	// Stream and decode body into struct
	if err := json.NewDecoder(resp.Body).Decode(&gqr); err != nil {
		return 0.0, err
	}
	// Parse price string to float
	currPrice, err := strconv.ParseFloat(gqr.Quote.Price, 64)
	if err != nil {
		return 0.0, err
	}
	// Return current stock price
	return currPrice, nil
}
