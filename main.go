package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"granblue.orbb.li/crawler/crawler"
)

type (
	Handler struct{}
)

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	crawler.Crawling(w, r)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Handler{},
	}
	log.Printf("Serving http")
	log.Fatal(httpServer.ListenAndServe())
}
