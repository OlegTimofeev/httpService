package main

import (
	"httpService/api"
	"log"
)

func main() {
	server := api.InitSWHandler()
	defer server.Shutdown()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
