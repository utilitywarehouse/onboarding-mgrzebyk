package main

import (
	"net/http"
	"onboarding-mgrzebyk/handler"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)
func recordMetrics() {
    go func() {
            for {
                counter.Inc()
                time.Sleep(2 * time.Second)
            }
    }()
}

var (
    counter = promauto.NewCounter(prometheus.CounterOpts{
        Name: "requests_count",
        Help: "A counter of the number of total requests received",
      })
)

func main() {
    recordMetrics()
	mux := http.NewServeMux()
	th := handler.TimeHandler()
	mux.Handle("/time", th)
    mux.Handle("/__/metrics",promhttp.Handler())

	_ = http.ListenAndServe("0.0.0.0:8080", mux)

}
