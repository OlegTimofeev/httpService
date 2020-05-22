package taskHandler

import (
	"httpService/internal/dataBase"
	"httpService/internal/models"
)

type TaskHandler interface {
	Init(worker func(tasks <-chan *models.FetchTask, response chan<- *models.TaskResponse), saver func(responses <-chan *models.TaskResponse), config dataBase.ConfigDB)
}

type ProdTaskHandler struct {
	TasksChan     chan *models.FetchTask
	ResponsesChan chan *models.TaskResponse
}
