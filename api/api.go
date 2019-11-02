package api

import (
	"database/sql"
	"net/http"
)

type API struct {
	db *sql.DB
}

// New returns a new API instance
func New(db *sql.DB) *API {
	return &API{
		db: db,
	}
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := api.NewRouter()
	router.ServeHTTP(w, r)
}

// WriteStatus writes status of the request in header and body of the response
func WriteStatus(w http.ResponseWriter, status int, text[]byte) {
	w.WriteHeader(status)
	w.Write(text)
}