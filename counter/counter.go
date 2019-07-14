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
	return &SafeCounter{
		v: make(map[string]int),
		max: maxCounter,
	}
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