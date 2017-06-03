package comments

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/comments"
)

func GetCommentChilds(id string) ([]comments.Comment, error) {
	var (
		comment	comments.Comment
		result	[]comments.Comment
	)
	result = make([]comments.Comment, 0)
	rows, err := mysql.DB.Query("SELECT id, user_id, post_id, COALESCE(parent_id, 0), depth, value FROM comments WHERE parent_id = ?;", id)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&comment.Id, &comment.User, &comment.Post, &comment.Parent, &comment.Depth, &comment.Text)
		result = append(result, comment)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return result, err
}
