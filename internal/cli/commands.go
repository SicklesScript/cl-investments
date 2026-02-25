package cli

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/SicklesScript/cl-investments/internal/auth"
	"github.com/SicklesScript/cl-investments/internal/database"
)

/*
Combines the signup and login logic to reduce
amount of commands

If user already has an account, they are logged in
If not, the user account is created and logged in
*/
func (s *State) LoginOrSignup(username, pass string) error {
	// Check if user exists
	user, err := s.DBQueries.GetUser(context.Background(), username)
	if errors.Is(err, sql.ErrNoRows) {
		// Hash password for acct creation
		hash, err := auth.HashPassword(pass)
		if err != nil {
			return err
		}
		// Create account
		s.DBQueries.CreateUser(context.Background(), database.CreateUserParams{
			Name:           username,
			HashedPassword: hash,
		})
		// Set new user to current user
		s.CurrentUser = username
		// Print welcome message
		fmt.Println("----------------------------")
		fmt.Printf("Welcome to CL-Investments %s!\n", s.CurrentUser)
		fmt.Println("----------------------------")
		return nil
	}
	// If user already exists, check that info matches db
	ok, err := auth.CheckPasswordHash(pass, user.HashedPassword)
	if err != nil {
		return err
	}
	// If password and username match, login user
	if ok {
		s.CurrentUser = user.Name
		// Print welcome message
		fmt.Println("----------------------------")
		fmt.Printf("Welcome Back to CL-Investments %s!\n", s.CurrentUser)
		fmt.Println("----------------------------")
		return nil
	} else {
		return errors.New("invalid username or password")
	}
}

/*
Handles logic to add transaction
Will check that transaction is possible given type
I.E. Unable to sell 5 shares if user only has 3.3 shares
*/
func (s *State) AddTransaction(ticker, shares, price, transType, username string, transDate time.Time) error {
	// Convert shares string to float, will need to change this in db later
	transShares, _ := strconv.ParseFloat(shares, 64)
	// Check that user has enough shares to make sell
	currShares, err := s.DBQueries.GetHolding(context.Background(), database.GetHoldingParams{
		Ticker:   ticker,
		Username: username,
	})
	if err != nil {
		return err
	}
	// If user tries to sell more shares than they hold, return error
	if transType == "SELL" && currShares < transShares {
		return errors.New("attempting to sell more shares than you hold")
	}
	// Create transaction
	s.DBQueries.AddTransaction(context.Background(), database.AddTransactionParams{
		Ticker:           ticker,
		Shares:           shares,
		TransactionPrice: price,
		TransactionDate:  transDate,
		Type:             transType,
		Username:         username,
	})
	// Print success message to terminal
	fmt.Printf("Successfully completed %s of %s\n", transType, ticker)
	// Get new holding info for company following successuful transaction
	// This might be a bad performance decision that will need to be analyzed later
	newHolding, err := s.DBQueries.GetHolding(context.Background(), database.GetHoldingParams{
		Ticker:   ticker,
		Username: username,
	})
	if err != nil {
		return err
	}
	// Print updated holding info to terminal
	fmt.Printf("Updated share count for %s: %.4f\n", ticker, newHolding)
	return nil
}

/*
Get holding value for entire portfolio
*/
func (s *State) GetHoldings(username string) error {
	// Get total value of all holdings for user
	total, err := s.DBQueries.GetHoldings(context.Background(), username)
	if err != nil {
		return err
	}
	// Print holding value to terminal
	fmt.Printf("Cost basis of Portfolio: $%.2f\n", total)
	return nil
}

/*
Get total return for each holding in portfolio
*/
func (s *State) GetTotalReturn(username string) error {
	// Get return data for each individual holding
	holdings, err := s.DBQueries.GetReturn(context.Background(), username)
	if err != nil {
		return err
	}

	// Variables for storing totals
	var totalProfit float64
	var totalVal float64
	var totalCostBasis float64

	// Iterate over holdings and print return for each ticker
	for _, holding := range holdings {
		// Placeholder num for current share price since api does not currently contain that info
		marketVal := holding.CurrentShares * 540.00
		totalVal += marketVal // Get total market value for entire port

		totalCostBasis += holding.CostBasis // Get total cost basis for entire port

		tickerProfit := marketVal - holding.CostBasis
		totalProfit += tickerProfit // Get total profit for entire port

		totalReturn := (tickerProfit / holding.CostBasis) * 100
		fmt.Printf("total return for %s: %.2f%%\n", holding.Ticker, totalReturn)
	}
	// Get total portfolio return
	portfolioReturn := (totalProfit / totalCostBasis) * 100
	fmt.Printf("total return for portfolio: %.2f%%\n", portfolioReturn)
	return nil
}

/*
This architecture might be bad.

Currently need this function to create list of user holdings to pass off to my alphalogic package
to calculate whole portfolio dividend growth
*/
func (s *State) GetAllHoldings(username string) ([]database.Transaction, error) {
	// Get all holdings for user
	holdings, err := s.DBQueries.GetAll(context.Background(), username)
	if err != nil {
		return []database.Transaction{}, err
	}
	// Return holdings for use in dividend growth calc
	return holdings, nil
}
