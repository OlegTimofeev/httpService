package models

import models2 "httpService/service/models"

const (
	StatusNew        = "New"
	StatusInProgress = "InProgress"
	StatusCompleted  = "Completed"
)

type FetchTask struct {
	ID      int                 `json:"task_id"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
	Status  string              `json:"status"`
}

func (task *FetchTask) ConvertForResp() *models2.FetchTask {
	ftResp := new(models2.FetchTask)
	ftResp.Method = task.Method
	ftResp.ID = int64(task.ID)
	ftResp.Body = task.Body
	ftResp.Headers = task.Headers
	ftResp.Path = task.Path
	return ftResp
}

type TaskResponse struct {
	ID          int                 `json:"resp_id"`
	Status      int                 `json:"status"`
	Method      string              `json:"method"`
	Path        string              `json:"path"`
	Headers     map[string][]string `json:"headers"`
	BodyLen     int                 `json:"body_len"`
	FetchTaskID int                 `json:"ft_id"`
}

func (tr *TaskResponse) ConvertForResp() *models2.TaskResponse {
	resp := new(models2.TaskResponse)
	resp.ID = int64(tr.FetchTaskID)
	resp.BodyLenght = int64(tr.BodyLen)
	resp.HTTPStatus = int64(tr.Status)
	resp.Path = tr.Path
	resp.Method = tr.Method
	resp.Headers = tr.Headers
	return resp
}
