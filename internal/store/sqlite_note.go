package store

import (
	"database/sql"
	"encoding/json"

	"github.com/NackBard/GoVault/internal/model"
)

type SQLiteNoteStore struct {
	db *sql.DB
}

func NewSQLiteNoteStore(db *sql.DB) *SQLiteNoteStore {
	return &SQLiteNoteStore{db}
}

func (s SQLiteNoteStore) Create(note model.Note) (model.Note, error) {
	tagsJSON, err := json.Marshal(note.Tags)
	if err != nil {
		return note, err
	}

	result, err := s.db.Exec("insert into notes (title, body, tags, created_at, updated_at) values (?, ?, ?, ?, ?)",
		note.Title, note.Body, tagsJSON, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		return note, err
	}

	index, err := result.LastInsertId()

	if err != nil {
		return note, err
	}

	note.ID = index

	return note, err
}
