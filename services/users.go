package services

import (
	"crypto/md5"
	"ecshopGoApi/dtos"
	"ecshopGoApi/infrastructure"
	"ecshopGoApi/models"
	"ecshopGoApi/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)


const (
	VENDOR_WEIXIN = 1
	VENDOR_WEIBO  = 2
	VENDOR_QQ     = 3
	VENDOR_TAOBAO = 4
	VENDOR_WXA    = 5
)


//WxSession 微信登陆接口返回session
type WxSession struct {
	SessionKey string `json:"session_key"`
	ExpireIn   int    `json:"expires_in"`
	OpenID     string `json:"openid"`
	ErrorCode  int `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}
type WxaConfig struct {
	AppId string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

func Auth(req dtos.AuthRequest)(result map[string]interface{},err error){
	user := models.Users{}
	if req.Vendor == VENDOR_WXA {
		//1.微信登录获取session_Key和openId
		session,err := WxLogin(req.JsCode)
		if err != nil {
			return nil,err
		}
		openId := session.OpenID
		//2.检查当前openId是否绑定，没有绑定创建用户；最后根据userId查询用户数据
		var userId int
		userId,err = checkBind(openId)
		if err != nil || userId <= 0{
			user,err = createAuthUser(int8(req.Vendor),openId,"wxa_"+openId,0,"wxa")
			if err != nil{
				return nil,err
			}
			userId = user.UserId
		}
		//3.查询用户信息并返回
		database := infrastructure.GetDb()
		err = database.Model(&models.Users{}).First(&user,userId).Error
		token,_ := utils.Encode(userId)
		return dtos.CreateAuthDto(false,openId,token,user),err
	}
	err = errors.New(fmt.Sprintf("Vendor错误：%d",req.Vendor))
	return nil,err

}

//WxLogin 微信用户授权
func  WxLogin(jscode string) (session WxSession, err error) {
	client := &http.Client{}

	database := infrastructure.GetDb()
	config := models.Config{}
	err = database.Model(&models.Config{}).Where("type=?","oauth").Where("status=?",1).Where("code=?","wechat.wxa").First(&config).Error
	if err !=nil{
		return
	}

	wxaConfig := WxaConfig{}
	if err = json.Unmarshal([]byte(config.Config),&wxaConfig);err !=nil{
		return
	}

	//生成要访问的url
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", wxaConfig.AppId, wxaConfig.AppSecret, jscode)
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	body, err := ioutil.ReadAll(response.Body)
	jsonStr := string(body)
	//解析json
	if err = json.Unmarshal(body, &session); err != nil {
		session.SessionKey = jsonStr
		return session, err
	}
	if session.ErrorCode != 0 {
		err = errors.New(session.ErrorMessage)
		return session,err
	}
	return session, err
}

func checkBind(openId string) (userId int,err error) {
	database := infrastructure.GetDb()
	sns := models.Sns{}
	if err = database.Model(&models.Sns{}).Where("open_id=?",openId).First(&sns).Error;err != nil{
		return
	}
	return sns.UserId,nil
}

func createAuthUser(vendor int8,openId string,nickName string,gender int,prefix string) (user models.Users,err error) {
	database := infrastructure.GetDb()
	userName := genUserName(prefix)
	user = models.Users{
		UserName:userName,
		Email: userName + "@sns.user",
		Password: setPassword(string(time.Now().Second())),
		RegTime: time.Now().Second(),
		UserRank: 0,
		Sex: gender,
		Alias: nickName,
		MobilePhone: "",
		RankPoints: 0,
	}
	err = database.Save(&user).Error
	if err != nil {
		return
	}
	sns := models.Sns{
		UserId: user.UserId,
		OpenId: openId,
		Vendor: vendor,
	}
	err = database.Save(&sns).Error
	return user,err
}

func genUserName(prefix string) string  {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%v_%d%d",prefix,time.Now().Second(),1000+rand.Intn(8999))
}

func setPassword(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}