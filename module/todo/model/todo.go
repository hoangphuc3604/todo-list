package model

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	TableNameTodo = "todo"
)

type TodoItem struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}


type TodoItemCreation struct {
	Title       string `json:"title" gorm:"column:title" validate:"required,min=3,max=100"`
	Description string `json:"description" gorm:"column:description" validate:"required,min=3,max=100"`
}


func (t *TodoItemCreation) Validate() error {
	return validate(t)
}

type TodoItemUpdate struct {
	Completed *bool `json:"completed" gorm:"column:completed" validate:"required"`
}

func (t *TodoItemUpdate) Validate() error {
	return validate(t)
}

func GetValidationMessage(field, tag string) string {
	messages := map[string]map[string]string{
		"Title": {
			"required": "Title can not be empty",
			"min":      "Title must be at least 3 characters",
			"max":      "Title must be at most 100 characters",
		},
		"Description": {
			"required": "Description can not be empty",
			"min":      "Description must be at least 3 characters",
			"max": "Description must be at most 100 characters",
		},
	}

	if msg, exists := messages[field][tag]; exists {
		return msg
	}
	return "Invalid field"
}

func validate(t interface{}) error {
	validate := validator.New()
	err := validate.Struct(t)
	if err == nil {
		return nil
	}

	var messages []string
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			messages = append(messages, GetValidationMessage(e.Field(), e.Tag()))
		}
	}

	return errors.New(strings.Join(messages, ", "))
}
