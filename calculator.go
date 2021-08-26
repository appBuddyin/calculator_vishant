package main

import (
	"log"
	"net/http"

	"calculator/math"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/calculator/{operation}/{var1:[0-9]+}/{var2:[0-9]+}", math.Calculate)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}
	log.Fatal(srv.ListenAndServe())
}
