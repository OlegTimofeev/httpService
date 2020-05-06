package dataBase

import "httpService/models"

var dataStore *MapStore

func InitDB() *MapStore {
	dataStore = new(MapStore)
	dataStore.Tasks = make(map[int]*models.FetchTask)
	dataStore.TaskID = 0
	return dataStore
}

func GetDB() *MapStore {
	return dataStore
}
