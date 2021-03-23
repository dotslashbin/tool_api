package main

import (
	"log"
	"net/http"
	"tool_api/router"
)

func main() {
	router := router.Mux()
	log.Fatal(http.ListenAndServe(":5000", router))
}
