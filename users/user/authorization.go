package user

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	jwtSecret   = os.Getenv("JWT_SECRET")
	sessionTime = GetSessionTime()
)

func AuthUser(user User) (string, error) {
	expirationTime := time.Now().Add(sessionTime)
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	return tokenString, err
}

func GetSessionTime() time.Duration {
	minutes, err := strconv.Atoi(os.Getenv("SESSION_TIME"))
	if err != nil {
		return 60 * time.Minute
	}
	return time.Duration(minutes)
}
