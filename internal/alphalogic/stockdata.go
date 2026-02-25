package alphalogic

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/SicklesScript/cl-investments/internal/database"
)

/*
Makes http request to alpha vantage overview api
Loads data into StockData struct
*/
func (sd *StockData) GetStockData(ticker string, apiKey string) error {
	// Generate full url
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=OVERVIEW&symbol=%s&apikey=%s", ticker, apiKey)

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
	if err := json.NewDecoder(resp.Body).Decode(&sd); err != nil {
		return err
	}
	// Return stock data
	return nil
}

/*
Print all stock data to terminal, separated by data type
*/
func (sd *StockData) DisplayStockData() {
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Name: %s\n", sd.Name)
	fmt.Printf("Description: %s\n", sd.Description)
	fmt.Printf("Sector: %s\n", sd.Sector)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("TTM PE: %s\n", sd.TTMPE)
	fmt.Printf("FWD PE: %s\n", sd.FWDPE)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("PEG Ratio: %s\n", sd.PriceToEarningsGrowth)
	fmt.Printf("Revenue Growth YOY: %s\n", sd.YOYRevenueGrowth)
	fmt.Printf("EPS Growth YOY: %s\n", sd.YOYEarningsGrowth)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Return on Equity: %s\n", sd.ROE)
	fmt.Printf("Return on Assets: %s\n", sd.ROA)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Dividend Yield: %s\n", sd.DivYield)
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Price Target: %s\n", sd.PriceTarget)
}

/*
Get total return for each holding in portfolio
*/
func (gqr *GlobalQuoteResponse) GetTotalReturn(username, apiKey string, holdings []database.GetReturnRow) error {
	// Variables for storing totals
	var totalProfit float64
	var totalVal float64
	var totalCostBasis float64

	// Iterate over holdings and print return for each ticker
	for _, holding := range holdings {
		// Get current price of holding
		currPrice, err := gqr.GetCurrentPrice(holding.Ticker, apiKey)
		if err != nil {
			return err
		}
		// Placeholder num for current share price since api does not currently contain that info
		marketVal := holding.CurrentShares * currPrice
		totalVal += marketVal // Get total market value for entire port

		totalCostBasis += holding.CostBasis // Get total cost basis for entire port

		tickerProfit := marketVal - holding.CostBasis
		totalProfit += tickerProfit // Get total profit for entire port

		totalReturn := (tickerProfit / holding.CostBasis) * 100
		fmt.Printf("total return for %s: %.2f%%\n", holding.Ticker, totalReturn)

		// Sleep to avoid messing up api
		time.Sleep(5 * time.Second)
	}
	// Get total portfolio return
	portfolioReturn := (totalProfit / totalCostBasis) * 100
	fmt.Printf("total return for portfolio: %.2f%%\n", portfolioReturn)
	return nil
}
