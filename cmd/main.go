package main

import (
	"log"

	todo "github.com/gimaevra94/todo"
	handler "github.com/gimaevra94/todo/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
