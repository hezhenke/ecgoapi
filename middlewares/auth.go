package middlewares

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"ecshopGoApi/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if exists && user.(models.Users).UserId != 0 {
			return
		} else {
			err, _ := c.Get("authErr")
			_ = c.AbortWithError(http.StatusUnauthorized, err.(error))
			return
		}
	}
}

func UserLoaderMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("X-" + os.Getenv("APP_NAME") + "-Authorization")

		if tokenString != "" {
			claims,err := utils.Decode(tokenString)
			if err != nil{
				return
			}
			userId := uint(claims["user_id"].(float64))

			if userId != 0 {
				user := models.Users{}
				database := infrastructure.GetDb()
				// We always need the Roles to be loaded to make authorization decisions based on Roles
				database.Preload("Roles").First(&user, userId)
				c.Set("currentUser", user)
				c.Set("currentUserId", user.UserId)
			}
		}
	}
}
