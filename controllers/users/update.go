package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/users"
)

func UpdateUser(user users.Profile) (bool, error) {
	stmt, err := mysql.DB.Prepare("UPDATE users SET mail = ?, firstname = ?, lastname = ? WHERE id = ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(user.Mail, user.Firstname, user.Lastname, user.Id)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
