package service

import (
	"fmt"
	"io"
	"net/http"

	task "github.com/Julia-Marcal/fake-fintech/internal/domain/entity"
)

func FetchPriceCoinCap(task task.AssetTask, apiKey string) ([]byte, error) {
	url := fmt.Sprintf("https://rest.coincap.io/v3/assets/%s?apiKey=%s", task.Symbol, apiKey)

	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from CoinCap: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", response.StatusCode)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return responseData, nil
}
