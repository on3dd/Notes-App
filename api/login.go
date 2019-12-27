package api

import (
	"database/sql"
	"log"
	"net/http"
)

func (api *API) SingIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var user User
	user.Name = r.FormValue("name")
	user.Password = r.FormValue("password")
	
	var password string
	result := api.db.QueryRow("SELECT password FROM users WHERE name = $1", user.Name)
	err := result.Scan(&password)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
	}

	if ok := CheckPasswordHash(password, user.Password); !ok {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error checking passwords")
	}

	if err := setSession(user.Name, w); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error setting session: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}
