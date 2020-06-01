package taskWorker

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"httpService/internal/dataBase"
	"httpService/internal/models"
	"httpService/internal/request"
	"httpService/nats"
)

type ChanWorkerPool struct {
	requester request.Requester
	r         models.CanSetResponse
	sc        *nats.Queue
}

func (h *ChanWorkerPool) AddRequest(ft *models.FetchTask) {
	h.sc.Publish(*ft)
}

func (h *ChanWorkerPool) ListenForTasks() {
	h.sc.Subscribe(h.msgHandler)
}

func NewWorkerPool(requester request.Requester, r models.CanSetResponse, config dataBase.ConfigDB) *ChanWorkerPool {
	pool := &ChanWorkerPool{
		requester: requester,
		r:         r,
		sc:        nats.NewStan(config),
	}
	pool.ListenForTasks()
	return pool
}

func (h *ChanWorkerPool) msgHandler(msg *stan.Msg) {
	var ft models.FetchTask
	json.Unmarshal(msg.Data, &ft)
	res, err := h.requester.DoRequest(&ft)
	h.r.SetResponse(ft.ID, res, err)
}
