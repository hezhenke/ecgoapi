package utils

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func Encode(uid int)  (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Unix() + 43200 * 60,
		"platform":"wechatminiprogram",
	})
	tokenSecret := os.Getenv("TOKEN_SECRET")
	return token.SignedString([]byte(tokenSecret))
}
