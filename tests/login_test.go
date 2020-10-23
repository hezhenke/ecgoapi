package tests

import (
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"testing"
	"time"
)

func TestGetArea(t *testing.T)  {

	database := infrastructure.GetDb()
	config := models.Config{}
	err := database.Model(&models.Config{}).Where("type=?","oauth").Where("status=?",1).Where("code=?","wechat.wxa").First(&config).Error
	if err != nil{
		log.Fatal(err.Error())
	}
	fmt.Println(config.Config)
}

func TestWeixinLogin(t *testing.T)  {
	/*
	session,err := services.WxLogin("021nlh000gaNdK1Fk2000d0Q0V3nlh0Q")
	fmt.Println(session)
	fmt.Print(session.OpenID == "")
	fmt.Println(err)
	 */
	maxLen := base64.URLEncoding.EncodedLen(len("abcdef"))
	dst :=  make([]byte,maxLen)
	base64.URLEncoding.Encode(dst,[]byte("abcdef"))
	fmt.Println(string(dst))
	dst = []byte("eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0")
	decodeLen := base64.URLEncoding.DecodedLen(len(dst))
	decodeS := make([]byte,decodeLen+10)
	n,err := base64.URLEncoding.Decode(decodeS,dst)
	fmt.Println(err.Error())
	fmt.Println(string(decodeS[:n]))
}

func TestJwt(t *testing.T){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("hmacSampleSecret"))

	base64str := jwt.EncodeSegment([]byte("123456"))
	fmt.Println(base64str)
	destr,_ := jwt.DecodeSegment("eyJleHAiOjI1OTIwMzAsInBsYXRmb3JtIjoid2VjaGF0bWluaXByb2dyYW0iLCJ1aWQiOjF9")
	fmt.Println(string(destr))

	fmt.Println(tokenString, err)
}

func TestPrice(t *testing.T){
	now := time.Now().Unix()
	fmt.Println(now)
}