package metrics

import (
	"github.com/morvanabonin/service-metrics/counter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
)

type (
	MetricsService interface {
		counter.Counter
	}

	Metric struct {
		metric counter.Counter
	}
)

var (
	countNumber = float64(counter.CountNumber())
	//countString = strconv.Itoa(countNumber)
)

var (
	opsCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "service_metrics",
			Name:      "counter",
			Help:      "Number Counter",
		},
		[]string{"count"},
	)
)

func init() {
	prometheus.MustRegister(opsCount)
	prometheus.MustRegister(collectors.NewBuildInfoCollector())
}

func Metrics() {
	opsCount.WithLabelValues("number").Add(countNumber)
}
