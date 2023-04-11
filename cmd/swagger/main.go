package main

import (
	"log"
	"net/http"
	"os"

	"github.com/flowchartsman/swaggerui"
)

func main() {
	spec, err := os.ReadFile("swaggerui/profile/v1/profile.swagger.json")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	http.Handle("/swaggerui/", http.StripPrefix("/swaggerui", swaggerui.Handler(spec)))
	log.Println("serving on :8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
