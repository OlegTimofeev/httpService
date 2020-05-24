package internal

import (
	"github.com/go-openapi/runtime/middleware"
	"httpService/internal/models"
	models2 "httpService/service/models"
	"httpService/service/restapi/operations"
	"net/http"
)

func CreateFetchTask(params operations.CreateFetchTaskParams) middleware.Responder {
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
	return operations.NewCreateFetchTaskOK().WithPayload(ft.ConvertToSwaggerModel())
}

func GetTasks(params operations.GetAllTasksParams) middleware.Responder {
	tasks, err := taskService.Store.GetAllTasks()
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to get tasks from database")
	}
	tasksResp := make([]*models2.FetchTask, len(tasks))
	for i := 0; i < len(tasks); i++ {
		tasksResp[i] = tasks[i].ConvertToSwaggerModel()
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
	if task.Status == models.StatusInProgress {
		return operations.NewGetTaskOK().WithPayload(&models2.FullTask{
			Request: ConvertToRequest(task.ConvertToSwaggerModel())})
	}
	if task.Status == models.StatusError {
		return operations.NewGetTaskOK().WithPayload(&models2.FullTask{
			Request: ConvertToRequest(task.ConvertToSwaggerModel()),
		})
	}
	resp, err := taskService.Store.GetTaskResponseByFtID(id)
	if err != nil {
		return middleware.Error(http.StatusNotFound, "Error : Unable to get tasks from database")
	}
	return operations.NewGetTaskOK().WithPayload(&models2.FullTask{
		Request:  ConvertToRequest(task.ConvertToSwaggerModel()),
		Response: ConvertToResponse(resp.ConvertToSwaggerModel()),
	})
}

func Worker(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse) {
	for task := range tasks {
		res, err := taskService.Requester.DoRequest(*task)
		if err != nil {
			task.Status = models.StatusError
			taskService.Store.UpdateFetchTask(*task)
			response <- &models.TaskResponse{
				ID:  task.ID,
				Err: err.Error(),
			}
			return
		}
		res.ID = task.ID
		response <- res
		task.Status = models.StatusCompleted
		taskService.Store.UpdateFetchTask(*task)
	}
}

func Saver(responses <-chan *models.TaskResponse) {
	for response := range responses {
		taskService.Store.AddTaskResponse(response)
	}
}

func ConvertToResponse(response *models2.TaskResponse) *models2.FullTaskResponse {
	if response == nil {
		return nil
	}
	return &models2.FullTaskResponse{
		BodyLenght: response.BodyLenght,
		HTTPStatus: response.HTTPStatus,
	}
}

func ConvertToRequest(request *models2.FetchTask) *models2.FullTaskRequest {
	if request == nil {
		return nil
	}
	return &models2.FullTaskRequest{
		ID:       request.ID,
		Progress: request.Progress,
		Body:     request.Body,
		Path:     request.Path,
		Method:   request.Method,
		Headers:  request.Headers,
	}
}
