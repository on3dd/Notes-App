package api

// Subject represents a Subject instance in the DB
type Subject struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
