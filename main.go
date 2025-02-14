package main

import (
	"fmt"
	"github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_db_service"
	"go-microspace-service/docs"
	_ "go-microspace-service/docs" // Нужен для подключения сгенерированных Swagger-документов
	"go-microspace-service/pkg/db_service"
	"go-microspace-service/routes"
	"log"
	"os"
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

	r := fiber.New()

	routes.ApiRoutes(r)

	// Маршрут для Swagger-документации
	if os.Getenv("APP_PRODUCTION") == "false" {
		docs.SwaggerInfo.Title = os.Getenv("APP_NAME")

		docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_USE_CUSTOM_BASEPATH")

		r.Get("/swagger/*", swagger.HandlerDefault)
		//r.Get("/swagger/*any", fiberSwagger.WrapHandler(swaggerFiles.Handler))
	}

	log.Fatal(r.Listen(":" + os.Getenv("APP_PORT")))

}
