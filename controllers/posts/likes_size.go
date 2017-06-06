package posts

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
)

func GetPostLikesCount(id string) (int, error) {
	var result	int
	row := mysql.DB.QueryRow("SELECT COUNT(user_id) FROM favourites WHERE post_id = ?;", id)
	err := row.Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
	}
	return result, err
}
