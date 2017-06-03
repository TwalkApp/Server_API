package comments

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/comments"
	"strconv"
)

func CreateComment(comment comments.Comment) (bool, error) {
	if comment.Parent != 0 {
		parent, err := GetComment(strconv.Itoa(comment.Parent))
		if err != nil {
			fmt.Print(err.Error())
			return false, err
		}
		comment.Depth = parent.Depth + 1
	}
	if comment.Depth > 3 {
		comment.Depth = 3
	}
	stmt, err := mysql.DB.Prepare("INSERT INTO comments (user_id, post_id, parent_id, depth, value) SELECT ?, ?, NULLIF(?, 0), ?, ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(comment.User, comment.Post, comment.Parent, comment.Depth, comment.Text)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
