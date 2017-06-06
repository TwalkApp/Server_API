package comments

import (
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
)

func GetCommentsCount(depth int, limitDepth int) (int, error) {
	var result	int
	query := "SELECT COUNT(id) FROM comments"

	if depth != -1 {
		query += " WHERE depth = " + strconv.Itoa(depth)
	} else if limitDepth != -1 {
		query += " WHERE depth BETWEEN 0 AND " + strconv.Itoa(limitDepth)
	}

	row := mysql.DB.QueryRow(query + ";")
	err := row.Scan(&result)
	if err != nil {
		fmt.Print(err.Error())
	}
	return result, err
}
