package main

import (
	"httpService/internal"
	"httpService/internal/dataBase"
	"log"
)

var TaskService *internal.TaskService

func main() {
	config := dataBase.ConfigDB{
		User:      "admin",
		Password:  "password",
		Dbname:    "httpService",
		StoreType: "postgres",
		PoolSize:  3,
	}
	TaskService = internal.NewTaskService(config)
	defer TaskService.Server.Shutdown()
	if err := TaskService.Server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
