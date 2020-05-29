package internal

import (
	"github.com/go-openapi/loads"
	"httpService/internal/dataBase"
	"httpService/internal/models"
	"httpService/internal/request"
	"httpService/internal/taskWorker"
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
	api.CreateFetchTaskHandler = operations.CreateFetchTaskHandlerFunc(CreateFetchTask)
	api.DeleteFetchTaskHandler = operations.DeleteFetchTaskHandlerFunc(DeleteFT)
	api.GetAllTasksHandler = operations.GetAllTasksHandlerFunc(GetTasks)
	api.GetTaskHandler = operations.GetTaskHandlerFunc(GetTask)

	server.ConfigureFlags()
	server.ConfigureAPI()
	return server
}

type TaskService struct {
	Store      dataBase.DataStore
	WorkerPool models.WorkerPool
	Server     *restapi.Server
}

func NewTaskService(config dataBase.ConfigDB, requester request.Requester) *TaskService {
	taskService = &TaskService{
		WorkerPool: taskWorker.NewWorkerPool(requester),
		Server:     initServer(),
		Store:      dataBase.NewDataStore(config),
	}
	return taskService
}
