package api

import "database/sql"

type Category struct {
	Id          int           `json:"id,omitempty"`
	Subject     int           `json:"subject,omitempty"`
	ParentId    sql.NullInt64 `json:"parent_id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
}
