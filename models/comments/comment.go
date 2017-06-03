package comments

type Comment struct {
	Id	int	`json:"id"`
	User  	int	`json:"user_id"`
	Post  	int	`json:"post_id"`
	Parent  int	`json:"parent_id"`
	Depth	int	`json:"depth"`
	Text	string	`json:"text"`
}
