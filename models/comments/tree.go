package comments

type Tree struct {
	*Comment
	Childs	[]Comment	`json:"chlids"`
}
