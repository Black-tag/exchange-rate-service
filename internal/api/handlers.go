package api

import (
	"encoding/json"
	"exchange-rate-service/internal/model"
	"exchange-rate-service/internal/repository"
	"fmt"
	"io"
	"log"

	"log/slog"
	"net/http"
)

func GetHealthOfAPI(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With(
		"handler", "gethealthofapi",
	)
	logger.Info("entered the handler")
	w.WriteHeader(http.StatusOK)
}



func GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With(
		"handler", "GetEchangeRate",
		"Method", r.Method,
	)
	logger.Info("entered handler")


	queryParams := r.URL.Query()
	from := queryParams.Get("from")
	to := queryParams.Get("to")
	date := queryParams.Get("date")

	logger = logger.With("from", from)
	logger = logger.With("to", to)
	logger = logger.With("date")
	rate, err := repository.FetchRates(from, to, date)
	if err != nil {
		logger.Error("cannot call external api")
		return

	}
	defer rate.Close()
	body, err := io.ReadAll(rate)
	if err != nil {
		logger.Error("cannot convert response body")
		log.Fatal(err)
	}
	logger = logger.With("function response", body)
	fmt.Println("rates", string(body))

	var data model.Rateresponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error("cannot decode json", "error", err)
	}
	
	w.Header().Set("Content-Type", "application/json")
	
}



