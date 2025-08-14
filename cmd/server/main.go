package main

import (
	"exchange-rate-service/internal/api"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)


func main() {


	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", api.GetHealthOfAPI)
	mux.HandleFunc("GET /api/convert", api.GetExchangeRate)

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