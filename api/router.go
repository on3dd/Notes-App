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

	r.HandleFunc("/api/v1/getNote", logHandlerCall(api.GetNote)).Methods("GET")
	r.HandleFunc("/api/v1/getNotes", logHandlerCall(api.GetNotes)).Methods("GET")
	r.HandleFunc("/api/v1/addNote", logHandlerCall(api.AddNote)).Methods("POST")
	r.HandleFunc("/api/v1/updateNote", logHandlerCall(api.UpdateNote)).Methods("PUT")
	r.HandleFunc("/api/v1/deleteNote", logHandlerCall(api.DeleteNote)).Methods("DELETE")

	r.HandleFunc("/api/v1/getUser", logHandlerCall(api.GetUser)).Methods("GET")
	r.HandleFunc("/api/v1/getUsers", logHandlerCall(api.GetUsers)).Methods("GET")
	r.HandleFunc("/api/v1/addUser", logHandlerCall(api.AddUser)).Methods("POST")
	r.HandleFunc("/api/v1/updateUser", logHandlerCall(api.UpdateUser)).Methods("PUT")
	r.HandleFunc("/api/v1/deleteUser", logHandlerCall(api.DeleteUser)).Methods("DELETE")

	r.HandleFunc("/api/v1/getCategory", logHandlerCall(api.GetCategory)).Methods("GET")
	r.HandleFunc("/api/v1/getCategories", logHandlerCall(api.GetCategories)).Methods("GET")
	r.HandleFunc("/api/v1/addCategory", logHandlerCall(api.AddCategory)).Methods("POST")
	r.HandleFunc("/api/v1/updateCategory", logHandlerCall(api.UpdateCategory)).Methods("PUT")
	r.HandleFunc("/api/v1/deleteCategory", logHandlerCall(api.DeleteCategory)).Methods("DELETE")

	r.HandleFunc("/api/v1/getSubject", logHandlerCall(api.GetSubject)).Methods("GET")
	r.HandleFunc("/api/v1/getSubjects", logHandlerCall(api.GetSubjects)).Methods("GET")
	r.HandleFunc("/api/v1/addSubject", logHandlerCall(api.AddSubject)).Methods("POST")
	r.HandleFunc("/api/v1/updateSubject", logHandlerCall(api.UpdateSubject)).Methods("PUT")
	r.HandleFunc("/api/v1/deleteSubject", logHandlerCall(api.DeleteSubject)).Methods("DELETE")

	r.HandleFunc("/api/v1/getTeacher", logHandlerCall(api.GetTeacher)).Methods("GET")
	r.HandleFunc("/api/v1/getTeachers", logHandlerCall(api.GetTeachers)).Methods("GET")
	r.HandleFunc("/api/v1/addTeacher", logHandlerCall(api.AddTeacher)).Methods("POST")
	r.HandleFunc("/api/v1/updateTeacher", logHandlerCall(api.UpdateTeacher)).Methods("PUT")
	r.HandleFunc("/api/v1/deleteTeacher", logHandlerCall(api.DeleteTeacher)).Methods("DELETE")

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