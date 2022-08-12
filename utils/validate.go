package utils

import (
	"fmt"
	"minitiktok/utils/logger"
	"time"

	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId int) string {
	mySigningKey := []byte("ItWillBeDone")

	// Create the Claims
	claims := MyCustomClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 1000, //过期时间
			Issuer:    "login",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		println("something goes wrong with generating token")
	}
	return ss // 生成的加密字符串
}


// 验证token的有效性

func ValidateToken(tokenString string) (int , error){
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ItWillBeDone"), nil
	})

	var userId int

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		userId = claims.UserId
	} else {
		logger.Error(fmt.Sprintf("token过期:%s",tokenString))
		// TODO: 设置token过期逻辑
		return 0, err
	}

	return userId, nil
}