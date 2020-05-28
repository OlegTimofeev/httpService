package dataBase

import (
	"httpService/internal/models"
)

type DataStore interface {
	models.CanSetResponse

	AddFetchTask(task *models.FetchTask) (*models.FetchTask, error)
	DeleteFetchTask(taskId int) error
	GetAllTasks() ([]*models.FetchTask, error)
	GetFetchTask(taskId int) (*models.FetchTask, error)
	GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error)
	AddTaskResponse(res *models.TaskResponse) (*models.TaskResponse, error)
	UpdateFetchTask(task models.FetchTask) error
}

func NewDataStore(config ConfigDB) DataStore {
	if config.StoreType == "postgres" {
		return NewPGStore(config)
	} else {
		return NewMapStore()
	}
}
