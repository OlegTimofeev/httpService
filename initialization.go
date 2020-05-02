package main

import (
	"github.com/labstack/echo"
)

var dataStore *MapStore

func initDB() {
	dataStore = new(MapStore)
	dataStore.tasks = make(map[int]*FetchTask)
	dataStore.taskID = 0
}

func initializeServer() (string, *echo.Echo) {
	initDB()
	port := ":8080"
	r := echo.New()
	r.GET("/sendTask", getResponse)
	r.GET("/getTasks", getTasks)
	r.DELETE("/delete/:ftId", deleteFT)
	r.GET("/getTask/:ftId", getTask)
	return port, r
}
