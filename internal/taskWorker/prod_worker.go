package taskWorker

import (
	"httpService/internal/models"
	"httpService/internal/request"
)

type ChanWorkerPool struct {
	requester request.Requester

	tasksChan chan *taskWithStore
}

type taskWithStore struct {
	FetchTask *models.FetchTask
	r         models.CanSetResponse
}

func (h *ChanWorkerPool) AddRequest(ft *models.FetchTask, r models.CanSetResponse) {
	h.tasksChan <- &taskWithStore{
		FetchTask: ft,
		r:         r,
	}
}

func (h *ChanWorkerPool) ListenForTasks() {
	for tr := range h.tasksChan {
		res, err := h.requester.DoRequest(tr.FetchTask)
		tr.r.SetResponse(tr.FetchTask.ID, res, err)
	}
}

func NewWorkerPool(requester request.Requester) *ChanWorkerPool {
	pool := &ChanWorkerPool{
		requester: requester,
		tasksChan: make(chan *taskWithStore, 1000),
	}
	go pool.ListenForTasks()
	return pool
}
