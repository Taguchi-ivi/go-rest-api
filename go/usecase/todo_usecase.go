package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITodoUsecase interface {
	GetAllTodos(userId uint) ([]model.TodoResponse, error)
	GetTodoById(userId uint, todoId uint) (model.TodoResponse, error)
	CreateTodo(todo model.Todo) (model.TodoResponse, error)
	UpdateTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error)
	DeleteTodo(userId uint, todoId uint) error
}

type todoUsecase struct {
	tr repository.ITodoRepository
	tv validator.ITodoValidator
}

func NewTodoUsecase(tr repository.ITodoRepository, tv validator.ITodoValidator) ITodoUsecase {
	return &todoUsecase{tr, tv}
}

func (tu *todoUsecase) GetAllTodos(userId uint) ([]model.TodoResponse, error) {
	todos := []model.Todo{}
	if err := tu.tr.GetAllTodos(&todos, userId); err != nil {
		return nil, err
	}
	resTodos := []model.TodoResponse{}
	for _, v := range todos {
		t := model.TodoResponse{
			ID:        v.ID,
			Title:     v.Title,
			DeleteFlg: v.DeleteFlg,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTodos = append(resTodos, t)
	}
	return resTodos, nil
}

func (tu *todoUsecase) GetTodoById(userId uint, todoId uint) (model.TodoResponse, error) {
	todo := model.Todo{}
	if err := tu.tr.GetTodoById(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		DeleteFlg: todo.DeleteFlg,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) CreateTodo(todo model.Todo) (model.TodoResponse, error) {
	if err := tu.tv.TodoValidate(todo); err != nil {
		return model.TodoResponse{}, err
	}
	if err := tu.tr.CreateTodo(&todo); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		DeleteFlg: todo.DeleteFlg,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) UpdateTodo(todo model.Todo, userId uint, todoId uint) (model.TodoResponse, error) {
	if err := tu.tv.TodoValidate(todo); err != nil {
		return model.TodoResponse{}, err
	}
	if err := tu.tr.UpdateTodo(&todo, userId, todoId); err != nil {
		return model.TodoResponse{}, err
	}
	resTodo := model.TodoResponse{
		ID:        todo.ID,
		Title:     todo.Title,
		DeleteFlg: todo.DeleteFlg,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
	return resTodo, nil
}

func (tu *todoUsecase) DeleteTodo(userId uint, todoId uint) error {
	if err := tu.tr.DeleteTodo(userId, todoId); err != nil {
		return err
	}
	return nil
}
