package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	ResponseCodeInternalServerError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_internal_server_error",
		Help: "amount of internal server errors",
	})
	ResponseCodeStatusOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_status_ok",
		Help: "amount of status ok",
	})
)

type Specification struct {
	RequestFrequencyInSec int
	Endpoint              string
}

func getResponseCode(requestCounter, ratio, successCode, errorCode int) int {
	rest := requestCounter % ratio
	if rest == 0 {
		return successCode
	} else {
		return errorCode
	}
}

func callServer(s Specification, quit chan struct{}) {
	ticker := time.NewTicker(time.Duration(s.RequestFrequencyInSec) * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Printf("Call endpoint: %s, with frequency: %d", s.Endpoint, s.RequestFrequencyInSec)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func main() {
	var s Specification
	err := envconfig.Process("res", &s)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("got endpoint: %s, frequency: %d", s.Endpoint, s.RequestFrequencyInSec)
	var quit chan struct{}

	if s.RequestFrequencyInSec > 0 && s.Endpoint != "" {
		quit = make(chan struct{})
		go callServer(s, quit)
	}

	setEndpointAndFrequency := func(w http.ResponseWriter, req *http.Request) {
		frequency := req.URL.Query()["frequency"]
		endpoint := req.URL.Query()["endpoint"]

		if len(frequency) != 0 {
			frequencyInt, err := strconv.Atoi(frequency[0])
			if err != nil {
				log.Printf("frequency is not a number: %s", frequency)
			}
			s.RequestFrequencyInSec = frequencyInt
			log.Printf("set Frequency to %d", s.RequestFrequencyInSec)

			// stop to old frequency
			close(quit)

			// start the new frequency
			quit = make(chan struct{})
			go callServer(s, quit)
		}
		if len(endpoint) != 0 {
			s.Endpoint = endpoint[0]
			log.Printf("set Endpoint to %s", s.Endpoint)
		}
	}

	http.HandleFunc("/control", setEndpointAndFrequency)
	log.Println("Listening for requests at http://localhost:8080/control")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
