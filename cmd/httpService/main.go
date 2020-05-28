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
		PoolSize:      3,
		NatsUrl:       "nats://localhost:4222",
		StanClusterID: "test-cluster",
	}
	TaskService = internal.NewTaskService(config, request.NewRequester())
	defer TaskService.Server.Shutdown()
	defer TaskService.WorkerPool.Close()
	if err := TaskService.Server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
