package alphalogic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/SicklesScript/cl-investments/internal/database"
)

/*
Makes http request to alpha vantage dividends api
Loads dividneds data into DividendData struct
*/
func (dd *DividendData) GetDividendData(ticker string, apiKey string) error {
	// Generate full url
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=DIVIDENDS&symbol=%s&apikey=%s", ticker, apiKey)

	// Create custom client with 10 second timeout feature
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	// Get response from alphavantage
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Stream and decode body into struct
	if err := json.NewDecoder(resp.Body).Decode(&dd); err != nil {
		return err
	}
	// Return stock data
	return nil
}

/*
Makes http request to alpha vantage dividends api
Loads dividneds data into DividendData struct
*/
func (dd *DividendData) GetDividendGrowth(ticker string, apiKey string) error {
	// Load dd struct
	err := dd.GetDividendData(ticker, apiKey)
	if err != nil {
		return err
	}
	// Pass shares (as string for parsing) as 1 to simulate single share
	shares := "1"
	ttmDiv, prevDiv := dd.CalculateDiv(shares)

	// Calculate div growth
	divGrowth := ((ttmDiv - prevDiv) / prevDiv) * 100

	// Print div growth to console
	fmt.Printf("TTM Dividend: %.2f\n", ttmDiv)
	fmt.Printf("Prev Dividend: %.2f\n", prevDiv)
	fmt.Printf("TTM Dividend Growth: %.2f%%\n", divGrowth)
	return nil
}

/*
Print all Dividend data to terminal
*/
func (dd *DividendData) DisplayDividendData() {
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Name: %s\n", dd.Symbol)
	// Initalize variable to hold total TTM Div
	var TTMDiv float64
	// Print last 4 dividend payments to console
	for i := range 8 {
		fmt.Printf("TTM Div Payments: %s\n", dd.DivData[i].Amount)
		// Parse string div data to float64
		div, err := strconv.ParseFloat(dd.DivData[i].Amount, 64)
		if err != nil {
			fmt.Printf("error parsing div info: %s\n", err)
		}
		// Load TTM div
		TTMDiv += div
	}
	// Print total TTM div
	fmt.Printf("TTM Total Dividend: %.2f\n", TTMDiv)
	fmt.Printf("Next Dividend Date: %s\n", dd.DivData[0].ExDivDate)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
}

/*
Calculate weighted dividend growth of entire portfolio
*/
func (dd *DividendData) GetPortfolioDividendGrowth(apiKey string, holdings []database.Transaction) error {
	// Div variables for running totals
	var TTMDiv float64
	var PrevDiv float64

	// Loop over holdings list to get all data
	for _, holding := range holdings {
		// Reset struct incase it is causing error

		// Get dividend data
		err := dd.GetDividendData(holding.Ticker, apiKey)
		if err != nil {
			return err
		}
		// Calculate ttm and prev div for holding
		ttmDiv, prevDiv := dd.CalculateDiv(holding.Shares)

		// Manually account for schd stock split
		if strings.ToLower(holding.Ticker) == "schd" {
			prevDiv = 0.99437 * 506.01787
		}
		// Manually account for avgo stock split
		if strings.ToLower(holding.Ticker) == "avgo" {
			prevDiv = 2.17 * 0.61649
		}
		// Logging
		fmt.Printf("Ticker: %s, TTM Div: %.2f, Prev Div: %.2f\n", holding.Ticker, ttmDiv, prevDiv)
		// Accumulate total
		TTMDiv += ttmDiv
		PrevDiv += prevDiv

		// Add delay to avoid api bottleneck
		time.Sleep(5 * time.Second)
	}

	// Calculate div growth
	divGrowth := ((TTMDiv - PrevDiv) / PrevDiv) * 100

	// Print div growth to console
	fmt.Printf("TTM Dividend: %.2f\n", TTMDiv)
	fmt.Printf("Prev Dividend: %.2f\n", PrevDiv)
	fmt.Printf("TTM Dividend Growth: %.2f\n", divGrowth)
	return nil
}
