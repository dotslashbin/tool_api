package main

import (
	"fmt"
	"log"
	"net/http"
	"tool_api/router"
)

func main() {
	router := router.Mux()
	fmt.Println("webserver running ... version 2")
	log.Fatal(http.ListenAndServe(":5000", router))
}
