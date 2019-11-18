package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Subject represents a Subject instance in the DB
type Subject struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetSubject gets single subject from DB by id
func (api *API) GetSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	row := api.db.QueryRow("SELECT * FROM subjects WHERE id = $1", id)

	var subject Subject
	err := row.Scan(&subject.Id, &subject.Name)
	if err == sql.ErrNoRows {
		WriteStatus(w, http.StatusNotFound, []byte(`{"status":"error"}`))
		return
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(subject)
	if err != nil {
		log.Fatal(err)
	}
}

// GetSubjects gets all subjects from DB
func (api *API) GetSubjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rows *sql.Rows
	var err error

	id := r.FormValue("id")
	if id == "" {
		rows, err = api.db.Query("SELECT * FROM subjects ORDER BY name")
	} else {
		rows, err = api.db.Query("SELECT * FROM subjects WHERE id = $1 ORDER BY name", id)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	var subjects []*Subject
	for rows.Next() {
		subject := &Subject{}
		err := rows.Scan(&subject.Id, &subject.Name)
		if err != nil {
			WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
			log.Fatal(err)
		}
		subjects = append(subjects, subject)
	}
	if err = rows.Err(); err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	if len(subjects) == 0 {
		err = json.NewEncoder(w).Encode(make([]Subject, 0))
	} else {
		err = json.NewEncoder(w).Encode(subjects)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}
}

// AddSubject adds a new subject to DB
func (api *API) AddSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var subject Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		log.Fatal(err)
	}

	var num int
	id := api.db.QueryRow("SELECT id FROM subjects ORDER BY id DESC LIMIT 1")
	err = id.Scan(&num)
	if err == sql.ErrNoRows {
		subject.Id = 1
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}
	subject.Id = num + 1

	_, err = api.db.Exec("INSERT INTO subjects VALUES($1, $2)",
		subject.Id, subject.Name)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}

// UpdateSubject updates a single subject in DB by id
func (api *API) UpdateSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	var subject Subject
	err := json.NewDecoder(r.Body).Decode(&subject)
	if err != nil {
		log.Fatal(err)
	}

	_, err = api.db.Exec("UPDATE subjects SET name = $2 WHERE id = ($1)",
		id, subject.Name)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}

// DeleteSubject deletes a single subject from DB by id
func (api *API) DeleteSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	_, err := api.db.Exec("DELETE FROM subjects WHERE id = $1", id)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}
