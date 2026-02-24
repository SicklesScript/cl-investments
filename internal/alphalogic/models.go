package alphalogic

type StockData struct {
	Symbol                string `json:"Symbol"`
	Name                  string `json:"Name"`
	Description           string `json:"Description"`
	Sector                string `json:"Sector"`
	TTMPE                 string `json:"TrailingPE"`
	FWDPE                 string `json:"ForwardPE"`
	PriceToEarningsGrowth string `json:"PEGRatio"`
	DivYield              string `json:"DividendYield"`
	PriceTarget           string `json:"AnalystTargetPrice"`
	YOYEarningsGrowth     string `json:"QuarterlyEarningsGrowthYOY"`
	YOYRevenueGrowth      string `json:"QuarterlyRevenueGrowthYOY"`
	ROA                   string `json:"ReturnOnAssetsTTM"`
	ROE                   string `json:"ReturnOnEquityTTM"`
}

type DividendData struct {
	Symbol  string `json:"symbol"`
	DivData []Data `json:"data"`
}

type Data struct {
	ExDivDate string `json:"ex_dividend_date"`
	Amount    string `json:"amount"`
}
