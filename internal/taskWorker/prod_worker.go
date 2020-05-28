package taskWorker

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"httpService/internal/dataBase"
	"httpService/internal/models"
	"httpService/internal/request"
	"httpService/nats"
	"log"
)

type ChanWorkerPool struct {
	requester request.Requester
	r         models.CanSetResponse
	tasksChan chan *taskWithStore
	sc        *nats.Queue
}

type taskWithStore struct {
	FetchTask *models.FetchTask
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
		tasksChan: make(chan *taskWithStore, 1000),
	}
	go pool.ListenForTasks()
	return pool
}

func (h *ChanWorkerPool) msgHandler(msg *stan.Msg) {
	var ft models.FetchTask
	if err := json.Unmarshal(msg.Data, &ft); err != nil {
		log.Fatal(err)
	}
	res, err := h.requester.DoRequest(&ft)
	h.r.SetResponse(ft.ID, res, err)
	fmt.Println(ft)
}

func (h *ChanWorkerPool) Close() {
	h.sc.Sc.Close()
}
