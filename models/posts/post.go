package posts

type Post struct {
	Id	int	`json:"id"`
	User  	int	`json:"user_id"`
	Title	string	`json:"title"`
	Desc	string	`json:"desc"`
}
