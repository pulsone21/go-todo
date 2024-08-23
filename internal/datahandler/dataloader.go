package datahandler

import (
	"github.com/pulsone21/go-todo/internal/entities"
)

type Handler interface {
	Load() ([]entities.Todo, error)
	Save([]entities.Todo) error
}
