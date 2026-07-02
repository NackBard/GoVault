package store

import "github.com/NackBard/GoVault/internal/model"

type NoteStore interface {
	Create(note model.Note) (model.Note, error)
	GetByID(id int) (model.Note, error)
	List() ([]model.Note, error)
	Update(note model.Note) (model.Note, error)
	Delete(id int) error
}

type TaskStore interface {
	Create(task model.Task) (model.Task, error)
	GetByID(id int) (model.Task, error)
	List(onlyPending bool) ([]model.Task, error)
	Update(task model.Task) (model.Task, error)
	Delete(id int) error
}
