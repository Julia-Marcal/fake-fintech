package services

import (
	"fmt"
	"io"
	"net/http"

	task "github.com/Julia-Marcal/fake-fintech/internal/domain"
)

func FetchPriceAlphaVantage(task task.AssetTask, apiKey string) ([]byte, error) {

	url := fmt.Sprintf("https://www.alphavantage.co/query?function=DIGITAL_CURRENCY_DAILY&symbol=%s&market=%s&apikey=%s", task.Symbol, task.Market, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from Alpha Vantage: %w", err)
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseData, nil
}
