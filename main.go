package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger) // To see logs whenever we do any kind of response and request.

	router.Get("/hello", basicHandler) // client asking server to get /hello endpoint

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
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
