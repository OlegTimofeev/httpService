package dataBase

import (
	"errors"
	"httpService/internal/models"
)

func NewMapStore() *MapStore {
	ms := new(MapStore)
	ms.Tasks = make(map[int]*models.FetchTask)
	ms.TaskID = 0
	return ms
}

func (ms *MapStore) AddFetchTask(ft *models.FetchTask) (*models.FetchTask, error) {
	ms.mutex.Lock()
	ft.ID = ms.GetTaskID()
	ms.Tasks[ft.ID] = ft
	ms.mutex.Unlock()
	return ft, nil
}

func (ms *MapStore) DeleteFetchTask(taskId int) error {
	ms.mutex.Lock()
	task := ms.Tasks[taskId]
	if task == nil {
		ms.mutex.Unlock()
		return errors.New("task not found")
	}
	delete(ms.Tasks, taskId)
	ms.mutex.Unlock()
	return nil
}

func (ms *MapStore) GetAllTasks() ([]*models.FetchTask, error) {
	var allTasks []*models.FetchTask
	ms.mutex.Lock()
	for _, ft := range ms.Tasks {
		allTasks = append(allTasks, ft)
	}
	ms.mutex.Unlock()
	return allTasks, nil
}
func (ms *MapStore) GetFetchTask(taskId int) (*models.FetchTask, error) {
	ms.mutex.Lock()
	task := ms.Tasks[taskId]
	ms.mutex.Unlock()
	if task == nil {
		return nil, errors.New("task not found")
	}
	return ms.Tasks[taskId], nil
}

func (ms *MapStore) GetTaskID() int {
	ms.mutex.Lock()
	ms.TaskID += 1
	ms.mutex.Unlock()
	return ms.TaskID
}
