package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/twalkapp/server/misc/config"
)

var DB	*sql.DB
var err	error

func init() {
	ConnectDatabase()
}

func ConnectDatabase() {
	DB, err = sql.Open("mysql", config.GetDatabaseSource())
}

func CheckDB() *sql.DB  {
	if pinger := DB.Ping() ; err != nil || pinger!=nil  {
		fmt.Println("FAILED TO CONNECT", "err" , err)
	}
	return DB
}
