package dataBase

import (
	"github.com/go-pg/pg"
	"httpService/models"
	"sync"
)

type DataStore interface {
	InitDB()
	AddFetchTask(task *models.FetchTask) (*models.FetchTask, error)
	DeleteFetchTask(taskId int) error
	GetAllTasks() ([]*models.FetchTask, error)
	GetFetchTask(taskId int) (*models.FetchTask, error)
	CheckConnection() error
}

type MapStore struct {
	Tasks  map[int]*models.FetchTask
	mutex  sync.Mutex
	TaskID int
}

type PostgresDB struct {
	pgdb *pg.DB
}
