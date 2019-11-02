package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Category struct {
	Id          int           `json:"id"`
	Subject     int           `json:"subject"`
	ParentId    sql.NullInt64 `json:"parent_id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
}

// GetCategory gets single category from DB by id
func (api *API) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	row := api.db.QueryRow("SELECT * FROM categories WHERE id = $1", id)

	var category Category
	err := row.Scan(&category.Id, &category.Subject, &category.ParentId, &category.Name, &category.Description)
	if err == sql.ErrNoRows {
		WriteStatus(w, http.StatusNotFound, []byte("{'status':'error'}"))
		return
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		log.Fatal(err)
	}
}

// GetCategories gets all categories from DB
func (api *API) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rows, err := api.db.Query("SELECT * FROM categories ORDER BY name")
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	var categories []*Category
	for rows.Next() {
		category := &Category{}
		err := rows.Scan(&category.Id, &category.Subject, &category.ParentId, &category.Name, &category.Description)
		if err != nil {
			WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
			log.Fatal(err)
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	if len(categories) == 0 {
		err = json.NewEncoder(w).Encode(make([]Category, 0))
	} else {
		err = json.NewEncoder(w).Encode(categories)
	}
	if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}
}

// AddCategory adds a new category to DB
func (api *API) AddCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatal(err)
	}

	var num int
	id := api.db.QueryRow("SELECT id FROM categories ORDER BY id DESC LIMIT 1")
	err = id.Scan(&num)
	if err == sql.ErrNoRows {
		category.Id = 1
	} else if err != nil {
		WriteStatus(w, http.StatusInternalServerError, []byte("{'status':'error'}"))
		log.Fatal(err)
	}
	category.Id = num + 1

	_, err = api.db.Exec("INSERT INTO categories VALUES($1, $2, $3, $4, $5)",
		category.Id, category.Subject, category.ParentId, category.Name, category.Description)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		log.Fatal(err)
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}

// UpdateCategory updates a single category in DB by id
func (api *API) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	var category Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Fatal(err)
	}

	_, err = api.db.Exec("UPDATE categories SET (name, description) = ($2, $3) WHERE id = ($1)",
		id, category.Name, category.Description)

	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}

// DeleteCategory deletes a single category from DB by id
func (api *API) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.FormValue("id")
	if id == "" {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	_, err := api.db.Exec("DELETE FROM categories WHERE id = $1", id)
	if err != nil {
		WriteStatus(w, http.StatusBadRequest, []byte("{'status':'error'}"))
		return
	}

	WriteStatus(w, http.StatusOK, []byte("{'status':'success'}"))
}
