package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// r := router.Router()

	fmt.Printf("Starting server on the port %s...\n", "PORT")

	if val, exist := os.LookupEnv("PORT"); exist {
		log.Fatal(val)
	} else {
		log.Fatal(":3001")
	}

	// log.Fatal(http.ListenAndServe("PORT", r))
}
