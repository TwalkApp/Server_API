package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/users"
)

func GetUser(id string) (users.Profile, error) {
	var user	users.Profile
	row := mysql.DB.QueryRow("SELECT id, username, mail, firstname, lastname FROM users WHERE id = ?;", id)
	err := row.Scan(&user.Id, &user.Username, &user.Mail, &user.Firstname, &user.Lastname)
	if err != nil {
		fmt.Print(err.Error())
	}
	return user, err
}
