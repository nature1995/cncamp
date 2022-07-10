package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"hw5/internal/logger"
)

const nameSpace = "http_server"

var funcLatency = NewFuncLatencyHistogramMetric(nameSpace, "func_latency_seconds", "function latency analysis for seconds")

// Register registers the FuncLatencyHistogramMetric
func Register() {
	err := prometheus.Register(funcLatency)
	if err != nil {
		logger.Logger.Error(err)
	}
}

// FuncLatencyTimer is a timer for function latency
type FuncLatencyTimer struct {
	histogram *prometheus.HistogramVec
	start     time.Time
	last      time.Time
}

// NewTimer creates a new NewTimer
func NewTimer() *FuncLatencyTimer {
	return NewFuncLatencyTimer(funcLatency)
}

// NewFuncLatencyTimer creates a new FuncLatencyTimer
func NewFuncLatencyTimer(histogram *prometheus.HistogramVec) *FuncLatencyTimer {
	now := time.Now()
	return &FuncLatencyTimer{
		histogram: histogram,
		start:     now,
		last:      now,
	}
}

// ComputeTotal records the total time since the timer was created
func (t *FuncLatencyTimer) ComputeTotal() {
	(*t.histogram).WithLabelValues("total").Observe(time.Since(t.start).Seconds())
}

// NewFuncLatencyHistogramMetric creates a new FuncLatencyHistogramMetric
func NewFuncLatencyHistogramMetric(ns, name, help string) *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: ns,
			Name:      name,
			Help:      help,
			Buckets:   prometheus.ExponentialBuckets(0.001, 2, 15),
		},
		[]string{"step"},
	)
}
