package request

import (
	"httpService/internal/models"
	models2 "httpService/service/models"
)

func (requester *TestRequester) DoRequest(ft models.FetchTask) (*models2.TaskResponse, error) {
	return requester.response, requester.err
}

func (requester *TestRequester) SetResponse(reqResponse *models2.TaskResponse) {
	requester.response = reqResponse
}

func (requester *TestRequester) SetError(err error) {
	requester.err = err
}
