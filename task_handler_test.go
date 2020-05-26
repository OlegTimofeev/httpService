package httpService

import (
	"httpService/internal/dataBase"
	"httpService/internal/models"
)

type TestTaskHandler struct {
	tasksChan     chan *models.FetchTask
	responsesChan chan *models.TaskResponse
	error         string
	response      *models.TaskResponse
	taskStatus    string
	store         dataBase.DataStore
}

func (h *TestTaskHandler) Init(worker func(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse), saver func(responses <-chan *models.TaskResponse), config dataBase.ConfigDB) {
	for i := 0; i < config.PoolSize; i++ {
		go worker(h.tasksChan, h.responsesChan)
		go saver(h.responsesChan)
	}
}

func (h *TestTaskHandler) SetError(err string) {
	h.error = err
}

func (h *TestTaskHandler) SetTaskStatus(status string) {
	h.taskStatus = status
}

func (h *TestTaskHandler) SetResponse(res *models.TaskResponse) {
	h.response = res
}

func (h *TestTaskHandler) SetStore(s dataBase.DataStore) {
	h.store = s
}

func (h *TestTaskHandler) WorkerTest(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse) {
	for task := range tasks {
		h.response.ID = task.ID
		response <- h.response
		task.Status = h.taskStatus
		h.store.UpdateFetchTask(*task)
	}
}

func (h *TestTaskHandler) SaverTest(responses <-chan *models.TaskResponse) {
	for response := range responses {
		h.store.AddTaskResponse(response)
	}
}

func NewResponseWithErr() *models.TaskResponse {
	return &models.TaskResponse{
		Err: "error",
	}
}

func NewResponse() *models.TaskResponse {
	return &models.TaskResponse{}
}
