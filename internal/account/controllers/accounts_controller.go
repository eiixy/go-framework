package controllers

import (
	"eii/internal/account/models/account"
	"eii/internal/account/models/scopes"
	"eii/internal/account/services"
	"eii/pkg/database"
	"github.com/gin-gonic/gin"
)

func GetAccounts(context *gin.Context) {
	accounts := database.DB.Scopes(scopes.Paginate(context)).Select(&account.Account{})
	success(context, accounts)
}
func CreateAccount(context *gin.Context) {
	account := &account.Account{
		Email: context.PostForm("email"),
	}
	services.Register(*account)
	success(context, nil)
}
