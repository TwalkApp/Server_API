package users

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/users"
	"github.com/twalkapp/server/models/misc"
	"strconv"
)

func GetAllUsers(pagination misc.Pagination) ([]users.Profile, error) {
	var (
		user  	users.Profile
		result	[]users.Profile
	)
	result = make([]users.Profile, 0)
	query := "SELECT id, username, mail, firstname, lastname FROM users "
	if pagination.IsSet() {
		query += "LIMIT " + strconv.Itoa(pagination.GetFrom()) + "," + strconv.Itoa(pagination.Size)
	}
	rows, err := mysql.DB.Query(query + ";")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Mail, &user.Firstname, &user.Lastname)
		result = append(result, user)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	return result, err
}
