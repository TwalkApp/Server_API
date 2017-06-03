package posts

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/posts"
)

func CreatePost(post posts.Post) (bool, error) {
	stmt, err := mysql.DB.Prepare("INSERT INTO posts (user_id, title, description) SELECT ?, ?, ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(post.User, post.Title, post.Desc)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
