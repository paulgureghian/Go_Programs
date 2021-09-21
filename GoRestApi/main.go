package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/paulgureghian/Go_Programs/GoRestApi/routes"
)

func main() {

	router := routes.MovieRoutes()

	http.Handle("/api/", router)

	fmt.Println("Go server")
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8081", router))

}
