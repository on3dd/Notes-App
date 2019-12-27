package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

// AddUser adds a new user to DB
func (api *API) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user User
	user.Name = r.FormValue("name")
	user.Password = r.FormValue("password")
	user.Email = r.FormValue("email")
	user.About = r.FormValue("about")

	var num int
	id := api.db.QueryRow("SELECT id FROM users ORDER BY id DESC LIMIT 1")
	err := id.Scan(&num)
	if err == sql.ErrNoRows {
		user.Id = 1
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatal(err)
	}
	user.Id = num + 1

	user.Password, err = HashPassword(user.Password)
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte(`{"status":"error"}`))
		log.Fatalf("Cannot hash user's password, error: %v", err)
	}

	_, err = api.db.Exec("INSERT INTO users VALUES($1, $2, $3, $4, $5)",
		user.Id, user.Name, user.Password, user.Email, user.About)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatal(err)
	}

	text := fmt.Sprintf(`{"status":"success", "id": %v}`, user.Id)
	WriteStatus(w, http.StatusOK, []byte(text))
}