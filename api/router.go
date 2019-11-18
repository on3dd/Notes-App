package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func (api *API) NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/notes/{id}", logHandlerCall(api.GetNote)).Methods("GET")
	r.HandleFunc("/api/v1/notes", logHandlerCall(api.GetNotes)).Methods("GET")
	r.HandleFunc("/api/v1/notes", logHandlerCall(api.AddNote)).Methods("POST")
	r.HandleFunc("/api/v1/notes/{id}", logHandlerCall(api.UpdateNote)).Methods("POST")
	r.HandleFunc("/api/v1/notes/{id}", logHandlerCall(api.DeleteNote)).Methods("DELETE")

	r.HandleFunc("/api/v1/users/{id}", logHandlerCall(api.GetUser)).Methods("GET")
	r.HandleFunc("/api/v1/users", logHandlerCall(api.GetUsers)).Methods("GET")
	r.HandleFunc("/api/v1/users", logHandlerCall(api.AddUser)).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", logHandlerCall(api.UpdateUser)).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", logHandlerCall(api.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/api/v1/categories/{id}", logHandlerCall(api.GetCategory)).Methods("GET")
	r.HandleFunc("/api/v1/categories", logHandlerCall(api.GetCategories)).Methods("GET")
	r.HandleFunc("/api/v1/categories", logHandlerCall(api.AddCategory)).Methods("POST")
	r.HandleFunc("/api/v1/categories/{id}", logHandlerCall(api.UpdateCategory)).Methods("PUT")
	r.HandleFunc("/api/v1/categories/{id}", logHandlerCall(api.DeleteCategory)).Methods("DELETE")

	r.HandleFunc("/api/v1/subjects/{id}", logHandlerCall(api.GetSubject)).Methods("GET")
	r.HandleFunc("/api/v1/subjects", logHandlerCall(api.GetSubjects)).Methods("GET")
	r.HandleFunc("/api/v1/subjects", logHandlerCall(api.AddSubject)).Methods("POST")
	r.HandleFunc("/api/v1/subjects/{id}", logHandlerCall(api.UpdateSubject)).Methods("PUT")
	r.HandleFunc("/api/v1/subjects/{id}", logHandlerCall(api.DeleteSubject)).Methods("DELETE")

	r.HandleFunc("/api/v1/teachers/{id}", logHandlerCall(api.GetTeacher)).Methods("GET")
	r.HandleFunc("/api/v1/teachers", logHandlerCall(api.GetTeachers)).Methods("GET")
	r.HandleFunc("/api/v1/teachers", logHandlerCall(api.AddTeacher)).Methods("POST")
	r.HandleFunc("/api/v1/teachers/{id}", logHandlerCall(api.UpdateTeacher)).Methods("PUT")
	r.HandleFunc("/api/v1/teachers/{id}", logHandlerCall(api.DeleteTeacher)).Methods("DELETE")

	return r
}

// logHandlerCall logs any handler call
func logHandlerCall(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()
		log.Printf("Handler function called: %v", name)
		handler(w, r)
	}
}
