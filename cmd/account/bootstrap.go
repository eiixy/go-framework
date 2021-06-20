package main

import (
	"eii/configs"
	"eii/internal/account/models/account"
	"eii/internal/account/routes"
	"eii/pkg/database"
	"gorm.io/gorm"
)
import "github.com/gin-gonic/gin"

func init() {
	configs.Initialize()

}

func SetUp() {
	SetUpDB()
	SetupRoutes()
}

func SetUpDB() {
	db := connectDB()
	migration(db)
}

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	routes.RegisterRoutes(router)
	return router
}

func connectDB() *gorm.DB {
	var db *gorm.DB
	db = database.Connect()
	return db
}

func migration(db *gorm.DB) {
	// 自动迁移
	err := db.AutoMigrate(
		&account.Account{},
		&account.AccountInfo{},
	)
	if err != nil {
		return
	}
}
