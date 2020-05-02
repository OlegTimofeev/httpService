package main

import "errors"

func (ms *MapStore) addFetchTask(ft *FetchTask) (*FetchTask, error) {
	ft.ID = ms.GetTaskID()
	ms.tasks[ft.ID] = ft
	return ft, nil
}

func (ms *MapStore) deleteFetchTask(taskId int) error {
	task := ms.tasks[taskId]
	if task == nil {
		return errors.New("task not found")
	}
	delete(ms.tasks, taskId)
	return nil
}

func (ms *MapStore) getAllTasks() ([]*FetchTask, error) {
	var allTasks []*FetchTask
	for _, ft := range ms.tasks {
		allTasks = append(allTasks, ft)
	}
	return allTasks, nil
}
func (ms *MapStore) getFetchTask(taskId int) (*FetchTask, error) {
	task := ms.tasks[taskId]
	if task == nil {
		return nil, errors.New("task not found")
	}
	return ms.tasks[taskId], nil
}

func (ms *MapStore) GetTaskID() int {
	ms.mutex.Lock()
	ms.taskID += 1
	ms.mutex.Unlock()
	return ms.taskID
}
