package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	LogsReceived = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "devlogger_logs_received_total",
			Help: "Total number of logs received",
		})

	LogsQueried = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "devlogger_logs_queried_total",
			Help: "Total number of logs queried",
		})
)

func Init() {
	prometheus.MustRegister(LogsReceived)
	prometheus.MustRegister(LogsQueried)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9090", nil)
	}()
}
