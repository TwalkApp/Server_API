package posts

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/posts"
	"github.com/twalkapp/server/models/misc"
	"strconv"
)

func GetAllPosts(pagination misc.Pagination) ([]posts.Post, error) {
	var (
		post  	posts.Post
		result	[]posts.Post
	)
	result = make([]posts.Post, 0)
	query := "SELECT id, user_id, title, description FROM posts "
	if pagination.IsSet() {
		query += "LIMIT " + strconv.Itoa(pagination.GetFrom()) + "," + strconv.Itoa(pagination.Size)
	}
	rows, err := mysql.DB.Query(query + ";")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.User, &post.Title, &post.Desc)
		result = append(result, post)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return result, err
}
