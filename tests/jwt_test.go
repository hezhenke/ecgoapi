package tests

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestJwtEncode(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":      1234,
		"exp":      time.Now().Unix() + 43200*60,
		"platform": "wechatminiprogram",
	})
	var tokenString, err = token.SignedString([]byte("iamascrety"))
	if(err !=nil){
		fmt.Println("abort in product vip cang")
	}
	fmt.Println(tokenString)
}

func TestJwtDecode(t *testing.T)  {
	var tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDU5NTQxNTQsInBsYXRmb3JtIjoid2VjaGF0bWluaXByb2dyYW0iLCJ1aWQiOjEyMzR9.TWNRD301waNa7_UZtqmsCMl_sEsdDZ_DRbCgOLvpEMk"

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("iamascrety"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("You look nice today")
		fmt.Println(claims)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
		} else {
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}
}
