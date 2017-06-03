package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
)

func DeleteUser(id string) (bool, error) {
	stmt, err := mysql.DB.Prepare("DELETE FROM users WHERE id = ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
