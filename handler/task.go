package handler

import (
	"net/http"
	"strconv"

	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/service"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=./task.go -destination=./task_mock.go -package=handler
type TaskHandler interface {
	Create(c echo.Context) error
	GetTasks(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
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

func (h *taskHandler) GetTasks(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	tasks, err := h.Service.GetTasksByID(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *taskHandler) Update(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string {"error": "erro ao ler os dados da requisição"})
	}

	err = h.Service.UpdateTask(userID, taskID, &task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *taskHandler) Delete(c echo.Context) error {
 	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id inválido"})
	}

	err = h.Service.DeleteTask(userID, taskID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}	

	return c.NoContent(http.StatusNoContent)
}	
