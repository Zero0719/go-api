package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (tokenString string, refreshTokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(Config.GetInt("auth.jwt_exp")) * time.Minute).Unix(),
	})
	tokenString, err = token.SignedString([]byte(Config.GetString("auth.jwt_secret")))
	if err != nil {
		return "", "", err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(Config.GetInt("auth.jwt_refresh_exp")) * time.Minute).Unix(),
	})
	refreshTokenString, err = refreshToken.SignedString([]byte(Config.GetString("auth.jwt_refresh_secret")))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
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

func ParseRefreshToken(refreshTokenString string) (int, error) {
	refreshToken, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (any, error) {
		return []byte(Config.GetString("auth.jwt_refresh_secret")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}
	if claims, ok := refreshToken.Claims.(jwt.MapClaims); ok {
		return int(claims["userId"].(float64)), nil
	} else {
		return 0, err
	}
}

func RefreshToken(refreshTokenString string) (tokenString string, newRefreshTokenString string, err error) {
	userId, err := ParseRefreshToken(refreshTokenString)
	if err != nil {
		return "", "", err
	}
	tokenString, newRefreshTokenString, err = GenerateToken(userId)
	if err != nil {
		return "", "", err
	}
	return tokenString, newRefreshTokenString, nil
}