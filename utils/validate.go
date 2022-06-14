package utils

import (
	"fmt"
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
			ExpiresAt: time.Now().Unix() + 86400, // 24h
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

// TODO: 不仅要解析token，还要验证token的有效性
func ValidateToken(tokenString string) int {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("ItWillBeDone"), nil
	})

	var userId int

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		userId = claims.UserId
	} else {
		fmt.Println(err)
	}

	return userId
}