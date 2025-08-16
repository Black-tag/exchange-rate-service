package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)




var (
	TotalApiRequests = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "Excahnge-rate-api",
		Name: "total_api_requests",
		Help: "Sum of total requests",
	},[]string{"endpoint", "method", "status"})

	TotalActicverequest = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "total_active_requests",
		Help: "current_request",
	})
)


func Init() {
	prometheus.MustRegister(TotalApiRequests, TotalActicverequest)
}