package main

import (
	"github.com/gilang-sas/todo-app/db"
	"github.com/gilang-sas/todo-app/server"
)

func main() {
	db.Init()
	server.Init()
}