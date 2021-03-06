package service

import (
	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/KoLLlaka/todo-app/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userID int, list todo.TodoList) (int, error) {
	return s.repo.Create(userID, list)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetByID(userID, listID int) (todo.TodoList, error) {
	return s.repo.GetByID(userID, listID)
}

func (s *TodoListService) Update(userID, listID int, updateInput todo.UpdateListInput) error {
	if err := updateInput.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, listID, updateInput)
}

func (s *TodoListService) Delete(userID, listID int) error {
	return s.repo.Delete(userID, listID)
}
