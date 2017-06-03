package comments

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/comments"
)

func UpdateComment(comment comments.Comment) (bool, error) {
	stmt, err := mysql.DB.Prepare("UPDATE comments SET value = ? WHERE id = ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(comment.Id, comment.Text)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
