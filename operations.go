package main

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"strings"
)

func getResponse(c echo.Context) error {
	ft := new(fetchTask)
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
	ur := userResponse{
		Headers:    resp.Header,
		HttpStatus: resp.StatusCode,
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	bodyString := string(body)
	ur.Body = bodyString
	ur.BodyLen = len(bodyString)
	return c.JSON(http.StatusOK, ur)
}

func getTasks(c echo.Context) error {
	return nil
}

func deleteFT(c echo.Context) error {
	return nil
}

func useTask(c echo.Context) error {
	return nil
}
