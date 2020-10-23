package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func Decode(tokenString string)(claims jwt.MapClaims,err error){
	token,err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err !=nil{
		return nil,err
	}
	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		return claims,nil
	}else{
		return nil,errors.New("jwt decode error")
	}

}