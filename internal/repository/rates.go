package repository

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"
)





func FetchRates() (io.ReadCloser, error) {

	// API_KEY := "0a40d8593652f00fe31321015eb90267"
	logger := slog.Default().With(
		"handler", "FetchRates",
	)
	
	
	 
	
	url := "https://api.exchangerate.host/live?access_key=0a40d8593652f00fe31321015eb90267"
	
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
	url := fmt.Sprintf("https://api.exchangerate.host/historical?access_key=0a40d8593652f00fe31321015eb90267&date=%s", date)
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

func ConvertAmount(from string, to string, amount int) (io.ReadCloser, error) {
	logger := slog.Default().With(
		"function", "FetchHistoricalData",
	)
	
	base := strings.ToUpper(from)
	symbol := strings.ToUpper(to)
	logger =logger.With("base", base)
	logger =logger.With("to", symbol)

	
	url := fmt.Sprintf("https://api.exchangerate.host/convert?access_key=0a40d8593652f00fe31321015eb90267&from=%s&to=%s&amount=%d", from, to, amount)
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