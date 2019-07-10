package main

import (
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
		"Starting the service on port %s...\ncommit: %s, build time: %s, release: %s",
		Port, version.Commit, version.BuildTime, version.Release,
	)	
	
	MaxCounter = counter.MaxCounter
	maxCounterEnv := os.Getenv("MAXCOUNTER")
	maxCounterInt, err := strconv.Atoi(maxCounterEnv)
	if err == nil {
		MaxCounter = maxCounterInt
	}

	log.Printf("The service will deliver id from 1 to %d", MaxCounter)
}

func main() {

	counterSafe := counter.New(MaxCounter)
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
	
	killSignal := <-interrupt
	switch killSignal {
		case os.Kill:
			log.Print("Got SIGKILL...")
		case os.Interrupt:
			log.Print("Got SIGINT...")
		case syscall.SIGTERM:
			log.Print("Got SIGTERM...")
	}

	log.Print("The service is shutting down...")
	srv.Shutdown(context.Background())
	log.Print("Done")
}