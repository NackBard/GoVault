package store

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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

func (s SQLiteNoteStore) List() ([]model.Note, error) {
	rows, err := s.db.Query("SELECT * FROM notes")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var n model.Note
		var tagsStr string

		if err := rows.Scan(&n.ID, &n.Title, &n.Body, &tagsStr, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}

		if err := json.Unmarshal([]byte(tagsStr), &n.Tags); err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, err
}

func (s SQLiteNoteStore) GetByID(id int) (model.Note, error) {
	row := s.db.QueryRow("SELECT * FROM notes WHERE id = ?", id)
	var note model.Note
	var tagsStr string

	if err := row.Scan(&note.ID, &note.Title, &note.Body, &tagsStr, &note.CreatedAt, &note.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Note{}, fmt.Errorf("note %d not found", id)
		}
		return model.Note{}, fmt.Errorf("get note: %w", err)
	}

	err := json.Unmarshal([]byte(tagsStr), &note.Tags)

	return note, err
}

func (s SQLiteNoteStore) Update(note model.Note) (model.Note, error) {
	tagsJSON, err := json.Marshal(note.Tags)
	if err != nil {
		return note, err
	}

	note.UpdatedAt = time.Now().UTC()
	result, err := s.db.Exec(`update notes set
	title = ?, body = ?, tags = ?,
	updated_at = ? where id = ?`, note.Title, note.Body, string(tagsJSON), note.UpdatedAt, note.ID)

	if err != nil {
		return note, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return note, err
	}

	if rowsAffected == 0 {
		return note, fmt.Errorf("no note found with id %d", note.ID)
	}

	return note, nil
}
func (s SQLiteNoteStore) Delete(id int) error {
	result, err := s.db.Exec("delete from notes where id = ?", id)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf("no note found with id %d", id)
	}

	return nil
}
