package main

import (
	"bytes"
	"google.golang.org/grpc/benchmark/latency"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

// bufConn is a net.Conn implemented by a bytes.Buffer (which is a ReadWriter).
type bufConn struct {
	*bytes.Buffer
}

func (bufConn) Close() error                       { panic("unimplemented") }
func (bufConn) LocalAddr() net.Addr                { panic("unimplemented") }
func (bufConn) RemoteAddr() net.Addr               { panic("unimplemented") }
func (bufConn) SetDeadline(t time.Time) error      { panic("unimplemneted") }
func (bufConn) SetReadDeadline(t time.Time) error  { panic("unimplemneted") }
func (bufConn) SetWriteDeadline(t time.Time) error { panic("unimplemneted") }

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
	var lty = 10

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

	bestApps := func(w http.ResponseWriter, req *http.Request) {
		var i = time.Duration(lty)
		slowConn, err := (&latency.Network{Kbps: 0, Latency: i * time.Millisecond, MTU: 5}).Conn(bufConn{&bytes.Buffer{}})
		if err != nil {
			log.Printf("Unexpected error creating connection: %v", err)
		}

		errWant := "measured network latency (10ms) higher than desired latency (5ms)"
		if _, err := (&latency.Network{Latency: 5 * time.Millisecond}).Conn(slowConn); err == nil || err.Error() != errWant {
			log.Printf("Conn() = _, %q; want _, %q", err, errWant)
		}

		/*n := &latency.Network{Kbps: 0, Latency: 1 * time.Millisecond, MTU: 5}
		n.Dialer (net.Dial) ("tcp", "localhost:8000")
		//clientConn, err := n.Conn(d)
		*/

		w.Header().Add("Content-Type", "application/json")
		io.WriteString(w, `{"bestApps":{"devOps": "DevOps"}}`)
	}

	introduceLatency := func(w http.ResponseWriter, req *http.Request) {
		ms := req.URL.Query()["ms"]
		if len(ms) == 0 {
			return
		}
		latencyInt, err := strconv.Atoi(ms[0])
		if err != nil {
			log.Printf("latency is not a number: %s", ms)
		}
		lty = latencyInt
		log.Printf("latency is set to %d ms", lty)

		// Infinitely fast CPU: time doesn't pass unless sleep is called.
		//tn := time.Unix(123, 0)
		//now := func() time.Time { return tn }
		//sleep := func(d time.Duration) { tn = tn.Add(d) }

		// Simulate a 10ms latency network, then attempt to simulate a 5ms latency
		// network and expect an error.
		/*
			slowConn, err := (&latency.Network{Kbps: 0, Latency: 100 * time.Millisecond, MTU: 5}).Conn(bufConn{&bytes.Buffer{}})
			if err != nil {
				log.Printf("Unexpected error creating connection: %v", err)
			}

			errWant := "measured network latency (10ms) higher than desired latency (5ms)"
			if _, err := (&latency.Network{Latency: 5 * time.Millisecond}).Conn(slowConn); err == nil || err.Error() != errWant {
				log.Printf("Conn() = _, %q; want _, %q", err, errWant)
			}
		*/
	}

	http.HandleFunc("/best-tools", bestTools)
	http.HandleFunc("/best-apps", bestApps)
	http.HandleFunc("/control/error", introduceHttpErrorCodes)
	http.HandleFunc("/control/latency", introduceLatency)
	log.Println("Listing for requests at http://localhost:8000/best-tools")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
