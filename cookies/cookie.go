package cookies

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSecret = []byte("mts/auth")

// New creates and return two cookies: access and refresh.
// Access cookie's max age is 1 minute.
// Refresh cookie's max age is 1 hour.
// Argument value is converted into JWT-token and setted as value
// in cookies.
func New(value string) ([]http.Cookie, error) {
	// converting value to JWT-token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": value,
	})
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}

	cookies := make([]http.Cookie, 2)

	accessCookie := http.Cookie{
		Name:   "Access",
		Value:  tokenString,
		MaxAge: int(1 * time.Minute / 1e9),
	}
	refreshCookie := http.Cookie{
		Name:   "Refresh",
		Value:  tokenString,
		MaxAge: int(1 * time.Hour / 1e9),
	}

	cookies[0], cookies[1] = accessCookie, refreshCookie
	return cookies, nil
}
