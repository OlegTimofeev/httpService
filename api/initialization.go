package api

import (
	"github.com/labstack/echo"
	"httpService/data_base"
)

func InitializeServer() (string, *echo.Echo) {
	data_base.InitDB()
	port := ":8080"
	r := echo.New()
	r.GET("/sendTask", GetResponse)
	r.GET("/getTasks", GetTasks)
	r.DELETE("/delete/:ftId", DeleteFT)
	r.GET("/getTask/:ftId", GetTask)
	return port, r
}
