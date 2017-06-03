package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"

	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/models/auth"
	"github.com/twalkapp/server/misc/jwt"
)

func AuthUser(login auth.Login) (gin.H, bool, error) {
	var (
		id string
		username string
		password []byte
	)
	row := mysql.DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?;", login.Username)
	err := row.Scan(&id, &username, &password)
	if err != nil {
		fmt.Print(err.Error())
		return nil, false, err
	}
	err = bcrypt.CompareHashAndPassword(password, []byte(login.Password))
	if err != nil {
		return nil, false, nil
	}
	profile, _ := GetUser(id)
	token := jwt.GenerateToken(profile)
	return gin.H{"token": token}, true, err
}
