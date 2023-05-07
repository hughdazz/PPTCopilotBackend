package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func CreateToken(user_id int) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	token_exp, _ := beego.AppConfig.String("Tokenexp")
	tokenexp, _ := strconv.Atoi(token_exp)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["user_id"] = user_id
	token.Claims = claims
	token_secrets, _ := beego.AppConfig.String("TokenSecrets")
	tokenString, _ := token.SignedString([]byte(token_secrets))
	return tokenString
}

func CheckToken(tokenString string) int {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		token_secrets, _ := beego.AppConfig.String("TokenSecrets")
		return []byte(token_secrets), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	return user_id
}

func GetUserId(tokenString string) int {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		token_secrets, _ := beego.AppConfig.String("TokenSecrets")
		return []byte(token_secrets), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_id := int(claims["user_id"].(float64))
	return user_id
}
