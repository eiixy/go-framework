package routes

import (
	"eii/internal/account/controllers"
	"github.com/gin-gonic/gin"
)

func v1(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		accounts(v1)
	}
}

func accounts(group *gin.RouterGroup) {
	group.Group("accounts")
	{
		group.GET("/", controllers.GetAccounts)
		group.POST("/", controllers.CreateAccount)
		group.GET("/:id", GetAccount)
		group.PUT("/:id", UpdateAccount)
		group.DELETE("/:id", DeleteAccount)
	}
}

func GetAccounts(context *gin.Context) {
	id := context.Param("id")
	context.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    id,
	})
}

func CreateAccount(context *gin.Context) {

}

func GetAccount(context *gin.Context) {

}

func UpdateAccount(context *gin.Context) {

}

func DeleteAccount(context *gin.Context) {

}
