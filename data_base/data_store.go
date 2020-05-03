package data_base

import (
	"httpService/models"
	"sync"
)

type DataStore interface {
	AddFetchTask(task *models.FetchTask) (*models.FetchTask, error)
	DeleteFetchTask(taskId int) error
	GetAllTasks() ([]*models.FetchTask, error)
}

type MapStore struct {
	Tasks  map[int]*models.FetchTask
	mutex  sync.Mutex
	TaskID int
}
