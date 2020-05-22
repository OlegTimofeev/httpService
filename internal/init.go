package internal

import (
	"github.com/go-openapi/loads"
	"httpService/internal/dataBase"
	"httpService/internal/models"
	"httpService/internal/request"
	"httpService/internal/taskHandler"
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
	TasksChan     chan *models.FetchTask
	ResponsesChan chan *models.TaskResponse
	Store         dataBase.DataStore
	Requester     request.Requester
	Server        *restapi.Server
	TaskHandler   taskHandler.TaskHandler
}

func NewTaskService(config dataBase.ConfigDB) *TaskService {
	taskService = &TaskService{
		TasksChan:     make(chan *models.FetchTask),
		ResponsesChan: make(chan *models.TaskResponse),
		Server:        initServer(),
		Requester:     request.NewRequester(),
		Store:         dataBase.NewDataStore(config),
	}
	taskService.TaskHandler = taskHandler.NewProdTaskHandler(taskService.TasksChan, taskService.ResponsesChan)
	return taskService
}

func (ts *TaskService) InitWorkers(config dataBase.ConfigDB) {
	taskService.TaskHandler = taskHandler.NewProdTaskHandler(taskService.TasksChan, taskService.ResponsesChan)
	taskService.TaskHandler.Init(Worker, Saver, config)
}

func (ts *TaskService) SetRequester(rm request.Requester) {
	ts.Requester = rm
}

func (ts *TaskService) SetTaskHandler(th taskHandler.TaskHandler) {
	ts.TaskHandler = th
}
