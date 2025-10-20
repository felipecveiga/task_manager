package handler

import (
	"net/http"
	"strconv"

	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/service"
	"github.com/labstack/echo/v4"
)

type TaskHandler interface {
	Create(c echo.Context) error
}

type taskHandler struct {
	Service service.TaskService
}

func NewTaskHandler(s service.TaskService) TaskHandler {
	return &taskHandler{
		Service: s,
	}
}

func (h *taskHandler) Create(c echo.Context) error {
	task := new(model.Task)
	if err := c.Bind(task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	err = h.Service.CreateTask(userID, task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, task)
}
