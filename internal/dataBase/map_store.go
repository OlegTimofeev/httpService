package dataBase

import (
	"errors"
	"httpService/internal/models"
	"sync"
)

type MapStore struct {
	Tasks     map[int]*models.FetchTask
	mutex     sync.Mutex
	TaskID    int
	Responses map[int]*models.TaskResponse
}

func NewMapStore() *MapStore {
	ms := new(MapStore)
	ms.Tasks = make(map[int]*models.FetchTask)
	ms.Responses = make(map[int]*models.TaskResponse)
	ms.TaskID = 0
	return ms
}

func (ms *MapStore) AddFetchTask(ft *models.FetchTask) (*models.FetchTask, error) {
	ft.ID = ms.GetTaskID()
	_, err := ms.AddTaskResponse(&models.TaskResponse{
		ID: ft.ID,
	})
	if err != nil {
		return nil, err
	}
	ms.Tasks[ft.ID] = ft
	return ft, nil
}

func (ms *MapStore) DeleteFetchTask(taskId int) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	task := ms.Tasks[taskId]
	if task == nil {
		return errors.New("task not found")
	}
	delete(ms.Responses, taskId)
	delete(ms.Tasks, taskId)
	return nil
}

func (ms *MapStore) GetAllTasks() ([]*models.FetchTask, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	var allTasks []*models.FetchTask
	for _, ft := range ms.Tasks {
		allTasks = append(allTasks, ft)
	}

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
	defer ms.mutex.Unlock()
	ms.TaskID += 1
	return ms.TaskID
}

func (ms *MapStore) GetTaskResponseByFtID(taskId int) (*models.TaskResponse, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	if ms.Tasks[taskId] == nil {
		return nil, errors.New("no task for current id" + string(taskId))
	}
	return ms.Responses[taskId], nil
}
func (ms *MapStore) AddTaskResponse(res *models.TaskResponse) (*models.TaskResponse, error) {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	return res, nil
}

func (ms *MapStore) UpdateFetchTask(task models.FetchTask) error {
	ms.Tasks[task.ID] = &task
	return nil
}

func (ms *MapStore) SetResponse(id int, response *models.TaskResponse, err error) error {
	ms.mutex.Lock()
	defer ms.mutex.Unlock()
	task := ms.Tasks[id]
	task.Status = models.StatusInProgress
	if err := ms.UpdateFetchTask(*task); err != nil {
		return err
	}
	response.ID = id
	if err != nil {
		task.Status = models.StatusError
	} else {
		task.Status = models.StatusCompleted
	}
	ms.Tasks[id] = task
	ms.Responses[id] = response
	return nil
}
