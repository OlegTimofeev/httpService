package internal

import (
	"github.com/go-openapi/loads"
	"httpService/internal/dataBase"
	"httpService/internal/request"
	"httpService/service/restapi"
	"httpService/service/restapi/operations"
	"log"
)

var taskService *TaskService

func initServer() *restapi.Server {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewFetchtaskHandlingServiceAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.Port = 8080
	api.CreateFetchTaskHandler = operations.CreateFetchTaskHandlerFunc(GetResponse)
	api.DeleteFetchTaskHandler = operations.DeleteFetchTaskHandlerFunc(DeleteFT)
	api.GetAllTasksHandler = operations.GetAllTasksHandlerFunc(GetTasks)
	api.GetTaskHandler = operations.GetTaskHandlerFunc(GetTask)

	server.ConfigureFlags()
	server.ConfigureAPI()
	return server
}

type TaskService struct {
	Store     dataBase.DataStore
	Requester request.RequesterModel
	Server    *restapi.Server
}

func NewTaskService(config dataBase.ConfigDB) *TaskService {
	taskService = &TaskService{
		Server:    initServer(),
		Requester: request.NewRequester(),
		Store:     dataBase.NewDataStore(config),
	}
	return taskService
}

func (ts *TaskService) SetRequester(rm request.RequesterModel) {
	ts.Requester = rm
}