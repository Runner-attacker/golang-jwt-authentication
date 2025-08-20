package initializers

import "github.com/golang-authentication-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
