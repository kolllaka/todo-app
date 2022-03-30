package main

import (
	"log"

	"github.com/KoLLlaka/todo-app/internal/todo"
)

const (
	port = "8080"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run(port); err != nil {
		log.Fatalf("error occured while running http server: %s\n", err.Error())
	}
}
