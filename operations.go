package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getResponse(c echo.Context) error {
	ft := new(FetchTask)
	if err := json.NewDecoder(c.Request().Body).Decode(ft); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	req, err := http.NewRequest(ft.Method, ft.Path, strings.NewReader(ft.Body))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	req.Header = ft.Headers
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ur := UserResponse{
		Headers:    resp.Header,
		HttpStatus: resp.StatusCode,
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	bodyString := string(body)
	ur.BodyLen = len(bodyString)
	ft, err = dataStore.addFetchTask(ft)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ur.FetchTaskId = ft.ID
	return c.JSON(http.StatusOK, ur)
}

func getTasks(c echo.Context) error {
	tasks, err := dataStore.getAllTasks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, tasks)
}

func deleteFT(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("ftId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = dataStore.deleteFetchTask(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Operation : successful")
}

func getTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("ftId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	task, err := dataStore.getFetchTask(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, task)
}
