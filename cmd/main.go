package main

import (
	"log"

	todo "github.com/gimaevra94/todo/pkg"
	handler "github.com/gimaevra94/todo/pkg/handlers"
)

func main() {
	handlers := new(handler.Handler)
	server := new(todo.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
