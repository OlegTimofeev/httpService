package api

import (
	"encoding/json"
	"github.com/labstack/echo"
	"httpService/data_base"
	"httpService/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetResponse(c echo.Context) error {
	ft := new(models.FetchTask)
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
	ur := models.UserResponse{
		Headers:    resp.Header,
		HttpStatus: resp.StatusCode,
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	bodyString := string(body)
	ur.BodyLen = len(bodyString)
	ft, err = data_base.GetDB().AddFetchTask(ft)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ur.FetchTaskId = ft.ID
	return c.JSON(http.StatusOK, ur)
}

func GetTasks(c echo.Context) error {
	tasks, err := data_base.GetDB().GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, tasks)
}

func DeleteFT(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("ftId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = data_base.GetDB().DeleteFetchTask(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, "Operation : successful")
}

func GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("ftId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	task, err := data_base.GetDB().GetFetchTask(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, task)
}
