package main

import (
	"log"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/KoLLlaka/todo-app/pkg/handler"
)

const (
	port = "8080"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}
}
