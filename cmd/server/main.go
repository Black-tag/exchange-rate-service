package main

import (
	"exchange-rate-service/internal/api"
	"exchange-rate-service/internal/metrics"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {
	metrics.Init()


	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", api.GetHealthOfAPI)
	mux.HandleFunc("GET /api/latest", api.GetLatestExchangeRate)
	mux.HandleFunc("GET /api/exchange", api.GetConvertedExchangeRate)
	mux.HandleFunc("GET /api/convert", api.GetHistoricalExchangeRate)
	mux.Handle("/metrics", promhttp.Handler())
	server :=&http.Server{
		Addr: ":8080",
		Handler: mux,
	} 
	logger.Info("satrted server")
	fmt.Printf("starting server on port:8080")
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("cannot start server", "err", err)
		fmt.Printf("Server failed: %v\n", err)

	}

}