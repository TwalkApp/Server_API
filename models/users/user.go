package users

type User struct {
	*Profile
	Password  string	`json:"password"`
}
