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
)


func Init() {
	prometheus.MustRegister(TotalApiRequests)
}