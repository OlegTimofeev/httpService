package main

import (
	"github.com/labstack/echo"
)

func initializeServer() (string, *echo.Echo) {
	port := ":8080"
	r := echo.New()
	r.GET("/sendTask", getResponse)
	r.GET("/getTasks", getTasks)
	r.DELETE("/delete/:ftId", deleteFT)
	return port, r
}
