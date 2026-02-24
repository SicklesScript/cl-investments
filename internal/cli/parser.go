package cli

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/SicklesScript/cl-investments/internal/database"
	//"github.com/SicklesScript/cl-investments/internal/database"
)

// Struct to parse holdings.csv file
type Holding struct {
	Ticker           string `csv:"Symbol"`
	Shares           string `csv:"Quantity"`
	TransactionPrice string `csv:"Avg. Price"`
}

/**/
func (s *State) ParseHoldingsCSV(filePath string) error {
	// Open csv file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)

	// Skip header row (title row)
	_, err = reader.Read()
	if err != nil {
		return err
	}

	// Loop over all records in struct to add to db
	for {
		// Get row
		record, err := reader.Read()
		// Break out of infinite loop if we reached EOF
		if err == io.EOF {
			break
		}
		// Else return err as normal
		if err != nil {
			return err
		}
		// Manually assign fields from the record to the struct
		holding := Holding{
			Ticker:           record[0],
			Shares:           record[2],
			TransactionPrice: record[3],
		}

		// Create transaction for each holding parsed from csv file
		_, err = s.DBQueries.AddTransaction(context.Background(), database.AddTransactionParams{
			Ticker:           holding.Ticker,
			TransactionDate:  time.Now(),
			TransactionPrice: holding.TransactionPrice,
			Shares:           holding.Shares,
			Type:             "BUY",
			Username:         s.CurrentUser,
		})
		if err != nil {
			return err
		}

	}
	// Print success message
	fmt.Println("Successfully loaded CSV holdings into database")
	return nil
}
