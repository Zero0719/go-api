package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/Zero0719/go-api/pkg/app"
	"github.com/Zero0719/go-api/pkg/config"
	"github.com/Zero0719/go-api/pkg/logger"
	"github.com/gin-gonic/gin"
	jwtpkg "github.com/golang-jwt/jwt"
)

var (
	ErrTokenExpired     = errors.New("令牌已过期")
	ErrTokenExpiredMaxRefresh = errors.New("令牌已过最大刷新时间")
	ErrTokenMalformed   = errors.New("令牌格式错误")
	ErrTokenInvalid     = errors.New("请求令牌无效")
	ErrHeaderEmpty      = errors.New("需要认证才能访问")
	ErrHeaderMalformed  = errors.New("Authorization 请求头格式错误")
)

type JWT struct {
	SignKey []byte
	MaxRefresh time.Duration
}

type JWTCustomClaims struct {
	UserID string `json:"user_id"`
	ExpireAtTime int64 `json:"expire_time"`
	jwtpkg.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(config.Get[string]("app.key")),
		MaxRefresh: time.Duration(config.GetInt64("jwt.max_refresh_time")) * time.Minute,
	}
}

func (jwt *JWT) IssueToken(userID string) string {
	expireTime := jwt.expireAtTime()
	claims := JWTCustomClaims{
		UserID: userID,
		ExpireAtTime: expireTime,
		StandardClaims: jwtpkg.StandardClaims{
			NotBefore: app.TimenowInTimezone().Unix(),
			IssuedAt: app.TimenowInTimezone().Unix(),
			ExpiresAt: expireTime,
			Issuer: config.Get[string]("app.name"),
		},
	}

	token, err := jwt.createToken(claims)
	if err != nil {
		logger.LogIf(err)
		return ""
	}

	return token
}

func (jwt *JWT) ParserToken(c *gin.Context) (*JWTCustomClaims, error) {
	tokenString, parseErr := jwt.GetTokenFromHeader(c)

	if parseErr != nil {
		return nil, parseErr
	}

	token, err := jwt.parseTokenString(tokenString)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if ok {
			if validationErr.Errors == jwtpkg.ValidationErrorMalformed {
				return nil, ErrTokenMalformed
			} else if validationErr.Errors == jwtpkg.ValidationErrorExpired {
				return nil, ErrTokenExpired
			}
		}

		return nil, ErrTokenInvalid
	}

	if claims, ok := token.Claims.(*JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrTokenInvalid
}

func (jwt *JWT) parseTokenString(tokenString string) (*jwtpkg.Token, error) {
	return jwtpkg.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwtpkg.Token) (interface{}, error) {
		return jwt.SignKey, nil
	})
}

func (jwt *JWT) RefreshToken(c *gin.Context) (string, error) {
	tokenStr, parseErr := jwt.GetTokenFromHeader(c)

	if parseErr != nil {
		return "", parseErr
	}

	token, err := jwt.parseTokenString(tokenStr)

	if err != nil {
		validationErr, ok := err.(*jwtpkg.ValidationError)
		if !ok || validationErr.Errors != jwtpkg.ValidationErrorExpired {
			return "", err
		}
	}

	claims := token.Claims.(*JWTCustomClaims)

	x := app.TimenowInTimezone().Add(jwt.MaxRefresh).Unix()
	if claims.ExpireAtTime >= x {
		claims.StandardClaims.ExpiresAt = jwt.expireAtTime()
		return jwt.createToken(*claims)
	}

	return "", ErrTokenExpiredMaxRefresh
}

func (jwt *JWT) createToken(claims JWTCustomClaims) (string, error) {
	token := jwtpkg.NewWithClaims(jwtpkg.SigningMethodHS256, claims)
	return token.SignedString(jwt.SignKey)
}

func (jwt *JWT) expireAtTime() int64 {
	timenow := app.TimenowInTimezone()

	var expireTime int64
	if config.GetBool("app.debug") {
		expireTime = config.GetInt64("jwt.debug_expire_time")
	} else {
		expireTime = config.GetInt64("jwt.expire_time")
	}

	expire := time.Duration(expireTime) * time.Minute
	return timenow.Add(expire).Unix()
}

func (jwt *JWT) GetTokenFromHeader(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", ErrHeaderEmpty
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return "", ErrHeaderMalformed
	}

	return parts[1], nil
}