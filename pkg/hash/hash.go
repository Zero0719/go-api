package hash

import (
	"github.com/Zero0719/go-api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

func BcryptCheck(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func BcryptIsHashed(password string) bool {
	return len(password) == 60
}