package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_midlwares"
	users_controller "go-microspace-service/controllers/users"
)

func ApiRoutes(r *fiber.App) {

	onlyUserGroup := r.Group("/user", gjmt_midlwares.AuthMiddleware())
	onlyUserGroup.Get("/", users_controller.GetUser)
	onlyUserGroup.Get("/profile", users_controller.GetUser)

}
