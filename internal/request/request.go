package request

import (
	"httpService/internal/models"
	models2 "httpService/service/models"
	"net/http"
)

type Requester interface {
	DoRequest(task models.FetchTask) (*models2.TaskResponse, error)
}

type HTTPRequester struct {
	client http.Client
}
