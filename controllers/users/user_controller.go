package users_controller

import (
	"github.com/gofiber/fiber/v3"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_midlwares"
)

// @Summary Проверка что пользователь авторизирован, получение профиля
// @Description Только с jwt токеном
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Response 200 {object} gjmt_models.Response[gjmt_models.User]
// @Response 401 {object} gjmt_models.ResponseErrorType
// @Router /user [get]
func GetUser(c fiber.Ctx) error {

	user, err := gjmt_midlwares.GetAuthUser(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(gjmt_models.ResponseWithError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(gjmt_models.ResponseWithValue(user))

}
