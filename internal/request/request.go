package request

import (
	"httpService/internal/models"
	models2 "httpService/service/models"
	"net/http"
)

type RequesterModel interface {
	DoRequest(task models.FetchTask) (*models2.TaskResponse, error)
}

type Request struct {
	client http.Client
}

type TestRequester struct {
	response *models2.TaskResponse
	err      error
}
