package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/gilang-sas/todo-app/controller"
)

func NewRouter() *echo.Echo {
	r := echo.New()

	r.Use(
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions},
			AllowCredentials: true,
		}),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, latency=${latency}, status=${status}\n",
		}),
	)
		r.GET("/ping", controller.Ping)
		r.GET("/task", controller.GetAllTask)
		r.POST("/task", controller.AddTask)
		r.PUT("/task/:id", controller.TaskComplete)
		r.PUT("/undoTask/:id", controller.UndoTask)
		r.DELETE("/delete/:id", controller.DeleteTask)
		r.DELETE("/delete", controller.DeleteAllTask)

	return r
}

func Init() error {
	r := NewRouter()
	return r.Start(":8000")
}