package counter

import (
	"sync"
	"github.com/prometheus/client_golang/prometheus"
)

const MaxCounter = 100

type SafeCounter struct {
	v 		map[string]int
	mux 	sync.Mutex
	max 	int
	Metrics	Metrics
}

type Metrics struct {
    CounterGauge *prometheus.GaugeVec
}

func New(maxCounter int) *SafeCounter {
	c := SafeCounter{
		v: make(map[string]int),
		max: maxCounter,
	}

	c.Metrics.CounterGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "idcyberark",
			Name:      "counter_info",
			Help:      "Value of the last id emitted for the counter",
		},
		[]string{"counter_name"},
	)

	return &c
}

func (c *SafeCounter) Inc(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	if (c.v[key]>=c.max) {
		c.v[key]=1
	} else {
		c.v[key]++
	}
	return c.v[key]
}