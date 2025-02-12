package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_db_service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"post-service/docs"
	_ "post-service/docs" // Нужен для подключения сгенерированных Swagger-документов
	"post-service/pkg/db_service"
	"post-service/routes"
)

//go:generate swag init --parseDependency --parseInternal --parseDepth 2

//Запустить
//go:generate go run main.go

// @title API
// @version 1.0
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Jwt Токен пользователя
func main() {

	fmt.Println("============ Start ============ ")
	fmt.Println(os.Getenv("APP_NAME"))
	fmt.Println(os.Getenv("APP_PORT"))

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Не настроен env! %w", err)
		return
	}

	gjmt_db_service.Connect()
	gjmt_db_service.Migrate()

	db_service.Migrate()

	r := gin.Default()

	routes.ApiRoutes(r)

	// Маршрут для Swagger-документации
	if os.Getenv("APP_PRODUCTION") == "false" {
		docs.SwaggerInfo.Title = os.Getenv("APP_NAME")

		docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_USE_CUSTOM_BASEPATH")

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.Run(":" + os.Getenv("APP_PORT"))

}
