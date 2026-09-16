package main

import (
	"math/rand"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
)

var sampleTasks = []Task{
	{ID: "1245ID", Name: "groceries", Done: false},
	{ID: "6789ID", Name: "create additional endpoints", Done: false},
}

func CreateTask(c *echo.Context) error {
	var t Task
	if err := c.Bind(&t); err != nil {
		return BadRequest("invalid request body", err)
	}
	if t.Name == "" {
		return BadRequest("name of task is required", nil)
	}
	if t.ID != "" {
		return BadRequest("task ID is NOT required", nil)
	}
	t.ID = strconv.Itoa(len(sampleTasks) + rand.Int())
	return c.JSON(http.StatusCreated, t)
}

func GetAllTasks(c *echo.Context) error {
	return c.JSON(http.StatusOK, sampleTasks)
}
