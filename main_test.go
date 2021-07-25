package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	handler "jk.com/note/handlers"
	model "jk.com/note/models"
)

type mockNoteModel struct{}

func (m *mockNoteModel) Create(note_in model.Note) (note model.Note, err error) {
	note = model.Note{
		ID:      38239,
		Content: "This is PG Unlimited Created!",
		Author:  "John Doe Dr.",
	}
	return note, nil
}

func (m *mockNoteModel) Get(id int) (note model.Note, err error) {
	note = model.Note{
		ID:      38239,
		Content: "This is PG Unlimited Get!",
		Author:  "John Doe Jr.",
	}
	return note, nil
}

func (m *mockNoteModel) List() (notes []model.Note, err error) {
	note := model.Note{
		ID:      38239,
		Content: "This is PG Unlimited List!",
		Author:  "John Doe",
	}
	notes = append(notes, note)
	return notes, nil
}

func (m *mockNoteModel) Update(note_in model.Note) (note model.Note, err error) {
	return
}

func (m *mockNoteModel) Delete(id int) (err error) {
	return
}

func TestNoteList(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/notes", nil)
	mock := mockNoteModel{}
	noteHandler := &handler.NoteHandler{&mock}
	http.HandlerFunc(noteHandler.List).ServeHTTP(rec, req)
	expected := model.Note{
		ID:      38239,
		Content: "This is PG Unlimited List!",
		Author:  "John Doe",
	}
	var got []model.Note
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	if len(got) != 1 {
		t.Errorf("Should be 1 but not")
	}
	if got[0].ID != expected.ID {
		t.Errorf("id not the same")
	}
	if got[0].Content != expected.Content {
		t.Errorf("content not the same")
	}
	if got[0].Author != expected.Author {
		t.Errorf("author not the same")
	}
}

func TestNoteGet(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/notes/38239", nil)
	mock := mockNoteModel{}
	noteHandler := &handler.NoteHandler{&mock}
	http.HandlerFunc(noteHandler.Get).ServeHTTP(rec, req)
	expected := model.Note{
		ID:      38239,
		Content: "This is PG Unlimited Get!",
		Author:  "John Doe Jr.",
	}
	var got model.Note
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	if got.ID != expected.ID {
		t.Errorf("id not the same")
	}
	if got.Content != expected.Content {
		t.Errorf("content not the same")
	}
	if got.Author != expected.Author {
		t.Errorf("author not the same")
	}
}

func TestNoteCreate(t *testing.T) {
	rec := httptest.NewRecorder()
	var jsonData = []byte(`{
		"content": "This is PG Unlimited Created!",
		"author": "John Doe Dr."
	}`)
	req, _ := http.NewRequest("POST", "/notes", bytes.NewBuffer(jsonData))
	mock := mockNoteModel{}
	noteHandler := &handler.NoteHandler{&mock}
	http.HandlerFunc(noteHandler.Create).ServeHTTP(rec, req)
	expected := model.Note{
		Content: "This is PG Unlimited Created!",
		Author:  "John Doe Dr.",
	}
	var got model.Note
	_ = json.Unmarshal(rec.Body.Bytes(), &got)
	if got.Content != expected.Content {
		t.Errorf("content not the same")
	}
	if got.Author != expected.Author {
		t.Errorf("author not the same")
	}
}
