package main

import "sync"

type DataStore interface {
	addFetchTask(task *FetchTask) (*FetchTask, error)
	deleteFetchTask(taskId int) error
	getAllTasks() ([]*FetchTask, error)
}

type MapStore struct {
	tasks  map[int]*FetchTask
	mutex  sync.Mutex
	taskID int
}
