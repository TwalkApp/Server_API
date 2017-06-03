package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/posts"
	Misc "github.com/twalkapp/server/misc/pagination"
	"github.com/twalkapp/server/misc/log"
	"github.com/twalkapp/server/models/posts"
)

func SetPostsRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		log.Info("GET /posts")
		pagination, err := Misc.GetPagination(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with pagination",
			})
			return
		}
		result, err := Controller.GetAllPosts(pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if pagination.IsSet() {
			count, err := Controller.GetPostsCount()
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
			c.JSON(http.StatusOK, result)
		}
	})

	routerGroup.POST("", func(c *gin.Context) {
		log.Info("POST /posts")
		var post posts.Post
		if c.BindJSON(&post) == nil {
			_, err := Controller.CreatePost(post)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while creating post",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Post successfully created",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

}
