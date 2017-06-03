package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/users"
	"github.com/twalkapp/server/models/auth"
	"github.com/twalkapp/server/misc/log"
)

func SetAuthRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.POST("/auth", func(c *gin.Context) {
		log.Info("POST /auth")
		var login auth.Login
		if c.BindJSON(&login) == nil {
			result, status, err := Controller.AuthUser(login)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while auth",
				})
				return
			}
			if status {
				c.JSON(http.StatusOK, result)
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Wrong password",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

}
