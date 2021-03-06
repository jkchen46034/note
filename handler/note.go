package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"jk.com/note/model"
)

var (
	createNoteRe = regexp.MustCompile(`^\/notes[\/]*$`)
	listNoteRe   = regexp.MustCompile(`^\/notes[\/]*$`)
	getNoteRe    = regexp.MustCompile(`^\/notes\/(\d+)$`)
	updateNoteRe = regexp.MustCompile(`^\/notes\/(\d+)$`)
	deleteNoteRe = regexp.MustCompile(`^\/notes\/(\d+)$`)
)

type NoteHandler struct {
	Model interface {
		Create(note_in model.Note) (note model.Note, err error)
		Get(id int) (note model.Note, err error)
		List() (notes []model.Note, err error)
		Update(note_in model.Note) (note model.Note, err error)
		Delete(id int) (err error)
	}
}

func (n *NoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodGet && listNoteRe.MatchString(r.URL.Path):
		n.List(w, r)
		return
	case r.Method == http.MethodGet && getNoteRe.MatchString(r.URL.Path):
		n.Get(w, r)
		return
	case r.Method == http.MethodPost && createNoteRe.MatchString(r.URL.Path):
		n.Create(w, r)
		return
	case r.Method == http.MethodPatch && updateNoteRe.MatchString(r.URL.Path):
		n.Update(w, r)
		return
	case r.Method == http.MethodDelete && deleteNoteRe.MatchString(r.URL.Path):
		n.Delete(w, r)
		return
	}
}

func (n *NoteHandler) List(w http.ResponseWriter, r *http.Request) {
	notes, err := n.Model.List()
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
	return
}

func (n *NoteHandler) Get(w http.ResponseWriter, r *http.Request) {
	matches := getNoteRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	id, _ := strconv.Atoi(matches[1])
	note, err := n.Model.Get(id)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
	return
}

func (n *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var note_in model.Note
	if err := json.NewDecoder(r.Body).Decode(&note_in); err != nil {
		internalServerError(w, r)
		return
	}
	note, err := n.Model.Create(note_in)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
	return
}

func (n *NoteHandler) Update(w http.ResponseWriter, r *http.Request) {
	matches := getNoteRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	var note_in model.Note
	if err := json.NewDecoder(r.Body).Decode(&note_in); err != nil {
		internalServerError(w, r)
		return
	}
	id, _ := strconv.Atoi(matches[1])
	note_in.ID = id

	note, err := n.Model.Update(note_in)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
	return
}

func (n *NoteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	matches := getNoteRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		notFound(w, r)
		return
	}
	id, _ := strconv.Atoi(matches[1])
	err := n.Model.Delete(id)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}
