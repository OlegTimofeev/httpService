package api

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"httpService/dataBase"
	"httpService/service/restapi"
	"httpService/service/restapi/operations"
	"log"
)

var db dataBase.DataStore

func InitSWHandler() *restapi.Server {
	db = new(dataBase.PostgresDB)
	db.InitDB()
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewFetchtaskHandlingServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8080
	api.CreateFetchTaskHandler = operations.CreateFetchTaskHandlerFunc(func(params operations.CreateFetchTaskParams) middleware.Responder {
		return GetResponse(params)
	})
	api.DeleteFetchTaskHandler = operations.DeleteFetchTaskHandlerFunc(func(params operations.DeleteFetchTaskParams) middleware.Responder {
		return DeleteFT(params)
	})
	api.GetAllTasksHandler = operations.GetAllTasksHandlerFunc(func(params operations.GetAllTasksParams) middleware.Responder {
		return GetTasks()
	})
	api.GetTaskHandler = operations.GetTaskHandlerFunc(func(params operations.GetTaskParams) middleware.Responder {
		return GetTask(params)
	})

	server.ConfigureFlags()
	server.ConfigureAPI()
	return server
}
