# CL-Investments

A powerful Command Line Interface (CLI) tool for tracking stock portfolios, analyzing dividend growth, and researching market data.

## Features

- **User Authentication**: Secure login and signup system using Argon2id password hashing.
- **Stock Research**: Fetch real-time stock data and metrics using the Alpha Vantage API.
- **Portfolio Management**:
    - Track transactions (Buy, Sell, Dividends).
    - Calculate cost basis and total returns.
    - Monitor portfolio performance.
- **Dividend Analysis**:
    - View detailed dividend data for individual stocks.
    - Analyze dividend growth for companies and ETFs.
    - Calculate weighted TTM (Trailing Twelve Months) dividend growth for your entire portfolio.
- **Bulk Import**: Upload holdings and transactions via CSV files.

## Tech Stack

- **Language**: [Go](https://go.dev/) (Golang)
- **Database**: [PostgreSQL](https://www.postgresql.org/)
- **ORM/Query Builder**: [sqlc](https://sqlc.dev/) for type-safe SQL.
- **Migrations**: [Goose](https://github.com/pressly/goose)
- **API**: [Alpha Vantage](https://www.alphavantage.co/)

## Getting Started

### Prerequisites

- Go 1.24+
- PostgreSQL
- Alpha Vantage API Key (Get one [here](https://www.alphavantage.co/support/#api-key))

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/SicklesScript/cl-investments.git
   cd cl-investments
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your environment variables:
   Create a `.env` file in the root directory:
   ```env
   DB_URL=postgres://username:password@localhost:5432/cl_investments?sslmode=disable
   API_KEY=your_alpha_vantage_api_key
   ```

4. Run database migrations:
   ```bash
   goose -dir sql/schema postgres "your_db_connection_string" up
   ```

### Running the Application

To start the CLI tracker:

```bash
go run cmd/tracker/main.go
```

## Usage

Once the application is running, you can use the following commands:

- `login <username> <password>`: Log in or create a new account.
- `research <ticker>`: Get current stock data for a ticker.
- `add <ticker> <shares> <price> <type> <date>`: Add a transaction (type: BUY, SELL, DIV; date: YYYY-MM-DD).
- `cost-basis`: View your current holdings and their cost basis.
- `return`: View individual and total portfolio returns.
- `div-data <ticker>`: View dividend information for a specific stock.
- `div-growth <ticker>`: Analyze dividend growth for a stock or ETF.
- `portfolio-growth`: View the weighted dividend growth of your entire portfolio.
- `upload <filepath>`: Import transactions from a CSV file.
- `exit`: Close the application.

## Project Structure

- `cmd/tracker/`: Main entry point for the CLI application.
- `internal/alphalogic/`: Logic for interacting with the Alpha Vantage API and performing financial calculations.
- `internal/auth/`: Authentication logic.
- `internal/cli/`: CLI input handling and command parsing.
- `internal/database/`: Generated Go code for database interactions (via sqlc).
- `sql/`: SQL schema migrations and queries.
