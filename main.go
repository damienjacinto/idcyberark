package main

import (
	"time"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"context"
	"strconv"
	"idcyberark/handlers"
	"idcyberark/counter"
	"idcyberark/version"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Port string = ""
	MaxCounter int
)

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		log.Fatal("Port is not set.")
	}

	log.Printf(
		"Starting the service on port %s...\ncommit: %s, release: %s",
		Port, version.Commit, version.Release,
	)	
	
	MaxCounter = counter.MaxCounter
	maxCounterEnv := os.Getenv("MAXCOUNTER")
	maxCounterInt, err := strconv.Atoi(maxCounterEnv)
	if (err == nil && maxCounterInt > 0) {
		MaxCounter = maxCounterInt
	}

	log.Printf("The service will deliver id from 1 to %d", MaxCounter)
}

func main() {

	counterSafe := counter.New(MaxCounter)

	counterGaugePrometheus(counterSafe)

	router := handlers.Router(counterSafe)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    ":" + Port,
		Handler: router,
	}
	
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
		
	gracefullShutdown(srv, interrupt)

	log.Println("Server shutdown")
}

func counterGaugePrometheus(counterSafe *counter.SafeCounter) {

	if err := prometheus.Register(counterSafe.Metrics.CounterGauge); err != nil {
		log.Println("CounterGauge not registered:", err)
	} else {
		log.Println("CounterGauge registered.")
	}
}

func gracefullShutdown(srv *http.Server, interrupt <-chan os.Signal) {
	killSignal := <-interrupt
	log.Print("The service is shutting down...")
	switch killSignal {
		case os.Kill:
			log.Print("Got SIGKILL...")
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Failed to gracefully shutdown:", err)
	}
}