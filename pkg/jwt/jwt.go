package jwt

import (
	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"

	db "github.com/komfy/api/pkg/database"
	err "github.com/komfy/api/pkg/error"
)

// CreateToken is used inside auth endpoint
// In order to create a token for graphql auth
func CreateToken(user *db.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": user.Username,
	})

	strToken, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		log.Print(err)

	}

	return strToken

}

// IsTokenValid check if a token is valid
func IsTokenValid(token string) (interface{}, error) {
	if token == "" {
		return nil, err.ErrTokenForgotten

	}

	parsedToken, parseError := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err.ErrSigningMethod
		}

		return []byte(os.Getenv("secret")), nil

	})

	if parseError != nil {
		return nil, parseError

	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		decodedtoken := make(map[string]string)
		decodedtoken["username"] = claims["username"].(string)
		return decodedtoken, nil

	}

	return nil, err.ErrTokenNotValid

}
