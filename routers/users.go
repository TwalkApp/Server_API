package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/users"
	Misc "github.com/twalkapp/server/misc/pagination"
	"github.com/twalkapp/server/models/users"
	"github.com/twalkapp/server/misc/log"
)

func SetUsersRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		log.Info("GET /users")
		pagination, err := Misc.GetPagination(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with pagination",
			})
			return
		}
		result, err := Controller.GetAllUsers(pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if pagination.IsSet() {
			count, err := Controller.GetUsersCount()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			pagination.SetInformations(count)
			c.JSON(http.StatusOK, gin.H{
				"pagination": pagination,
				"users":      result,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	routerGroup.POST("", func(c *gin.Context) {
		log.Info("POST /users")
		var user users.User
		if c.BindJSON(&user) == nil {
			_, err := Controller.CreateUser(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while creating user",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf(" %s successfully created", user.Username),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

}
