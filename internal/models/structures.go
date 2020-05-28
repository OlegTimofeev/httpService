package models

import models2 "httpService/service/models"

const (
	StatusNew        = "New"
	StatusInProgress = "InProgress"
	StatusCompleted  = "Completed"
	StatusError      = "Error"
)

type FetchTask struct {
	ID      int                 `json:"task_id"`
	Method  string              `json:"method"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Body    string              `json:"body"`
	Status  string              `json:"status"`
}

func (task *FetchTask) ConvertToSwaggerModel() *models2.FetchTask {
	return &models2.FetchTask{
		Method:   task.Method,
		ID:       int64(task.ID),
		Body:     task.Body,
		Progress: task.Status,
		Headers:  task.Headers,
		Path:     task.Path}
}

type TaskResponse struct {
	ID      int    `json:"resp_id"`
	Status  int    `json:"status"`
	BodyLen int    `json:"body_len"`
	Err     string `json:"err"`
}

func (tr *TaskResponse) ConvertToSwaggerModel() *models2.TaskResponse {
	return &models2.TaskResponse{
		ID:         int64(tr.ID),
		BodyLenght: int64(tr.BodyLen),
		HTTPStatus: int64(tr.Status),
	}
}
