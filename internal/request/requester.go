package request

import (
	"httpService/internal/models"
	models2 "httpService/service/models"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func NewRequester() *HTTPRequester {
	rqst := new(HTTPRequester)
	rqst.client = http.Client{
		Timeout: 15 * time.Second,
	}
	return rqst
}

func (requester *HTTPRequester) DoRequest(ft models.FetchTask) (*models2.TaskResponse, error) {
	req, err := http.NewRequest(ft.Method, ft.Path, strings.NewReader(ft.Body))
	if err != nil {
		return nil, err
	}
	req.Header = ft.Headers
	resp, err := requester.client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	taskResponse := models2.TaskResponse{
		ID:         int64(ft.ID),
		HTTPStatus: int64(resp.StatusCode),
		Method:     ft.Method,
		Path:       ft.Path,
		BodyLenght: int64(len(bodyString)),
		Headers:    resp.Header,
	}
	return &taskResponse, nil
}
