package users

type Profile struct {
	Id        int		`json:"id"`
	Username  string        `json:"username"`
	Mail      string        `json:"mail"`
	Firstname string        `json:"firstname"`
	Lastname  string        `json:"lastname"`
}
