package main

import (
	"httpService/internal"
	"httpService/internal/dataBase"
	"httpService/internal/request"
	"log"
)

var TaskService *internal.TaskService

func main() {
	config := dataBase.ConfigDB{
		User:          "admin",
		Password:      "password",
		Dbname:        "httpService",
		StoreType:     "postgres",
		NatsUrl:       "nats://localhost:4222",
		StanClusterID: "test-cluster",
	}
	TaskService = internal.NewTaskService(config, request.NewRequester())
	defer TaskService.Server.Shutdown()
	if err := TaskService.Server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
