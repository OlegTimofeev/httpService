package httpService

import (
	"errors"
	client2 "github.com/go-openapi/runtime/client"
	util2 "github.com/itimofeev/go-util"
	"github.com/stretchr/testify/suite"
	"httpService/internal"
	"httpService/internal/dataBase"
	"httpService/service/client"
	"httpService/service/client/operations"
	models2 "httpService/service/models"
	"net/http"
	"testing"
)

func (hs *HandlersSuit) SetupTest() {
	config := dataBase.ConfigDB{
		User:      "admin",
		Password:  "password",
		Dbname:    "httpService",
		StoreType: "postgres",
	}
	hs.taskService = internal.NewTaskService(config)
	hs.requester = new(TestRequester)
	hs.taskService.SetRequester(hs.requester)
	response := new(models2.TaskResponse)
	hs.requester.SetResponse(response)
	httpClient := &http.Client{Transport: util2.NewTransport(hs.taskService.Server.GetHandler())}
	c := client2.NewWithClient(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes, httpClient)
	hs.taskClient = client.New(c, nil)
	//from initialization
	hs.task.Method = "POST"
	hs.task.Path = "https://www.google.com/"
	hs.task.Body = ""
}

type HandlersSuit struct {
	task        models2.FetchTask
	taskService *internal.TaskService
	taskClient  *client.FetchtaskHandlingService
	requester   *TestRequester
	suite.Suite
}

func (hs *HandlersSuit) TestAddFetchTask() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
}

func (hs *HandlersSuit) TestGetResponseError() {
	hs.requester.SetResponse(nil)
	hs.requester.SetError(errors.New("error : unable to get response"))
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().Error(err)
	hs.Require().Nil(getResponseOK)
}

func (hs *HandlersSuit) TestGetFetchTask() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	getResponseOK, err = hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	getTaskOK, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(int64(1)))
	hs.Require().NoError(err)
	hs.Require().NotNil(getTaskOK)
}

func (hs *HandlersSuit) TestGetAllFetchTasks() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	getResponseOK, err = hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	countOfTasks := 2
	getTasksOK, err := hs.taskClient.Operations.GetAllTasks(operations.NewGetAllTasksParams())
	hs.Require().NoError(err)
	hs.Require().EqualValues(countOfTasks, len(getTasksOK.Payload))
}

func (hs *HandlersSuit) TestDeleteFetchTasks() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Method: hs.task.Method,
		Path:   hs.task.Path,
		Body:   hs.task.Body,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	countOfTasks := 1
	getTasksOK, err := hs.taskClient.Operations.GetAllTasks(operations.NewGetAllTasksParams())
	hs.Require().NoError(err)
	hs.Require().EqualValues(countOfTasks, len(getTasksOK.Payload))
	_, err = hs.taskClient.Operations.DeleteFetchTask(operations.NewDeleteFetchTaskParams().WithTaskID(int64(1)))
	hs.Require().NoError(err)
	countOfTasks = 0
	getTasksOK, err = hs.taskClient.Operations.GetAllTasks(operations.NewGetAllTasksParams())
	hs.Require().NoError(err)
	hs.Require().EqualValues(countOfTasks, len(getTasksOK.Payload))
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersSuit))
}
