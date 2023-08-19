package routes

import (
	"api-gin2/api/controllers"

	"github.com/gin-gonic/gin"
)

func AppRoutes(router *gin.Engine) *gin.RouterGroup {

	v1 := router.Group("/v1")
	{
		v1.GET("/users/:id", controllers.FindById)
		v1.GET("/users", controllers.FindAll)
		v1.POST("/user", controllers.CreateTweet)
		v1.PUT("/user/:id", controllers.UpdateTweet)
		v1.DELETE("/user/:id", controllers.Delete)
	}

	return v1

}
