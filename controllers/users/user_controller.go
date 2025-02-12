package users_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/slavaWins/go-jwt-microservice-template/gjmt_models"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_midlwares"
	"net/http"
)

// @Summary Проверка что пользователь авторизирован
// @Description Только с jwt токеном
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Response 200 {object} gjmt_models.Response[gjmt_models.User]
// @Response 401 {object} gjmt_models.ResponseErrorType
// @Router /user [get]
func GetUser(c *gin.Context) {

	user, err := gjmt_midlwares.GetAuthUser(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gjmt_models.ResponseWithError(err.Error()))
	}

	c.JSON(http.StatusOK, gjmt_models.ResponseWithValue(user))
}
