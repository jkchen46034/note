package model

import (
	"database/sql"
)

type NoteModel struct {
	DB *sql.DB
}

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (m *NoteModel) Create(note_in Note) (note Note, err error) {
	stmt := `insert into notes(content, author) values ($1, $2) returning id, content, author;`
	err = m.DB.QueryRow(stmt, note_in.Content, note_in.Author).Scan(&note.ID, &note.Content, &note.Author)
	return
}

func (m *NoteModel) Get(id int) (note Note, err error) {
	stmt := `select id, content, author from notes where id = $1`
	err = m.DB.QueryRow(stmt, id).Scan(&note.ID, &note.Content, &note.Author)
	return
}

func (m *NoteModel) List() (notes []Note, err error) {
	stmt := `select id, content, author from notes order by id asc;`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var n Note
		err = rows.Scan(&n.ID, &n.Content, &n.Author)
		if err != nil {
			return
		}
		notes = append(notes, n)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return
}

func (m *NoteModel) Update(note_in Note) (note Note, err error) {
	stmt := `update notes set content = $2, author = $3 where id = $1 returning id, content, author`
	err = m.DB.QueryRow(stmt, note_in.ID, note_in.Content, note_in.Author).Scan(&note.ID, &note.Content, &note.Author)
	return
}

func (m *NoteModel) Delete(id int) (err error) {
	stmt := `delete from notes where id = $1`
	_, err = m.DB.Query(stmt, id)
	return
}
