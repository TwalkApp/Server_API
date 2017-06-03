package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/twalkapp/server/middlewares"
)

func InitRoutes(routerGroup *gin.RouterGroup ) {
	SetAuthRoutes(routerGroup)

	needAuthGroup := routerGroup.Group("")
	needAuthGroup.Use(middlewares.CheckAuth())
	{
		usersGroup := needAuthGroup.Group("/users")
		SetUsersRoutes(usersGroup)

		userGroup := usersGroup.Group("/:uid")
		SetUserRoutes(userGroup)

		postsGroup := needAuthGroup.Group("/posts")
		SetPostsRoutes(postsGroup)

		postGroup := postsGroup.Group("/:pid")
		SetPostRoutes(postGroup)

		commentsGroup := postGroup.Group("/comments")
		SetCommentsRoutes(commentsGroup)

		commentGroup := commentsGroup.Group("/:cid")
		SetCommentRoutes(commentGroup)
	}
}
