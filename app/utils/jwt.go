package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(Config.GetInt("auth.jwt_exp")) * time.Minute).Unix(),
	})
	return token.SignedString([]byte(Config.GetString("auth.jwt_secret")))
}

func ParseToken(tokenString string) (int, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(Config.GetString("auth.jwt_secret")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}
	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 返回int类型的userId，将interface{}转换为int
		return int(claims["userId"].(float64)), nil
	} else {
		return 0, err
	}
}