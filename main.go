package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/kelseyhightower/envconfig"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	ResponseCodeSuccess                   int
	ResponseCodeFailure                   int
	ResponseCodeSuccessFailureRatio       int
	ResponseCodeSuccessFailureRatioModulo int
}

func (s *Specification) init() {
	if s.ResponseCodeSuccess == 0 {
		s.ResponseCodeSuccess = 200
	}
	if s.ResponseCodeFailure == 0 {
		s.ResponseCodeFailure = 500
	}
	if s.ResponseCodeSuccessFailureRatio == 0 {
		s.ResponseCodeSuccessFailureRatioModulo = 1
	}
	if s.ResponseCodeSuccessFailureRatio == 50 {
		s.ResponseCodeSuccessFailureRatioModulo = 2
	}
}

func getResponseCode(requestCounter, ratio, successCode, errorCode int) int {
	rest := requestCounter % ratio
	if rest == 0 {
		return successCode
	} else {
		return errorCode
	}
}

func setRestRatio(errorratioInt int) int {
	restratio := 100 / errorratioInt
	return restratio
}

func main() {
	var s Specification
	err := envconfig.Process("res", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	s.init()

	var requestCounter = 1

	bestTools := func(w http.ResponseWriter, req *http.Request) {
		rest := requestCounter % s.ResponseCodeSuccessFailureRatioModulo
		if rest != 0 {
			w.WriteHeader(s.ResponseCodeSuccess)
			ResponseCodeInternalServerError.Inc()
		} else {
			w.WriteHeader(s.ResponseCodeFailure)
			ResponseCodeStatusOK.Inc()
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"bestTools":{"cidcd": "Jenkins"}}`)
		requestCounter++
		log.Printf("requestCounter: %d", requestCounter)
		log.Printf("ratio: %d", s.ResponseCodeSuccessFailureRatio)
	}

	introduceHttpErrorCodes := func(w http.ResponseWriter, req *http.Request) {
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
			s.ResponseCodeSuccessFailureRatio = setRestRatio(errorratioInt
			logResponse := "set erroratio to " + strconv.Itoa(ratio) + " (" + strconv.Itoa(errorratioInt) + "%)"
			log.Println(logResponse)
		}
	}

	http.HandleFunc("/best-tools", bestTools)
	http.HandleFunc("/control/error", introduceHttpErrorCodes)
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Listening for requests at http://localhost:8080/best-tools")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
