package data

type Note struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func (note *Note) Create() (err error) {
	stmt := `insert into notes(content, author) values ($1, $2) returning id;`
	err = Db.QueryRow(stmt, note.Content, note.Author).Scan(&note.ID)
	return
}

func (note *Note) Get(id int) (err error) {
	stmt := `select id, content, author from notes where id = $1`
	err = Db.QueryRow(stmt, id).Scan(&note.ID, &note.Content, &note.Author)
	return
}

func (note *Note) List() (notes []Note, err error) {
	stmt := `select id, content, author from notes order by id asc;`
	rows, err := Db.Query(stmt)
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
	err = rows.Err()
	if err != nil {
		return notes, err
	}
	return
}

func (note *Note) Update(id int) (err error) {
	stmt := `update notes set content = $2, author = $3 where id = $1 returning id, content, author`
	err = Db.QueryRow(stmt, id, note.Content, note.Author).Scan(&note.ID, &note.Content, &note.Author)
	return
}

func (note *Note) Delete(id int) (err error) {
	stmt := `delete from notes where id = $1`
	_, err = Db.Query(stmt, id)
	return
}
