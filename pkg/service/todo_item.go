package service

import (
	"github.com/KoLLlaka/todo-app/internal/todo"
	"github.com/KoLLlaka/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userID, listID int, input todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list does not exists or does not belong to user
		return 0, err
	}

	return s.repo.Create(listID, input)
}

func (s *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list does not exists or does not belong to user
		return []todo.TodoItem{}, err
	}

	return s.repo.GetAll(listID)
}

func (s *TodoItemService) GetByID(userID, listID, itemID int) (todo.TodoItem, error) {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list does not exists or does not belong to user
		return todo.TodoItem{}, err
	}

	return s.repo.GetByID(listID, itemID)
}

func (s *TodoItemService) Delete(userID, listID, itemID int) error {
	_, err := s.listRepo.GetByID(userID, listID)
	if err != nil {
		// list does not exists or does not belong to user
		return  err
	}

	return s.repo.Delete(listID, itemID)
}