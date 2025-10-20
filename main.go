package main

import (
	"github.com/felipecveiga/task_manager/db"
	"github.com/felipecveiga/task_manager/handler"
	"github.com/felipecveiga/task_manager/repository"
	"github.com/felipecveiga/task_manager/service"
	"github.com/labstack/echo/v4"
)

func main() {

	clientDB := db.Connection()

	//User
	userRepo := repository.NewUserRepository(clientDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Task
	taskRepo := repository.NewTaskRepository(clientDB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	e := echo.New()

	// Rotas User
	e.POST("/user", userHandler.Create)

	// Rotas Task
	e.POST("/user/:id/tasks", taskHandler.Create)
	e.GET("/user/:id/tasks", taskHandler.GetTasks)
	e.DELETE("/user/:id/tasks/:task_id", taskHandler.Delete)

	e.Logger.Fatal(e.Start(":8080"))
}
