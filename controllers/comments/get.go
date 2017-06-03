package comments

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/comments"
)

func GetComment(id string) (comments.Tree, error) {
	var comment comments.Comment
	row := mysql.DB.QueryRow("SELECT id, user_id, post_id, COALESCE(parent_id, 0), depth, value FROM comments WHERE id = ?;", id)
	err := row.Scan(&comment.Id,
		&comment.User,
		&comment.Post,
		&comment.Parent,
		&comment.Depth,
		&comment.Text)
	if err != nil {
		fmt.Print(err.Error())
	}
	childs, err := GetCommentChilds(id)
	if err != nil {
		fmt.Print(err.Error())
		return comments.Tree{Comment: &comment, Childs: make([]comments.Comment, 0)}, err
	}
	return comments.Tree{Comment: &comment, Childs: childs}, err
}
