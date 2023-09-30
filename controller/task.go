package controller

import (
	"go-crean-aarchitecture/model"
	"go-crean-aarchitecture/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(e echo.Context) error
	Delete(e echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu: tu}
}

func (tc *taskController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	task, err := tc.tu.GetTask(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, task)
}

func (tc *taskController) Create(c echo.Context) error {
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(400, err)
	}
	taskID, err := tc.tu.CreateTask(task.Title)
	task.ID = taskID
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, task)
}

func (tc *taskController) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}

	task := model.Task{ID: id}
	if err := c.Bind(&task); err != nil {
		return c.JSON(400, err)
	}
	err = tc.tu.UpdateTask(task.ID, task.Title)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, task)
}

func (tc *taskController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err)
	}
	err = tc.tu.DeleteTask(id)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "deleted")
}
