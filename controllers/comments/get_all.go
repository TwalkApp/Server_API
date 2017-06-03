package comments

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/comments"
	"strconv"
	"github.com/twalkapp/server/models/misc"
)

func GetAllComments(depth int, limitDepth int, pagination misc.Pagination) ([]comments.Comment, error) {
	var (
		comment	comments.Comment
		result	[]comments.Comment
	)
	result = make([]comments.Comment, 0)

	query := "SELECT id, user_id, post_id, COALESCE(parent_id, 0), depth, value FROM comments "

	if depth != -1 {
		query += "WHERE depth = " + strconv.Itoa(depth)
	} else if limitDepth != -1 {
		query += "WHERE depth BETWEEN 0 AND " + strconv.Itoa(limitDepth)
	}
	if pagination.IsSet() {
		query += "LIMIT " + strconv.Itoa(pagination.GetFrom()) + "," + strconv.Itoa(pagination.Size)
	}

	rows, err := mysql.DB.Query(query + ";")
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
