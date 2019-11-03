package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Note represents a Note instance in the DB
type Note struct {
	Id          int            `json:"id"`
	Author      int            `json:"author_id"`
	CategoryId  int            `json:"category_id"`
	TeacherId   int            `json:"teacher_id"`
	PostedAt    string         `json:"posted_at"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"descirption"`
	Link        string         `json:"link"`
}

// GetNote gets single note from DB by id
func (api *API) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	row := api.db.QueryRow("SELECT * FROM notes WHERE id = $1", id)

	var note Note
	err := row.Scan(&note.Id, &note.Author, &note.CategoryId, &note.TeacherId, &note.PostedAt,
		&note.Title, &note.Description, &note.Link)
	if err == sql.ErrNoRows {
		WriteStatus(w, http.StatusNotFound, []byte("{'status':'error'}"))
		return
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		log.Fatal(err)
	}
}

// GetNotes gets all notes from DB
func (api *API) GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rows *sql.Rows
	var err error

	num := r.FormValue("num")
	if num == "" {
		rows, err = api.db.Query("SELECT * FROM notes ORDER BY posted_at DESC")
	} else {
		rows, err = api.db.Query("SELECT * FROM notes ORDER BY posted_at DESC LIMIT $1", num)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	var notes []*Note
	for rows.Next() {
		note := &Note{}
		err := rows.Scan(&note.Id, &note.Author, &note.CategoryId, &note.TeacherId, &note.PostedAt,
			&note.Title, &note.Description, &note.Link)
		if err != nil {
			WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
			log.Fatal(err)
		}
		notes = append(notes, note)
	}
	if err = rows.Err(); err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	if len(notes) == 0 {
		err = json.NewEncoder(w).Encode(make([]Note, 0))
	} else {
		err = json.NewEncoder(w).Encode(notes)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}
}

// AddNote adds a new note to DB
func (api *API) AddNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	var num int
	id := api.db.QueryRow("SELECT id FROM notes ORDER BY id DESC LIMIT 1")
	err = id.Scan(&num)
	if err == sql.ErrNoRows {
		note.Id = 1
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}
	note.Id = num + 1

	_, err = api.db.Exec("INSERT INTO notes VALUES($1, $2, $3, $4, $5, $6, $7, $8)",
		note.Id, note.Author, note.CategoryId, note.TeacherId, time.Now(), note.Title, note.Description, note.Link)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}

// UpdateNote updates a single note in DB by id
func (api *API) UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatal(err)
	}

	_, err = api.db.Exec("UPDATE notes SET (title, descirption, link) = ($2, $3, $4) WHERE id = ($1)",
		id, note.Title, note.Description, note.Link)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}

// DeleteNote deletes a single note from DB by id
func (api *API) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	_, err := api.db.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}
