package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

type API struct {
	db *sql.DB
	verifyKey []byte
	singKey []byte
}

// New returns a new API instance
func New(db *sql.DB, sk, vk []byte) *API {
	return &API{
		db: db,
		singKey: sk,
		verifyKey: vk,
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

// CreateDirIfNotExist creates new path if it doesn't exist
func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}