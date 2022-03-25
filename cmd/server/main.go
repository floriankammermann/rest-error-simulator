package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ResponseCodeInternalServerError = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_internal_server_error",
		Help: "amount of response internal server errors",
	})
	ResponseCodeStatusOK = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_ok",
		Help: "amount of response status ok",
	})
	ResponseCodeAll = promauto.NewCounter(prometheus.CounterOpts{
		Name: "response_all",
		Help: "all response",
	})
)

type Specification struct {
	ResponseCodeSuccess             int
	ResponseCodeFailure             int
	ResponseCodeSuccessFailureRatio int
	LatencyInMs                     int
}

var failureRatioModulo int
var latencyinms int

func (s *Specification) init() {
	if s.ResponseCodeSuccess == 0 {
		s.ResponseCodeSuccess = 200
	}
	if s.ResponseCodeFailure == 0 {
		s.ResponseCodeFailure = 500
	}
	if s.ResponseCodeSuccessFailureRatio == 0 {
		s.ResponseCodeSuccessFailureRatio = 1
	}
}

func calculateFailureRationModulo(errorratioInt int) int {
	restratio := 100 / errorratioInt
	return restratio
}

func getResponseCode(requestCounter, ratioModulo, successCode, errorCode int) int {
	// if we do not deliver every request as success on ratioModulo 100, we do not get 100% successCode
	if ratioModulo == 100 {
		return successCode
	}
	rest := requestCounter % ratioModulo
	if rest != 0 {
		return successCode
	} else {
		return errorCode
	}
}

func main() {
	s := &Specification{}
	err := envconfig.Process("res", s)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.init()
	failureRatioModulo = calculateFailureRationModulo(s.ResponseCodeSuccessFailureRatio)
	log.Printf("start with responseCodeSuccess: %d, responseCodeFailure: %d, responseCodeFailureRatio: %d, failureRatioModulo: %d", s.ResponseCodeSuccess, s.ResponseCodeFailure, s.ResponseCodeSuccessFailureRatio, failureRatioModulo)

	var requestCounter = 0

	bestTools := func(w http.ResponseWriter, req *http.Request) {
		// we have to add 1 to the requestCounter to prohibit modulo operation with 0
		responseCode := getResponseCode(requestCounter+1, failureRatioModulo, s.ResponseCodeSuccess, s.ResponseCodeFailure)
		w.WriteHeader(responseCode)
		if responseCode == s.ResponseCodeSuccess {
			log.Printf("return success responseCode %d", s.ResponseCodeSuccess)
			ResponseCodeStatusOK.Inc()
		} else {
			log.Printf("return failure responseCode %d", s.ResponseCodeFailure)
			ResponseCodeInternalServerError.Inc()
		}
		time.Sleep(time.Duration(s.LatencyInMs) * time.Millisecond)
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"bestTools":{"cidcd": "Jenkins"}}`)
		requestCounter++
		ResponseCodeAll.Inc()
		log.Printf("requestCounter: %d", requestCounter)
		log.Printf("ratioModulo: %d", failureRatioModulo)
	}

	introduceHttpErrorCodes := func(w http.ResponseWriter, req *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Methods", "OPTION")
		w.Header().Add("Content-Type", "application/json")

		if req.Method != "POST" {
			w.WriteHeader(405)
			io.WriteString(w, "only POST allowed")
			return
		}

		errorcode := req.URL.Query()["errorcode"]
		errorratio := req.URL.Query()["errorratio"]

		if len(errorcode) != 0 {
			errorcodeInt, err := strconv.Atoi(errorcode[0])
			if err != nil {
				log.Printf("errorcode is not a number: %s", errorcode)
			}
			s.ResponseCodeFailure = errorcodeInt
			log.Printf("set ResponseCode to %d", s.ResponseCodeFailure)
		}
		if len(errorratio) != 0 {
			errorratioInt, err := strconv.Atoi(errorratio[0])
			if err != nil {
				log.Printf("errorratio is not a number: %s", errorratio)
			}
			failureRatioModulo = calculateFailureRationModulo(errorratioInt)
			log.Printf("set failureRatioModulo to %d", failureRatioModulo)
		}
	}

	introduceLatency := func(w http.ResponseWriter, req *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Methods", "OPTION")
		w.Header().Add("Content-Type", "application/json")

		if req.Method != "POST" {
			w.WriteHeader(405)
			io.WriteString(w, "only POST allowed")
			return
		}

		latencyinmsStr := req.URL.Query()["latencyinms"]

		if len(latencyinmsStr) != 0 {
			latencyinms, err := strconv.Atoi(latencyinmsStr[0])
			if err != nil {
				log.Printf("latencyinms is not a number: %s", latencyinmsStr)
			}
			s.LatencyInMs = latencyinms
			log.Printf("set latencyinms to %d", latencyinms)
		}
	}

	controlParams := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		responseBody := fmt.Sprintf(
			`{"responseCodeSuccess":%d, "responseCodeFailure":%d, "responseCodeSuccessFailureRatio":%d, "ratioModulo":%d, "requestCounter":%d}`,
			s.ResponseCodeSuccess,
			s.ResponseCodeFailure,
			s.ResponseCodeSuccessFailureRatio,
			failureRatioModulo,
			requestCounter,
		)
		io.WriteString(w, responseBody)
	}

	http.HandleFunc("/best-tools", bestTools)
	http.HandleFunc("/control/error", introduceHttpErrorCodes)
	http.HandleFunc("/control/latency", introduceLatency)
	http.HandleFunc("/control", controlParams)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening for requests at http://localhost:8080/best-tools")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
