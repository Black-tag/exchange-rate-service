package api

import (
	"encoding/json"
	"exchange-rate-service/internal/model"
	"exchange-rate-service/internal/repository"
	"exchange-rate-service/internal/services"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

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



func GetLatestExchangeRate(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With(
		"handler", "GetEchangeRate",
		"Method", r.Method,
	)
	logger.Info("entered handler")

	queryParams := r.URL.Query()
	from := queryParams.Get("from")
	to := queryParams.Get("to")
	
	rate, err := repository.FetchRates()
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

	var data model.RateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("error  parsing json")
		logger.Error("cannot decode json", "error", err)
	}
	From := strings.ToUpper(from)
	To := strings.ToUpper(to)
	
	fmt.Print(data.Quotes)
	amount := services.ConvertCurrency(From, To, data.Quotes)
	fmt.Println("amount", amount)
	
	

	
	w.Header().Set("Content-Type", "application/json")
	
}


func GetConvertedExchangeRate (w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With(
		"handler", "GetHistoricalExchangeRate",
		"method", r.Method,
		"url", r.URL,
	)
	logger.Info("entered GethistoricalExchange handler")
	
	queryParams := r.URL.Query()
	from := queryParams.Get("from")
	to := queryParams.Get("to")
	amountStr := queryParams.Get("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		logger.Error("failed to convert amount from str to int")
		return
	}


	rate , err := repository.ConvertAmount(from, to, amount)
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

	var data model.RateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error("cannot decode json", "error", err)
	}
	
	w.Header().Set("Content-Type", "application/json")

}


func GetHistoricalExchangeRate(w http.ResponseWriter, r *http.Request) {
	logger := slog.Default().With(
		"handler", "GetHistoricalExchangeRate",
		"method", r.Method,
		"url", r.URL,
	)
	logger.Info("entered GethistoricalExchange handler")


	queryParams := r.URL.Query()
	from := queryParams.Get("from")
	to := queryParams.Get("to")
	date := queryParams.Get("date")
	if date == "" {
		fmt.Println("missing date and calling latest")
		GetLatestExchangeRate(w, r)
		return
		
	}

	logger = logger.With("from", from)
	logger = logger.With("to", to)
	logger = logger.With("date", date)

	currentDate := time.Now().Truncate(24 * time.Hour)

	givenDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		log.Fatalf("error parsing given date: %v", err)

	}

	ninetydaysBeforeNow := currentDate.AddDate(0,0,-90)
	fmt.Println(ninetydaysBeforeNow)

	if currentDate.Sub(givenDate) > 90 * 24 * time.Hour {
		http.Error(w, "given date is more than 90 days in past", http.StatusBadRequest)
		return

	}
	
	

	rate , err := repository.FetchHistoricalData(date)
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

	var data model.RateResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		logger.Error("cannot decode json", "error", err)
	}
	From := strings.ToUpper(from)
	To := strings.ToUpper(to)
	
	fmt.Print(data.Quotes)
	amount := services.ConvertCurrency(From, To, data.Quotes)
	fmt.Println("amount", amount)
	
	w.Header().Set("Content-Type", "application/json")

}



