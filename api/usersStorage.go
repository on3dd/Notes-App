package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// GetUser gets single user from DB by id
func (api *API) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	row := api.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user User
	err := row.Scan(&user.Id, &user.Name)
	if err == sql.ErrNoRows {
		WriteStatus(w, http.StatusNotFound, []byte(`{"status":"error"}`))
		return
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}

// GetUsers gets all users from DB
func (api *API) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := api.db.Query("SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.About, &user.Password)
		if err != nil {
			WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
			log.Fatal(err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}

	if len(users) == 0 {
		err = json.NewEncoder(w).Encode(make([]User, 0))
	} else {
		err = json.NewEncoder(w).Encode(users)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}
}

// UpdateUser updates a single user in DB by id
func (api *API) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}

	_, err = api.db.Exec("UPDATE users SET (name, password) = ($2, $3) WHERE id = ($1)",
		id, user.Name, user.Password)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}

// DeleteUser deletes a single user from DB by id
func (api *API) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	_, err := api.db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte(`{"status":"error"}`))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}
