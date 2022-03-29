package main

import (
	"fmt"
	"go-todo/router"
	"log"
	"net/http"
	"os"
)

func main() {

	os.Setenv("PORT", "8080")

	r := router.Router()

	fmt.Printf("Starting server on the port %s...\n", os.Getenv("PORT"))

	if val, exist := os.LookupEnv("PORT"); exist {
		log.Fatal(val)
	} else {
		log.Fatal(":3001")
	}

	log.Fatal(http.ListenAndServe("PORT", r))

}
