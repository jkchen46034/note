package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"jk.com/note/data"
)

func main() {
	mux := http.NewServeMux()
	noteHandler := &NoteHandler{}
	mux.Handle("/notes", noteHandler)
	mux.Handle("/notes/", noteHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

var (
	createNoteRe = regexp.MustCompile(`^\/notes[\/]*$`)
	listNoteRe   = regexp.MustCompile(`^\/notes[\/]*$`)
	getNoteRe    = regexp.MustCompile(`^\/notes\/(\d+)$`)
	updateNoteRe = regexp.MustCompile(`^\/notes\/(\d+)$`)
	deleteNoteRe = regexp.MustCompile(`^\/notes\/(\d+)$`)
)

type NoteHandler struct{}

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
	note := data.Note{}
	notes, err := note.List()
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
	note := data.Note{}
	id, _ := strconv.Atoi(matches[1])
	err := note.Get(id)
	if err != nil {
		internalServerError(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
	return
}

func (n *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var note data.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		internalServerError(w, r)
		return
	}
	err := note.Create()
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
	var note data.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		internalServerError(w, r)
		return
	}
	id, _ := strconv.Atoi(matches[1])
	err := note.Update(id)
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
	note := data.Note{}
	id, _ := strconv.Atoi(matches[1])
	err := note.Delete(id)
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
