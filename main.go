package main

import (
	"fmt"
	"net/http"

	"github.com/felipecveiga/task_manager/db"
	"github.com/labstack/echo/v4"
)

func main() {

	clientDB := db.Connection()

	//repository := repository.NewUserRepository(clientDB)
	//service := service.NewUserService(repository)
	//handler := handler.NewUserHandler(service)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	fmt.Println(clientDB)
	e.Logger.Fatal(e.Start(":8080"))
}
