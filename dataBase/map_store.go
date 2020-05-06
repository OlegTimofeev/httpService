package dataBase

import (
	"errors"
	"httpService/models"
)

func (ms *MapStore) AddFetchTask(ft *models.FetchTask) (*models.FetchTask, error) {
	ft.ID = ms.GetTaskID()
	ms.Tasks[ft.ID] = ft
	return ft, nil
}

func (ms *MapStore) DeleteFetchTask(taskId int) error {
	task := ms.Tasks[taskId]
	if task == nil {
		return errors.New("task not found")
	}
	delete(ms.Tasks, taskId)
	return nil
}

func (ms *MapStore) GetAllTasks() ([]*models.FetchTask, error) {
	var allTasks []*models.FetchTask
	for _, ft := range ms.Tasks {
		allTasks = append(allTasks, ft)
	}
	return allTasks, nil
}
func (ms *MapStore) GetFetchTask(taskId int) (*models.FetchTask, error) {
	task := ms.Tasks[taskId]
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
