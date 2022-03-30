package main

import (
	"fmt"
	"go-todo/router"
	"log"
	"net/http"
	"os"
)

func main() {

	r := router.Router()

	fmt.Printf("Starting server on the port %s...\n", os.Getenv("PORT"))

	if val, exist := os.LookupEnv("PORT"); exist {
		log.Fatal(http.ListenAndServe(":"+val, r))
	} else {
		log.Fatal(http.ListenAndServe(":3001", r))
	}

}
