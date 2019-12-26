package api

// User represents a User instance in the DB
type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
