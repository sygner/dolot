package constants

import (
	"dolott_user_gw_http/internal/types"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var MAIN_LUNC_WALLET_ADDRESS = "terra1jqt929u9tp2q6s9a79k9jx3zxec097z759e267"
var MAIN_LUNC_USER_WALLET_ID int32 = 0
var MAIN_LUNC_WALLET_ID int32 = 0

type CoinPaprikaResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int    `json:"rank"`
	Quotes struct {
		USD struct {
			Price float64 `json:"price"`
		} `json:"USD"`
	} `json:"quotes"`
}

func GetLUNCPriceCoinPaprika() (float64, *types.Error) {
	url := "https://api.coinpaprika.com/v1/tickers/luna-terra"

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return 0, types.NewInternalError(fmt.Sprintf("failed to fetch LUNC price from CoinPaprika: %v", err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, types.NewInternalError(fmt.Sprintf("unexpected status code from CoinPaprika: %d", resp.StatusCode))
	}

	var cpResp CoinPaprikaResponse
	if err := json.NewDecoder(resp.Body).Decode(&cpResp); err != nil {
		return 0, types.NewInternalError(fmt.Sprintf("failed to decode CoinPaprika response: %v", err))
	}

	fmt.Println(cpResp)
	return cpResp.Quotes.USD.Price, nil
}
