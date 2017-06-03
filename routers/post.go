package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/posts"
	"github.com/twalkapp/server/misc/log"
	"github.com/twalkapp/server/models/posts"
)

func SetPostRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		pid := c.Param("pid")
		log.Info("GET /posts/" + pid)
		result, err := Controller.GetPost(pid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		log.Info("PUT /posts/" + c.Param("pid"))
		var post posts.Post
		if c.BindJSON(&post) == nil {
			_, err := Controller.UpdatePost(post)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while updating post",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Post successfully updated",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

	routerGroup.DELETE("", func(c *gin.Context) {
		pid := c.Param("pid")
		log.Info("DELETE /posts/" + pid)
		_, err := Controller.DeletePost(pid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while deletion",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted post %s", pid),
		})

	})

}
