package api

type Credentials struct {
	Name string `json:"name",db:"name"`
	Password string `json:"password",db:"password"`
}
