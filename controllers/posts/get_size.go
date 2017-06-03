package posts

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
)

func GetPostsCount() (int, error) {
	var result	int
	row := mysql.DB.QueryRow("SELECT COUNT(id) FROM posts;")
	err := row.Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
	}
	return result, err
}

