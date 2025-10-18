package main

import (
	"net/http"

	"github.com/felipecveiga/task_manager/db"
	"github.com/labstack/echo/v4"
)

func main() {

	db := db.Connection()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
