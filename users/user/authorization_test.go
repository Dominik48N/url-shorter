package user_test

import (
	"os"
	"testing"
	"time"

	"github.com/Dominik48N/url-shorter/users/user"
	"github.com/golang-jwt/jwt"
)

func TestAuthUser(t *testing.T) {
	userTest := user.User{
		Username: "Dominik48N",
	}

	tokenString, err := user.AuthUser(userTest)
	if err != nil {
		t.Fatal(err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		t.Errorf("AuthUser(%v) = %s; want a valid JWT token", userTest, tokenString)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] != userTest.Username {
			t.Errorf("AuthUser(%v) = %s; want a token with username claim %s", userTest, tokenString, userTest.Username)
		}
		exp := int64(claims["exp"].(float64))
		if sessionTime := user.GetSessionTime(); exp < time.Now().Unix() || exp > time.Now().Add(sessionTime).Unix() {
			t.Errorf("AuthUser(%v) = %s; want a token with exp claim within %v", userTest, tokenString, sessionTime)
		}
	} else {
		t.Errorf("AuthUser(%v) = %s; want a token with valid claims", userTest, tokenString)
	}
}
