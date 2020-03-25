package jwt

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
)

func Create(user *structs.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ID":       user.ID,
		"Username": user.Username,
	})

	strToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", nil

	}

	return strToken, nil

}

func IsValid(token string) (interface{}, error) {
	if token == "" {
		return nil, err.ErrTokenForgotten

	}

	parsedToken, parseError := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err.ErrSigningMethod
		}

		return []byte(os.Getenv("JWT_SECRET")), nil

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
