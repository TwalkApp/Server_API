package posts

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/posts"
)

func GetPost(id string) (posts.Post, error) {
	var post	posts.Post
	row := mysql.DB.QueryRow("SELECT id, user_id, title, description FROM posts WHERE id = ?;", id)
	err := row.Scan(&post.Id, &post.User, &post.Title, &post.Desc)
	if err != nil {
		fmt.Print(err.Error())
	}
	return post, err
}
