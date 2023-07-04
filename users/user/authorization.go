package user

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtSecret   = os.Getenv("JWT_SECRET")
	sessionTime = getSessionTime()
)

func AuthUser(user User) (string, error) {
	expirationTime := time.Now().Add(sessionTime)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}

func getSessionTime() time.Duration {
	minutes, err := strconv.Atoi(os.Getenv("SESSION_TIME"))
	if err != nil {
		return 60 * time.Minute
	}
	return time.Duration(minutes)
}
