package main

import (
	"httpService/api"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(api.InitializeServer()))
}
