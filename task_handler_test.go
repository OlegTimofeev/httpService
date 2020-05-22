package httpService

import (
	"httpService/internal/dataBase"
	"httpService/internal/models"
)

type TestTaskHandler struct {
	tasksChan     chan *models.FetchTask
	responsesChan chan *models.TaskResponse
	haveError     bool
	store         dataBase.DataStore
}

func (h *TestTaskHandler) Init(worker func(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse), saver func(responses <-chan *models.TaskResponse), config dataBase.ConfigDB) {
	for i := 0; i < config.PoolSize; i++ {
		go worker(h.tasksChan, h.responsesChan)
		go saver(h.responsesChan)
	}
}

func (h *TestTaskHandler) HaveError(f bool) {
	h.haveError = f
}

func (h *TestTaskHandler) SetStore(s dataBase.DataStore) {
	h.store = s
}

func (h *TestTaskHandler) WorkerTest(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse) {
	for task := range tasks {
		if h.haveError {
			task.Status = models.StatusError
			h.store.UpdateFetchTask(*task)
			response <- &models.TaskResponse{
				ID:  task.ID,
				Err: "err",
			}
			return
		}
		response <- &models.TaskResponse{
			ID: task.ID,
		}
		task.Status = models.StatusCompleted
		h.store.UpdateFetchTask(*task)
	}
}

func (h *TestTaskHandler) SaverTest(responses <-chan *models.TaskResponse) {
	for response := range responses {
		h.store.AddTaskResponse(response)
	}
}
