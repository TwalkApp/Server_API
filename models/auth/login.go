package auth

type Login struct {
	Username  string        `json:"login"`
	Password  string        `json:"password"`
}
