package services

import (
	"context"
	"github.com/minhtuhcmus/manabie-todo-server/graph/model"
)

type TodoService interface {
	CreateTodo(ctx context.Context, newTodo *model.NewTodo) (*model.Todo, error)
}
