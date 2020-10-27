package controller

import (

	"github.com/gilang-sas/todo-app/model"
	"github.com/labstack/echo"
)

func GetAllTask(c echo.Context) error {
	data, err := model.GetAllTask()
	if err != nil {
		return Unprocessable(c, err.Error())
	}

	return Data(c, data)

}

func AddTask(c echo.Context) error {
	var t model.ToDo
	if err := c.Bind(&t); err != nil {
		return Bad(c)
	}

	if err := model.InsertTask(&t); err != nil {
		return Unprocessable(c, err.Error()) 
	}

	return OK(c, "succedd add task ")
}

func TaskComplete(c echo.Context) error {
	id := c.Param("id")
	err := model.TaskComplete(id)
	if err != nil {
		return Unprocessable(c, err.Error())
	}

	return OK(c, "Task Complete..")

}

func UndoTask(c echo.Context) error {
	id := c.Param("id")
	err := model.UndoTask(id)
	if err != nil {
		return Unprocessable(c, err.Error())
	}

	return OK(c, "succeed undo task")

}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	err := model.DeleteTask(id)
	if err != nil {
		return Unprocessable(c, err.Error())
	}

	return OK(c, "succeed delete task")

}

func DeleteAllTask(c echo.Context) error {
	err := model.DeleteAllTask()
	if err != nil {
		return Unprocessable(c, err.Error())
	}

	return OK(c, "succeed delete all task")
}