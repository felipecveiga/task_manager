package handler

import (
	"net/http"

	"github.com/felipecveiga/task_manager/model"
	"github.com/felipecveiga/task_manager/service"
	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=./user.go -destination=./user_mock.go -package=handler
type Handler interface {
	Create(c echo.Context) error
}

type handler struct {
	Service service.Service
}

func NewUserHandler(s service.Service) Handler {
	return &handler{
		Service: s,
	}
}

func (h *handler) Create(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.Service.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}
