package handlers

import (
	"log"
	"time"
	"sync/atomic"
	"idcyberark/counter"
	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(c *counter.SafeCounter) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)

	go func() {
		log.Printf("Ready probe is negative by default...")
		time.Sleep(5 * time.Second)
		isReady.Store(true)
		log.Printf("Ready probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/id/{jenkins}", idcyberark(c)).Methods("GET")
	r.HandleFunc("/health", health)
	r.HandleFunc("/ready", ready(isReady))
	return r
}