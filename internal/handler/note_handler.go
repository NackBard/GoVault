package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/NackBard/GoVault/internal/model"
	"github.com/NackBard/GoVault/internal/store"
)

type NoteHandler struct {
	store store.NoteStore
}

func NewNoteHandler(s store.NoteStore) *NoteHandler {
	return &NoteHandler{store: s}
}

func (h NoteHandler) HandleList(w http.ResponseWriter, r *http.Request) {
	list, err := h.store.List()

	if err != nil {
		respondError(w, http.StatusInternalServerError, store.ErrInternalServer.Error())
		return
	}

	respondJSON(w, http.StatusOK, list)
}

func (h NoteHandler) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		respondError(w, http.StatusBadRequest, store.ErrInvalidRequestBody.Error())
		return
	}

	createdNote, err := h.store.Create(note)

	if err != nil {
		respondError(w, http.StatusInternalServerError, store.ErrInternalServer.Error())
		return
	}

	respondJSON(w, http.StatusCreated, createdNote)
}

func (h NoteHandler) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("id")
	id, err := strconv.Atoi(value)

	if err != nil {
		respondError(w, http.StatusBadRequest, store.ErrInvalidId.Error())
		return
	}

	note, err := h.store.GetByID(id)

	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			respondError(w, http.StatusNotFound, store.ErrNotFound.Error())
			return
		}
		respondError(w, http.StatusInternalServerError, store.ErrInternalServer.Error())
		return
	}

	respondJSON(w, http.StatusOK, note)
}

func (h NoteHandler) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		respondError(w, http.StatusBadRequest, store.ErrInvalidRequestBody.Error())
		return
	}

	value := r.PathValue("id")
	id, err := strconv.Atoi(value)
	if err != nil {
		respondError(w, http.StatusBadRequest, store.ErrInvalidId.Error())
		return
	}

	note.ID = int64(id)
	updateNote, err := h.store.Update(note)

	if err != nil {
		respondError(w, http.StatusInternalServerError, store.ErrInternalServer.Error())
		return
	}

	respondJSON(w, http.StatusOK, updateNote)
}

func (h NoteHandler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("id")
	id, err := strconv.Atoi(value)

	if err != nil {
		respondError(w, http.StatusBadRequest, store.ErrInvalidId.Error())
		return
	}

	err = h.store.Delete(id)

	if err != nil {
		respondError(w, http.StatusInternalServerError, store.ErrInternalServer.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
