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

func SetUserRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		uid := c.Param("uid")
		log.Info("GET /users/" + uid)
		result, err := Controller.GetUser(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		log.Info("PUT /users/" + c.Param("uid"))
		var user users.Profile
		if c.BindJSON(&user) == nil {
			_, err := Controller.UpdateUser(user)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while updating user",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf(" %s successfully updated", user.Username),
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

	routerGroup.DELETE("", func(c *gin.Context) {
		uid := c.Param("uid")
		log.Info("DELETE /users/" + uid)
		_, err := Controller.DeleteUser(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while deletion",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user %s", uid),
		})
	})

	routerGroup.PUT("/password", func(c *gin.Context) {
		uid := c.Param("uid")
		log.Info("PUT /users/" + uid + "/password")
		var password users.Password
		if c.BindJSON(&password) == nil {
			_, err := Controller.UpdatePassword(uid, password)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while updating user password",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Password successfully updated",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

	routerGroup.GET("/posts", func(c *gin.Context) {
		uid := c.Param("uid")
		log.Info("GET /users/" + uid + "/posts")
		pagination, err := Misc.GetPagination(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with pagination",
			})
			return
		}
		result, err := Controller.GetUserPosts(uid, pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if pagination.IsSet() {
			count, err := Controller.GetPostsCount(uid)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			pagination.SetInformations(count)
			c.JSON(http.StatusOK, gin.H{
				"pagination": pagination,
				"posts":      result,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

}
