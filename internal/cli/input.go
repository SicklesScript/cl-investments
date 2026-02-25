package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() []string {
	fmt.Printf("> ")
	// Create scanner to read from cl
	scanner := bufio.NewScanner(os.Stdin)
	// Get next input
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}
	// Clean line of leading/trailing whitespaces
	line := CleanInput(scanner.Text())
	// Get all words from input
	lines := strings.Fields(line)

	return lines
}

func CleanInput(s string) string {
	// Remove whitespaces from input
	trimmed := strings.TrimSpace(s)

	return trimmed
}

func PrintCommands() {
	fmt.Println("COMMANDS")
	fmt.Println("----------------------------")
	fmt.Println("* login")
	fmt.Println("Sign up or login to your account")
	fmt.Println("Usage: signup <username> <password>")
	fmt.Println("----------------------------")
	fmt.Println("* research")
	fmt.Println("Get individual stock data")
	fmt.Println("Usage: research <stock>")
	fmt.Println("----------------------------")
	fmt.Println("* add")
	fmt.Println("Add holding to your portfolio")
	fmt.Println("Usage: add <stock> <shares> <pps> <BUY, SELL, DIV> <yyyy-mm-dd>")
	fmt.Println("----------------------------")
	fmt.Println("* return")
	fmt.Println("Generate total return report for each holding")
	fmt.Println("Usage: return")
	fmt.Println("----------------------------")
	fmt.Println("* div-data")
	fmt.Println("Generate TTM dividend data for company")
	fmt.Println("Usage: div-data <ticker>")
	fmt.Println("----------------------------")
	fmt.Println("* div-growth")
	fmt.Println("Generate TTM dividend growth for company")
	fmt.Println("Usage: div-growth <ticker>")
	fmt.Println("----------------------------")
	fmt.Println("* portfolio-growth")
	fmt.Println("Generate TTM dividend growth for entire portfolio")
	fmt.Println("Usage: portfolio-growth")
	fmt.Println("----------------------------")
	fmt.Println("* upload")
	fmt.Println("Upload CSV file of current holdings")
	fmt.Println("Usage: upload <file path>")
	fmt.Println("----------------------------")
	fmt.Println("* exit")
	fmt.Println("Close application")
	fmt.Println("Usage: exit")
	fmt.Println("----------------------------")
}
