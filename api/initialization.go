package api

import (
	"github.com/labstack/echo"
	"httpService/dataBase"
)

var db dataBase.DataStore

func InitializeServer() (string, *echo.Echo) {
	db = new(dataBase.PostgresDB)
	db.InitDB()
	port := ":8080"
	r := echo.New()
	r.POST("/sendTask", GetResponse)
	r.GET("/getTasks", GetTasks)
	r.DELETE("/delete/:ftId", DeleteFT)
	r.GET("/getTask/:ftId", GetTask)
	return port, r
}
