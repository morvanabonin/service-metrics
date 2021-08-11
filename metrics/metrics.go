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

// func metricsCountQuery() int {
// 	total := make(chan int)

// 	go counter.CountNumber(total)
// 	return printer(total)
// }

// func printer(in <-chan int) int {
// 	var total int
// 	for v := range in {
// 		total = v
// 	}
// 	return total
// }

var (
	countNumber = float64(counter.MetricsCountQuery())
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
