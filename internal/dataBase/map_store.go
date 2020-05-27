package dataBase

import (
	"errors"
	"httpService/internal/models"
	"sync"
)

type MapStore struct {
	Tasks  map[int]*models.FetchTask
	mutex  sync.Mutex
	TaskID int
}

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
	defer ms.mutex.Unlock()
	task := ms.Tasks[taskId]
	if task == nil {
		return errors.New("task not found")
	}
	delete(ms.Tasks, taskId)
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
	defer ms.mutex.Unlock()
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

func (ms *MapStore) GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error) {
	return nil, nil
}
func (ms *MapStore) AddTaskResponse(res *models.TaskResponse) (*models.TaskResponse, error) {
	return nil, nil
}

func (ms *MapStore) UpdateFetchTask(task models.FetchTask) error {
	return nil
}

func (ms *MapStore) SetResponse(id int, response *models.TaskResponse, err error) error {
	return nil
}
