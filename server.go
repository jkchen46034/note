package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"jk.com/note/handlers"
	"jk.com/note/models"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/notedb?sslmode=disable&binary_parameters=yes")
	if err != nil {
		log.Fatalln(err)
	}
	noteModel := model.NoteModel{db}
	noteHandler := &handler.NoteHandler{&noteModel}

	mux := http.NewServeMux()
	mux.Handle("/notes", noteHandler)
	mux.Handle("/notes/", noteHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
