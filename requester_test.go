package httpService

import (
	"httpService/internal/models"
)

type TestRequester struct {
	response *models.TaskResponse
	err      error
}

func (requester *TestRequester) DoRequest(task *models.FetchTask) (*models.TaskResponse, error) {
	return requester.response, requester.err
}

func (requester *TestRequester) SetResponse(reqResponse *models.TaskResponse) {
	requester.response = reqResponse
}

func (requester *TestRequester) SetError(err error) {
	requester.err = err
}
