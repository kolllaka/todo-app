package main

import (
	"log"

	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/KoLLlaka/todo-app/pkg/handler"
	"github.com/KoLLlaka/todo-app/pkg/repository"
	"github.com/KoLLlaka/todo-app/pkg/service"
)

const (
	port = "8080"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}
}
