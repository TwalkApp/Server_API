package config

import (
	"strconv"

	"github.com/jinzhu/configor"
)

var Conf = struct {
	Name	string	`default:"Twalk"`
	Port	int	`default:"80"`

	Database	struct {
		User		string	`default:"root"`
		Password	string	`default:"password"`
		Address		string	`default:"127.0.0.1"`
		Port		int	`default:"3306"`
		Table		string	`default:"Twalk"`
	}

	JWT	struct {
		Secret		string	`default:"secret"`
		Duration	int	`default:"1"`
	}

	Pagination	struct {
		Size	int	`default:"-1"`
	}
}{}

func init() {
	configor.Load(&Conf, "config.json")
}

func GetDatabaseSource() string {
	return Conf.Database.User + ":" + Conf.Database.Password + "@tcp(" + Conf.Database.Address + ":" + strconv.Itoa(Conf.Database.Port) + ")/" + Conf.Database.Table
}
