package api

import (
	"github.com/go-openapi/runtime/middleware"
	"httpService/models"
	models2 "httpService/service/models"
	"httpService/service/restapi/operations"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetResponse(params operations.CreateFetchTaskParams) middleware.Responder {
	if err := db.CheckConnection(); err != nil {
		return middleware.Error(500, "Error : Connection lost")
	}
	ft := new(models.FetchTask)
	ft.Body = params.Task.Body
	ft.Path = params.Task.Path
	ft.Headers = params.Task.Headers
	ft.Method = params.Task.Method
	req, err := http.NewRequest(ft.Method, params.Task.Path, strings.NewReader(params.Task.Body))
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable produce request")
	}
	req.Header = params.Task.Headers
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return middleware.Error(http.StatusNotFound, "Error : Page not found")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to read response body")
	}
	bodyString := string(body)
	ft, err = db.AddFetchTask(ft)
	if err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Unable to add task to database")
	}
	taskResponse := models2.TaskResponse{
		ID:         int64(ft.ID),
		HTTPStatus: int64(resp.StatusCode),
		Method:     ft.Method,
		Path:       ft.Path,
		BodyLenght: int64(len(bodyString)),
		Headers:    resp.Header,
	}
	return operations.NewCreateFetchTaskOK().WithPayload(&taskResponse)
}

func GetTasks() middleware.Responder {
	if err := db.CheckConnection(); err != nil {
		return middleware.Error(500, "Error : Connection lost")
	}
	tasks, err := db.GetAllTasks()
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
	if err := db.CheckConnection(); err != nil {
		return middleware.Error(http.StatusInternalServerError, "Error : Connection lost")
	}
	id := int(params.TaskID)
	if err := db.DeleteFetchTask(id); err != nil {
		return middleware.Error(http.StatusNotFound, "Error : Unable to delete task from database")
	}
	return operations.NewDeleteFetchTaskOK()
}

func GetTask(params operations.GetTaskParams) middleware.Responder {
	if err := db.CheckConnection(); err != nil {
		return middleware.Error(500, "Error : Connection lost")
	}
	id := int(params.TaskID)
	task, err := db.GetFetchTask(id)
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
