package db_service

import (
	"fmt"
	"github.com/slavaWins/go-jwt-microservice-template/pkg/gjmt_db_service"
	"go-microspace-service/models"
)

func Migrate() {

	fmt.Println("[Migrate] Migratation ")
	db := gjmt_db_service.Connect()

	db.AutoMigrate(&models.PostWall{})

}
