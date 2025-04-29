package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("failed to listen to server: %v", err)
		os.Exit(1)
	}
}

func basicHandler(rw http.ResponseWriter, req *http.Request) {
	// stand response writer and request set up.
	// req is the inbound request received from the client
	rw.Write([]byte("Hi Hi Hi sheeps"))
}
