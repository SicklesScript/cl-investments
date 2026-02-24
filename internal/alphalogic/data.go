package alphalogic

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
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
Print all Dividend data to terminal
*/
func (dd *DividendData) DisplayDividendData() {
	fmt.Println("---------------------------------------------")
	fmt.Println("---------------------------------------------")
	fmt.Printf("Name: %s\n", dd.Symbol)
	// Initalize variable to hold total TTM Div
	var TTMDiv float64
	// Print last 4 dividend payments to console
	for i := range 4 {
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
