package api

import (
	"github.com/labstack/echo"
	"httpService/dataBase"
)

func InitializeServer() (string, *echo.Echo) {
	dataBase.InitDB()
	port := ":8080"
	r := echo.New()
	r.GET("/sendTask", GetResponse)
	r.GET("/getTasks", GetTasks)
	r.DELETE("/delete/:ftId", DeleteFT)
	r.GET("/getTask/:ftId", GetTask)
	return port, r
}
