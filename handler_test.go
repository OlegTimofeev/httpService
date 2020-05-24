package httpService

import (
	client2 "github.com/go-openapi/runtime/client"
	util2 "github.com/itimofeev/go-util"
	"github.com/stretchr/testify/suite"
	"httpService/internal"
	"httpService/internal/dataBase"
	"httpService/internal/models"
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
		PoolSize:  2,
	}
	hs.taskService = internal.NewTaskService(config)
	hs.requester = new(TestRequester)
	hs.taskService.SetRequester(hs.requester)
	th := new(TestTaskHandler)
	th.responsesChan = hs.taskService.ResponsesChan
	th.tasksChan = hs.taskService.TasksChan
	th.store = hs.taskService.Store
	th.Init(th.WorkerTest, th.SaverTest, config)
	th.response = NewResponse()
	hs.taskHandler = th
	hs.taskService.SetTaskHandler(th)
	response := new(models.TaskResponse)
	hs.requester.SetResponse(response)
	httpClient := &http.Client{Transport: util2.NewTransport(hs.taskService.Server.GetHandler())}
	c := client2.NewWithClient(client.DefaultHost, client.DefaultBasePath, client.DefaultSchemes, httpClient)
	hs.taskClient = client.New(c, nil)
	hs.task.Path = "https://123445"
}

type HandlersSuit struct {
	task        models2.FetchTask
	taskService *internal.TaskService
	taskClient  *client.FetchtaskHandlingService
	requester   *TestRequester
	taskHandler *TestTaskHandler
	suite.Suite
}

func (hs *HandlersSuit) TestAddFetchTask() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Path: hs.task.Path,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
}

func (hs *HandlersSuit) TestGetTaskError() {
	hs.taskHandler.response = NewResponseWithErr()
	hs.taskHandler.SetTaskStatus(models.StatusError)
	getResponseOK, errResp := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{

		Path: hs.task.Path,
	}))
	hs.Require().NotNil(getResponseOK)
	hs.Require().NoError(errResp)
	hs.Require().EqualValues(models.StatusNew, getResponseOK.Payload.Progress)
	getTaskOk, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(getResponseOK.Payload.ID))
	hs.Require().NoError(err)
	hs.Require().EqualValues(models.StatusError, getTaskOk.Payload.Request.Progress)
}

func (hs *HandlersSuit) TestGetTaskWithError() {
	hs.taskHandler.response = NewResponseWithErr()
	hs.taskHandler.SetTaskStatus(models.StatusError)
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Path: hs.task.Path,
	}))
	hs.Require().NotNil(getResponseOK)
	hs.Require().NoError(err)
	hs.Require().EqualValues(models.StatusNew, getResponseOK.Payload.Progress)
	getTaskOk, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(getResponseOK.Payload.ID))
	hs.Require().NoError(err)
	hs.Require().Nil(getTaskOk.Payload.Response)
}

func (hs *HandlersSuit) TestGetFetchTask() {
	hs.taskHandler.SetTaskStatus(models.StatusCompleted)
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Path: hs.task.Path,
	}))
	hs.Require().NotNil(getResponseOK)
	hs.Require().NoError(err)
	hs.Require().EqualValues(models.StatusNew, getResponseOK.Payload.Progress)
	getTaskOK1, err := hs.taskClient.Operations.GetTask(operations.NewGetTaskParams().WithTaskID(getResponseOK.Payload.ID))
	hs.Require().NoError(err)
	hs.Require().NotNil(getTaskOK1)
	hs.EqualValues(getResponseOK.Payload.ID, getTaskOK1.Payload.Request.ID)
}

func (hs *HandlersSuit) TestGetAllFetchTasks() {
	getResponseOK, err := hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Path: hs.task.Path,
	}))
	hs.Require().NoError(err)
	hs.Require().NotNil(getResponseOK)
	getResponseOK, err = hs.taskClient.Operations.CreateFetchTask(operations.NewCreateFetchTaskParams().WithTask(operations.CreateFetchTaskBody{
		Path: hs.task.Path,
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
		Path: hs.task.Path,
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
