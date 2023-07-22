package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITodoValidator interface {
	TodoValidate(todo model.Todo) error
}

type todoValidator struct{}

func NewTodoValidator() ITodoValidator {
	return &todoValidator{}
}

func (tv *todoValidator) TodoValidate(todo model.Todo) error {
	return validation.ValidateStruct(&todo,
		validation.Field(
			&todo.Title,
			validation.Required.Error("title is required"),
		),
	)
}
