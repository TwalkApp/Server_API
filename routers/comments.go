package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	Controller "github.com/twalkapp/server/controllers/comments"
	Misc "github.com/twalkapp/server/misc/pagination"
	"github.com/twalkapp/server/models/comments"
	"github.com/twalkapp/server/misc/log"
	"strconv"
)

func SetCommentsRoutes(routerGroup *gin.RouterGroup ) {

	routerGroup.GET("", func(c *gin.Context) {
		pid := c.Param("pid")
		depth, err := strconv.Atoi(c.DefaultQuery("depth", "-1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid depth value",
			})
			return
		}
		limitDepth, err := strconv.Atoi(c.DefaultQuery("limit_depth", "-1"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid limit_depth value",
			})
			return
		}
		log.Info("GET /posts/" + pid + "/comments")
		pagination, err := Misc.GetPagination(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Error with pagination",
			})
			return
		}
		result, err := Controller.GetAllComments(depth, limitDepth, pagination)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if pagination.IsSet() {
			count, err := Controller.GetCommentsCount(depth, limitDepth)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			pagination.SetInformations(count)
			c.JSON(http.StatusOK, gin.H{
				"pagination": pagination,
				"comments":   result,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	})

	routerGroup.POST("", func(c *gin.Context) {
		pid := c.Param("pid")
		log.Info("POST /posts/" + pid + "/comments")
		var comment comments.Comment
		if c.BindJSON(&comment) == nil {
			_, err := Controller.CreateComment(comment)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Error while creating comment",
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Comment successfully created",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error in body format",
		})
	})

}
