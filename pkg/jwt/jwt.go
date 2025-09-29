package jwt

import (
	"go-api/internal/config"
	"time"

	jwtPkg "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId int) (tokenString string, refreshTokenString string, err error) {
	jwtConfig := config.Get().Auth
	token := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, jwtPkg.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(jwtConfig.JwtExp) * time.Minute).Unix(),
	})
	tokenString, err = token.SignedString([]byte(jwtConfig.JwtSecret))
	if err != nil {
		return "", "", err
	}
	refreshToken := jwtPkg.NewWithClaims(jwtPkg.SigningMethodHS256, jwtPkg.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Duration(jwtConfig.JwtRefreshExp) * time.Minute).Unix(),
	})
	refreshTokenString, err = refreshToken.SignedString([]byte(jwtConfig.JwtRefreshSecret))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func ParseToken(tokenString string) (int, error) {
	config := config.Get().Auth
	token, err := jwtPkg.Parse(tokenString, func(token *jwtPkg.Token) (any, error) {
		return []byte(config.JwtSecret), nil
	}, jwtPkg.WithValidMethods([]string{jwtPkg.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}
	
	if claims, ok := token.Claims.(jwtPkg.MapClaims); ok {
		// 返回int类型的userId，将interface{}转换为int
		return int(claims["userId"].(float64)), nil
	} else {
		return 0, err
	}
}

func ParseRefreshToken(refreshTokenString string) (int, error) {
	config := config.Get().Auth
	refreshToken, err := jwtPkg.Parse(refreshTokenString, func(token *jwtPkg.Token) (any, error) {
		return []byte(config.JwtRefreshSecret), nil
	}, jwtPkg.WithValidMethods([]string{jwtPkg.SigningMethodHS256.Alg()}))
	if err != nil {
		return 0, err
	}
	if claims, ok := refreshToken.Claims.(jwtPkg.MapClaims); ok {
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