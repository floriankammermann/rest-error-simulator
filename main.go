package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	bestTools := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, `{"bestTools":{"cidcd": "Jenkins"}}`)
		w.Header().Add("Content-Type", "application/json")
	}

	http.HandleFunc("/best-tools", bestTools)
	log.Println("Listing for requests at http://localhost:8000/best-tools")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
