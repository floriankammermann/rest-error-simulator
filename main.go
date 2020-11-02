package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

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

	var responseCode = http.StatusOK
	var responseCodeSuccess = http.StatusOK
	var ratio = 1
	var requestCounter = 1

	bestTools := func(w http.ResponseWriter, req *http.Request) {
		rest := requestCounter % ratio
		if rest != 0 {
			w.WriteHeader(responseCodeSuccess)
		} else {
			w.WriteHeader(responseCode)
		}
		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"bestTools":{"cidcd": "Jenkins"}}`)
		requestCounter++
		log.Printf("requestCounter: %d", requestCounter)
		log.Printf("ratio: %d", ratio)
	}

	introduceHttpErrorCodes := func(w http.ResponseWriter, req *http.Request) {
		errorcode := req.URL.Query()["errorcode"]
		if len(errorcode) == 0 {
			return
		}
		errorcodeInt, err := strconv.Atoi(errorcode[0])
		if err != nil {
			log.Printf("errorcode is not a number: %s", errorcode)
		}
		log.Printf("set errorcode to %s", errorcode)
		responseCode = errorcodeInt

		errorratio := req.URL.Query()["errorratio"]
		if len(errorratio) == 0 {
			return
		}
		errorratioInt, err := strconv.Atoi(errorratio[0])
		if err != nil {
			log.Printf("errorratio is not a number: %s", errorratio)
		}
		/*if errorratioInt == 50 {
			log.Println("set errorratio to 2 (50%)")
			ratio = 2
		} else if errorratioInt == 20 {
			log.Println("set errorratio to 5 (20%)")
			ratio = 5
		} else if errorratioInt == 33 {
			log.Println("set errorratio to 3 (30%)")
			ratio = 3
		}*/
		ratio = setRestRatio(errorratioInt)
		logResponse := "set erroratio to " + strconv.Itoa(ratio) + " (" + strconv.Itoa(errorratioInt) + "%)"
		log.Println(logResponse)

	}

	http.HandleFunc("/best-tools", bestTools)
	http.HandleFunc("/control/error", introduceHttpErrorCodes)
	log.Println("Listing for requests at http://localhost:8000/best-tools")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
