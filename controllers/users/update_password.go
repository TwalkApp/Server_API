package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/users"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(id string, password users.Password) (bool, error) {
	stmt, err := mysql.DB.Prepare("UPDATE users SET password = ? WHERE id = ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password.New), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(hash, id)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
