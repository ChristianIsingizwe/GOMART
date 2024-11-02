package helpers

import (
	"os"
	"regexp"

	"github.com/go-playground/validator/v10"
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

func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateAccessToken(userID string, role string, tokenVersion int) (string, error) {
	claims := jwt.MapClaims{
		"userID":       userID,
		"role":         role,
		"tokenVersion": tokenVersion,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessTokenSecretKey))
}

func GenerateRefreshToken(userID string, role string, tokenVersion int) (string, error) {
	claims := jwt.MapClaims{
		"userID":       userID,
		"role":         role,
		"tokenVersion": tokenVersion,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(refreshTokenSecretKey))
}


func StrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*()_+{}:"|<>?,./;'[\]\\-]`).MatchString(password)
	isAtLeast8 := len(password) >=8

	return hasUpper && hasLower && hasNumber && hasSpecial && isAtLeast8
}
