package internal

import (
	"github.com/go-openapi/runtime/middleware"
	"httpService/internal/models"
	models2 "httpService/service/models"
	"httpService/service/restapi/operations"
	"net/http"
)

func GetResponse(params operations.CreateFetchTaskParams) middleware.Responder {
	ft := new(models.FetchTask)
	ft.Body = params.Task.Body
	ft.Path = params.Task.Path
	ft.Headers = params.Task.Headers
	ft.Method = params.Task.Method
	ft, err := taskService.Store.AddFetchTask(ft)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to add tasks to database")
	}
	taskResponse, err := taskService.Requester.DoRequest(*ft)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get response")
	}
	return operations.NewCreateFetchTaskOK().WithPayload(taskResponse)
}

func GetTasks(params operations.GetAllTasksParams) middleware.Responder {
	tasks, err := taskService.Store.GetAllTasks()
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get tasks from database")
	}
	tasksResp := make([]*models2.FetchTask, len(tasks))
	for i := 0; i < len(tasks); i++ {
		tasksResp[i] = convertForResp(tasks[i])
	}
	return operations.NewGetAllTasksOK().WithPayload(tasksResp)
}

func DeleteFT(params operations.DeleteFetchTaskParams) middleware.Responder {
	id := int(params.TaskID)
	if err := taskService.Store.DeleteFetchTask(id); err != nil {
		return middleware.Error(http.StatusNotFound, "Error : Unable to delete task from database")
	}
	return operations.NewDeleteFetchTaskOK()
}

func GetTask(params operations.GetTaskParams) middleware.Responder {
	id := int(params.TaskID)
	task, err := taskService.Store.GetFetchTask(id)
	if err != nil {
		return middleware.Error(http.StatusNotFound, "Error : Unable to get tasks from database")
	}
	return operations.NewGetTaskOK().WithPayload(convertForResp(task))
}

func convertForResp(task *models.FetchTask) *models2.FetchTask {
	ftResp := new(models2.FetchTask)
	ftResp.Method = task.Method
	ftResp.ID = int64(task.ID)
	ftResp.Body = task.Body
	ftResp.Headers = task.Headers
	ftResp.Path = task.Path
	return ftResp
}
