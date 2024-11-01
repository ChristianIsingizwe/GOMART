package helpers

import (
	"os"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var refreshTokenSecretKey = os.Getenv("REFRESH_TOKEN_SECRET_KEY")
var accessTokenSecretKey = os.Getenv("ACCESS_TOKEN_SECRET_KEY")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}


func CheckPassword(hashedPassword, password string) error{
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateAccessToken(userID string, tokenVersion int) (string, error){
	claims := jwt.MapClaims{
		"userID": userID,
		"tokenVersion": tokenVersion,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenSecretKey))
}

func GenerateRefreshToken(userID string, tokenVersion int) (string, error){
	claims := jwt.MapClaims{
		"userID": userID,
		"tokenVersion": tokenVersion,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenSecretKey))
}

