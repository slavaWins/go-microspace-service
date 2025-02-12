package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_midlwares"
	users_controller "post-service/controllers/users"
)

func ApiRoutes(r *gin.Engine) {

	protected := r.Group("/")
	protected.Use(gjmt_midlwares.AuthMiddleware())
	{
		protected.GET("/user", users_controller.GetUser)
	}

}
