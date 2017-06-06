package users

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/posts"
	"github.com/twalkapp/server/models/misc"
)

func GetUserPosts(id string, pagination misc.Pagination) ([]posts.Post, error) {
	var (
		post  	posts.Post
		result	[]posts.Post
	)
	result = make([]posts.Post, 0)
	query := "SELECT id, title, description FROM posts WHERE user_id = ?"
	if pagination.IsSet() {
		query += " LIMIT " + strconv.Itoa(pagination.GetFrom()) + "," + strconv.Itoa(pagination.Size)
	}
	rows, err := mysql.DB.Query(query + ";", id)
	for rows.Next() {
		err = rows.Scan(&post.Id, &post.Title, &post.Desc)
		post.User, _ = strconv.Atoi(id)
		result = append(result, post)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return result, err
}
