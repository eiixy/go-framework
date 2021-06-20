package services

import (
	"eii/internal/account/models/account"
	"eii/pkg/database"
)

func Register(account account.Account) {
	database.DB.Create(account)
}
