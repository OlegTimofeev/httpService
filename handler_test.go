package httpService

import (
	"errors"
	client2 "github.com/go-openapi/runtime/client"
	util2 "github.com/itimofeev/go-util"
	"github.com/stretchr/testify/suite"
	"httpService/internal"
	"httpService/internal/dataBase"
	"httpService/internal/models"
	"httpService/service/client"
	"httpService/service/client/operations"
	"net/http"
	"testing"
)

func (hs *HandlersSuit) SetupTest() {
	config := dataBase.ConfigDB{
		User:          "admin",
		Password:      "password",
		Dbname:        "httpService",
		StoreType:     "postgres",
		NatsUrl:       "nats://localhost:4222",
		StanClusterID: "test-cluster",
	}
	hs.response = new(models.TaskResponse)
	hs.requester = &TestRequester{
		response: hs.response,
		err:      hs.err,
	}
	hs.taskService = internal.NewTaskService(config, hs.requester)
	response := new(models.TaskResponse)
	hs.requester.SetResponse(response)
	httpClient := &http.Client{Transport: util2.NewTransport(hs.taskService.Server.GetHandler())}
	c := client2.NewWithClient(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes, httpClient)
	hs.taskClient = client.New(c, nil)
}

type HandlersSuit struct {
	response    *models.TaskResponse
	err         error
	taskService *internal.TaskService
	taskClient  *client.FetchtaskHandlingService
	requester   *TestRequester
	suite.Suite
}

func (hs *HandlersSuit) TestAddFetchTask() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	ft, err := hs.taskService.Store.GetFetchTask(int(getResponseOK.Payload.ID))
	hs.Require().NotNil(ft)
	hs.Require().NoError(err)
}

func (hs *HandlersSuit) TestGetTaskWithError() {
	hs.err = errors.New("error")
	task, err := hs.taskService.Store.AddFetchTask(&models.FetchTask{})
	hs.Require().NotNil(task)
	hs.Require().NoError(err)
	err = hs.taskService.Store.SetResponse(task.ID, hs.response, hs.err)
	hs.Require().NoError(err)
	getTaskOk, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(int64(task.ID)))
	hs.Require().NoError(err)
	hs.Require().Nil(getTaskOk.Payload.Response)
	hs.Require().EqualValues(models.StatusError, getTaskOk.Payload.Request.Progress)
}

func (hs *HandlersSuit) TestGetFetchTask() {
	task, err := hs.taskService.Store.AddFetchTask(&models.FetchTask{})
	hs.Require().NotNil(task)
	hs.Require().NoError(err)
	err = hs.taskService.Store.SetResponse(task.ID, hs.response, hs.err)
	getTaskOK, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(int64(task.ID)))
	hs.Require().NoError(err)
	hs.Require().NotNil(getTaskOK)
	hs.EqualValues(int64(task.ID), getTaskOK.Payload.Request.ID)
	hs.Require().EqualValues(models.StatusCompleted, getTaskOK.Payload.Request.Progress)
}

func (hs *HandlersSuit) TestGetAllFetchTasks() {
	getTasksOK, err := hs.taskClient.Operations.GetAllTasks(operations.NewGetAllTasksParams())
	hs.Require().NoError(err)
	countOfTasks := len(getTasksOK.Payload)
	countOfAddedTasks := 2
	for i := 0; i < countOfAddedTasks; i++ {
		_, err = hs.taskService.Store.AddFetchTask(&models.FetchTask{})
		hs.Require().NoError(err)
	}
	getTasksOK, err = hs.taskClient.Operations.GetAllTasks(operations.NewGetAllTasksParams())
	hs.Require().NoError(err)
	hs.Require().EqualValues(countOfTasks+countOfAddedTasks, len(getTasksOK.Payload))
}

func (hs *HandlersSuit) TestDeleteFetchTasks() {
	task, err := hs.taskService.Store.AddFetchTask(&models.FetchTask{})
	hs.Require().NotNil(task)
	hs.Require().NoError(err)
	err = hs.taskService.Store.SetResponse(task.ID, hs.response, hs.err)
	_, err = hs.taskClient.Operations.DeleteFetchTask(operations.NewDeleteFetchTaskParams().WithTaskID(int64(task.ID)))
	hs.Require().NoError(err)
	getTaskOk, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(int64(task.ID)))
	hs.Require().Nil(getTaskOk)
	hs.Require().Error(err)
	res, err := hs.taskService.Store.GetTaskResponseByFtID(task.ID)
	hs.Require().Nil(res)
	hs.Require().Error(err)
}

func TestHandlers(t *testing.T) {
	suite.Run(t, new(HandlersSuit))
}
