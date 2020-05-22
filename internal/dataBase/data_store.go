package dataBase

import (
	"github.com/go-pg/pg"
	"httpService/internal/models"
	"sync"
)

type DataStore interface {
	AddFetchTask(task *models.FetchTask) (*models.FetchTask, error)
	DeleteFetchTask(taskId int) error
	GetAllTasks() ([]*models.FetchTask, error)
	GetFetchTask(taskId int) (*models.FetchTask, error)
	GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error)
	AddTaskResponse(res *models.TaskResponse) (*models.TaskResponse, error)
	UpdateFetchTask(task models.FetchTask) error
}

type MapStore struct {
	Tasks  map[int]*models.FetchTask
	mutex  sync.Mutex
	TaskID int
}

type PostgresDB struct {
	pgdb *pg.DB
}

func NewDataStore(config ConfigDB) DataStore {
	if config.StoreType == "postgres" {
		return NewPGStore(config)
	} else {
		return NewMapStore()
	}
}
