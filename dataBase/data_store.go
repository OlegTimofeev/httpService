package dataBase

import (
	"httpService/models"
	"sync"
)

type DataStore interface {
	addFetchTask(task *models.FetchTask) (*models.FetchTask, error)
	deleteFetchTask(taskId int) error
	getAllTasks() ([]*models.FetchTask, error)
}

type MapStore struct {
	Tasks  map[int]*models.FetchTask
	mutex  sync.Mutex
	TaskID int
}
