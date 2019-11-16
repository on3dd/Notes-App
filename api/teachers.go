package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

// Teacher represents a Teacher instance in the DB
type Teacher struct {
	Id       int    `json:"id"`
	SubjectId int `json:"subject_id"`
	Name     string `json:"name"`
}

// GetTeacher gets single teacher from DB by id
func (api *API) GetTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	row := api.db.QueryRow("SELECT * FROM teachers WHERE id = $1", id)

	var teacher Teacher
	err := row.Scan(&teacher.Id, &teacher.SubjectId, &teacher.Name)
	if err == sql.ErrNoRows {
		WriteStatus(w, http.StatusNotFound, []byte(`{"status":"error"}`))
		return
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(teacher)
	if err != nil {
		log.Fatal(err)
	}
}

// GetTeachers gets all teachers from DB
func (api *API) GetTeachers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var rows *sql.Rows
	var err error

	id := r.FormValue("id")
	if id == "" {
		rows, err = api.db.Query("SELECT * FROM teachers ORDER BY name")
	} else {
		rows, err = api.db.Query("SELECT * FROM teachers WHERE subject_id = $1 ORDER BY name", id)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	var teachers []*Teacher
	for rows.Next() {
		teacher := &Teacher{}
		err := rows.Scan(&teacher.Id, &teacher.SubjectId, &teacher.Name)
		if err != nil {
			WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
			log.Fatal(err)
		}
		teachers = append(teachers, teacher)
	}
	if err = rows.Err(); err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	if len(teachers) == 0 {
		err = json.NewEncoder(w).Encode(make([]Teacher, 0))
	} else {
		err = json.NewEncoder(w).Encode(teachers)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}
}

// TODO: Make this shit work
// AddTeacher adds a new teacher to DB
func (api *API) AddTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		log.Fatal(err)
	}

	//var num int
	//id := api.db.QueryRow("SELECT id FROM teacher ORDER BY id DESC LIMIT 1")
	//err = id.Scan(&num)
	//if err == sql.ErrNoRows {
	//	teacher.Id = 1
	//} else if err != nil {
	//	WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
	//	log.Fatal(err)
	//}
	//teacher.Id = num + 1

	_, err = api.db.Exec("INSERT INTO teachers VALUES($1, $2, $3)",
		&teacher.Id, &teacher.SubjectId, &teacher.Name)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}

// UpdateTeacher updates a single teacher in DB by id
func (api *API) UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	var teacher Teacher
	err := json.NewDecoder(r.Body).Decode(&teacher)
	if err != nil {
		log.Fatal(err)
	}

	_, err = api.db.Exec("UPDATE teachers SET name = $2 WHERE id = ($1)",
		id, teacher.Name)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}

// DeleteTeacher deletes a single teacher from DB by id
func (api *API) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	_, err := api.db.Exec("DELETE FROM teachers WHERE id = $1", id)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte(`{"status":"success"}`))
}