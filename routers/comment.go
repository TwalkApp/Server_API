package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/comments"
	"github.com/twalkapp/server/misc/log"
	"github.com/twalkapp/server/models/comments"
)

func SetCommentRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		id := c.Param("cid")
		log.Info("GET /posts/" + c.Param("pid") + "/comments/" + id)
		result, err := Controller.GetComment(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	routerGroup.PUT("", func(c *gin.Context) {
		log.Info("PUT /posts/" + c.Param("pid") + "/comments/" + c.Param("cid"))
		var comment comments.Comment
		if c.BindJSON(&comment) == nil {
			_, err := Controller.UpdateComment(comment)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while updating comment",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Comment successfully updated",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})

	})

	routerGroup.DELETE("", func(c *gin.Context) {
		cid := c.Param("cid")
		log.Info("PUT /posts/" + c.Param("pid") + "/comments/" + cid)
		_, err := Controller.DeleteComment(cid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error while deletion",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted comment %s", cid),
		})
	})

}
