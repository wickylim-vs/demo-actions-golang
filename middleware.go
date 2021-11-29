package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	opsRequest = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "vs_app",
		Name:      "http_request_ops_total",
		Help:      "The total number of requests",
	})
)

//statusRecorder to record the status code from the ResponseWriter
type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(statusCode int) {
	rec.statusCode = statusCode
	rec.ResponseWriter.WriteHeader(statusCode)
}

func measureResponseDuration(next http.Handler) http.Handler {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	responseTimeHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "vs_app",
		Name:      "http_request_duration_seconds",
		Help:      "Histogram of response time for handler in seconds",
		Buckets:   buckets,
	}, []string{"method", "status_code"})

	prometheus.MustRegister(responseTimeHistogram)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := statusRecorder{w, 200}

		next.ServeHTTP(&rec, r)

		duration := time.Since(start)
		statusCode := strconv.Itoa(rec.statusCode)
		responseTimeHistogram.WithLabelValues(r.Method, statusCode).Observe(duration.Seconds())
	})
}
