package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SicklesScript/cl-investments/internal/alphalogic"
	"github.com/SicklesScript/cl-investments/internal/cli"
	"github.com/SicklesScript/cl-investments/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get apiKey from .env file
	apiKey := os.Getenv("API_KEY")

	// Make connection to DB
	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("unable to connect to db: %s", err)
	}
	// Get db queries
	queries := database.New(db)

	// Initalize state struct
	state := cli.State{
		DBQueries:   queries,
		CurrentUser: "",
	}

	// Initalize StockData struct
	var stockData alphalogic.StockData
	// Initalize DividendData struct
	var dividendData alphalogic.DividendData

	// Print commands for user
	cli.PrintCommands()

	for {
		// Extract user entry as slice of words
		words := cli.GetInput()

		// Switch series to direct flow of commands
		switch words[0] {
		case "research":
			ticker := words[1]
			// Load stock data struct
			err = stockData.GetStockData(ticker, apiKey)
			if err != nil {
				fmt.Printf("Error retreiving stock data: %s", err)
			} else {
				// Display stock data to terminal
				stockData.DisplayStockData()
			}
		case "login":
			// Get username and password from user input
			username := words[1]
			password := words[2]
			// If user exists, login, else signup and then login
			err := state.LoginOrSignup(username, password)
			if err != nil {
				fmt.Printf("error in login/signup: %s\n", err)
			}
		case "add":
			// Store user arsg
			ticker := words[1]
			shares := words[2]
			price := words[3]
			transType := words[4]
			date := words[5]

			layout := "2006-01-02"
			t, _ := time.Parse(layout, date)
			// Create transaction and provide user with receipt
			err := state.AddTransaction(ticker, shares, price, transType, state.CurrentUser, t)
			if err != nil {
				fmt.Printf("error adding transaction: %s\n", err)
			}
		case "cost-basis":
			// Get total cost value for user
			err := state.GetHoldings(state.CurrentUser)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		case "return":
			// Get total return for each stock in portfolio
			err := state.GetTotalReturn(state.CurrentUser)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		case "dividend":
			ticker := words[1]
			// Get dividend data for stock
			err := dividendData.GetDividendData(ticker, apiKey)
			if err != nil {
				fmt.Printf("error: %s", err)
			}
			// Print div data to console
			dividendData.DisplayDividendData()
		case "growth":
			ticker := words[1]
			// Get dividend growth data of etf or company
			err := dividendData.GetDividendGrowth(ticker, apiKey)
			if err != nil {
				fmt.Printf("error: %s", err)
			}
		case "upload":
			// Get filepath arg from user
			filePath := words[1]
			// Parse csv and print data to console
			err := state.ParseHoldingsCSV(filePath)
			if err != nil {
				fmt.Printf("error: %s\n", err)
			}
		case "exit":
			return
		default:
			fmt.Println("----------------------------")
			fmt.Println("Unknown command. Following is a list of commands: ")
			cli.PrintCommands()
		}
	}
}
