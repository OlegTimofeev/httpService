package taskHandler

import (
	"httpService/internal/dataBase"
	"httpService/internal/models"
)

func (h *ProdTaskHandler) Init(worker func(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse), saver func(responses <-chan *models.TaskResponse), config dataBase.ConfigDB) {
	for i := 0; i < config.PoolSize; i++ {
		go worker(h.TasksChan, h.ResponsesChan)
		go saver(h.ResponsesChan)
	}
}

func NewProdTaskHandler(ftChan chan *models.FetchTask, trChan chan *models.TaskResponse) *ProdTaskHandler {
	return &ProdTaskHandler{
		TasksChan:     ftChan,
		ResponsesChan: trChan,
	}
}
