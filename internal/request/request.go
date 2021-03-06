package request

import (
	"httpService/internal/models"
	"net/http"
)

type Requester interface {
	DoRequest(task *models.FetchTask) (*models.TaskResponse, error)
}

type HTTPRequester struct {
	client http.Client
}
