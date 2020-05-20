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
	ft.Status = models.StatusNew
	ft, err := taskService.Store.AddFetchTask(ft)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to add tasks to database")
	}
	taskService.TasksChan <- ft
	err = <-taskService.ErrorsChan
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get response")
	}
	taskResponse, err := taskService.Store.GetTaskResponseByFtID(ft.ID)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get response")
	}
	ft.Status = models.StatusCompleted
	if err = taskService.Store.UpdateFetchTask(ft); err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get response")
	}
	return operations.NewCreateFetchTaskOK().WithPayload(taskResponse.ConvertForResp())
}

func GetTasks(params operations.GetAllTasksParams) middleware.Responder {
	tasks, err := taskService.Store.GetAllTasks()
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get tasks from database")
	}
	tasksResp := make([]*models2.FetchTask, len(tasks))
	for i := 0; i < len(tasks); i++ {
		tasksResp[i] = tasks[i].ConvertForResp()
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
	return operations.NewGetTaskOK().WithPayload(task.ConvertForResp())
}

func Worker(tasks chan *models.FetchTask, response chan *models.TaskResponse, errors chan error) {
	for task := range tasks {
		res, err := taskService.Requester.DoRequest(*task)
		if err != nil {
			errors <- err
			return
		}
		task.Status = models.StatusInProgress
		if err = taskService.Store.UpdateFetchTask(task); err != nil {
			errors <- err
			return
		}
		res.FetchTaskID = task.ID
		response <- res
	}
}

func Saver(response chan *models.TaskResponse, errors chan error) {
	_, err := taskService.Store.AddTaskResponse(<-response)
	if err != nil {
		errors <- err
		return
	}
	errors <- nil
}
