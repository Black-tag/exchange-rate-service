package repository

import (
	"exchange-rate-service/internal/config"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)





func FetchRates() (io.ReadCloser, error) {

	
	logger := slog.Default().With(
		"handler", "FetchRates",
	)
	
	
	 
	
	
	url := fmt.Sprintf("%s/live?access_key=%s", config.BaseURL, config.APIKey)
	logger = logger.With("url", url)
	logger.Info("making request to the rate excahnge api")

	httpClient := &http.Client{Timeout: 5* time.Second}

	resp, err := httpClient.Get(url)
	if err != nil {
		logger.Error("couldnt fetch response", "error", err)
		return nil, err

	}
	
	logger.Info("successful request")
	
	return resp.Body, nil
}


func FetchHistoricalData(date string) (io.ReadCloser, error) {
	logger := slog.Default().With(
		"function", "FetchHistoricalData",
	)
	logger =logger.With("date", date)
	url := fmt.Sprintf("%s/historical?access_key=%s&date=%s", config.BaseURL, config.APIKey, date)
	logger = logger.With("url",url)
	logger.Info("making request to the rate excahnge api")

	httpClient := &http.Client{Timeout: 5* time.Second}

	resp, err := httpClient.Get(url)
	if err != nil {
		logger.Error("couldnt fetch response", "error", err)
		return nil, err

	}
	
	logger.Info("successful request")
	
	return resp.Body, nil

}

func ConvertAmount(from string, to string, amount int) (io.ReadCloser, error) {
	logger := slog.Default().With(
		"function", "FetchHistoricalData",
	)
	
	base := strings.ToUpper(from)
	symbol := strings.ToUpper(to)
	logger =logger.With("base", base)
	logger =logger.With("to", symbol)

	
	url := fmt.Sprintf("%s/convert?access_key=%s&from=%s&to=%s&amount=%d",
		config.BaseURL, config.APIKey, base, symbol, amount)
	logger = logger.With("url", url)
	logger.Info("making request to the rate excahnge api")

	httpClient := &http.Client{Timeout: 5* time.Second}

	resp, err := httpClient.Get(url)
	if err != nil {
		logger.Error("couldnt fetch response", "error", err)
		return nil, err

	}
	
	logger.Info("successful request")
	
	return resp.Body, nil
}