package main

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/twalkapp/server/routers"
	"github.com/twalkapp/server/storage/mysql"
	"github.com/twalkapp/server/misc/config"
)

var router  *gin.Engine;

func init() {
	mysql.CheckDB()
	router = gin.New()
	group:=router.Group("/")
	routers.InitRoutes(group)
}

func main() {
	router.Run(":" + strconv.Itoa(config.Conf.Port))
}
