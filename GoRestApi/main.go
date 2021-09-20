package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Go server")
	}).Methods(http.MethodGet)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", r))
}
