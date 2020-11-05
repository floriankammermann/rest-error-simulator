package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Specification struct {
	Port                  int
	RequestFrequencyInSec int
	Endpoint              string
}

func execRequest(s Specification) {
	resp, err := http.Get(s.Endpoint)
	if err != nil {
		log.Printf("got error, while calling %s: %s", s.Endpoint, err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("got error, while reading response of %s: %s", s.Endpoint, err)
		return
	}

	log.Printf("endpoint: %s, responseCode %d, responseBody: %s", s.Endpoint, resp.StatusCode, string(body))
}

func callServer(s Specification, quit chan struct{}) {
	ticker := time.NewTicker(time.Duration(s.RequestFrequencyInSec) * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Printf("Call endpoint: %s, with frequency: %d", s.Endpoint, s.RequestFrequencyInSec)
			execRequest(s)
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
	if s.Port == 0 {
		s.Port = 8080
	}
	log.Printf("Listening for requests at http://localhost:%d/control", s.Port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(s.Port), nil))
}
