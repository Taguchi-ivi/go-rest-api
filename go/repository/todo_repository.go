package repository

import (
	"fmt"
	"go-rest-api/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITodoRepository interface {
	GetAllTodos(todos *[]model.Todo, userId uint) error
	GetTodoById(todo *model.Todo, userId uint, todoId uint) error
	CreateTodo(todo *model.Todo) error
	UpdateTodo(todo *model.Todo, userId uint, todoId uint) error
	DeleteTodo(userId uint, todoId uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) ITodoRepository {
	return &todoRepository{db}
}

func (tr *todoRepository) GetAllTodos(todos *[]model.Todo, userId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).Order("created_at").Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) GetTodoById(todo *model.Todo, userId uint, todoId uint) error {
	if err := tr.db.Joins("User").Where("user_id = ?", userId).
		First(todo, todoId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) CreateTodo(todo *model.Todo) error {
	if err := tr.db.Create(todo).Error; err != nil {
		return err
	}
	return nil
}

func (tr *todoRepository) UpdateTodo(todo *model.Todo, userId uint, todoId uint) error {
	result := tr.db.Model(todo).Clauses(clause.Returning{}).
		Where("id=? AND user_id=?", todoId, userId).
		// Updates(&model.Todo{
		Updates(map[string]interface{}{
			"title":      todo.Title,
			"delete_flg": todo.DeleteFlg,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (tr *todoRepository) DeleteTodo(userId uint, todoId uint) error {
	result := tr.db.Where("id=? AND user_id=?", todoId, userId).Delete(&model.Todo{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
