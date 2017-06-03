package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/users"
)

func CreateUser(user users.User) (bool, error) {
	stmt, err := mysql.DB.Prepare("INSERT INTO users (username, mail, firstname, lastname, password) SELECT ?, ?, ?, ?, ?;")
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}
	_, err = stmt.Exec(user.Username, user.Mail, user.Firstname, user.Lastname, hash)
	if err != nil {
		defer stmt.Close()
		fmt.Print(err.Error())
		return false, err
	}
	return true, err
}
