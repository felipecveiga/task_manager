package main

import (
	"fmt"

	"github.com/felipecveiga/task_manager/db"
	"github.com/felipecveiga/task_manager/handler"
	"github.com/felipecveiga/task_manager/repository"
	"github.com/felipecveiga/task_manager/service"
	"github.com/labstack/echo/v4"
)

func main() {

	clientDB := db.Connection()

	repository := repository.NewUserRepository(clientDB)
	service := service.NewUserService(repository)
	handler := handler.NewUserHandler(service)

	e := echo.New()
	e.POST("/user", handler.Create)

	fmt.Println(clientDB)
	e.Logger.Fatal(e.Start(":8080"))
}
