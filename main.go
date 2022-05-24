package main

import (
	"log"
	"net/http"
	"os"

	"github.com/megpha/website/web"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Handler: web.CreateRouter(),
		Addr:    ":" + port,
	}

	log.Fatal(server.ListenAndServe())
}
