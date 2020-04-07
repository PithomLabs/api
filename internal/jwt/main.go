package jwt

import (
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	err "github.com/komfy/api/internal/error"
	"github.com/komfy/api/internal/structs"
)

func Create(user *structs.User) (string, error) {
	claims := jwt.MapClaims{
		"ID":       user.ID,
		"Username": user.Username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	strToken, err := token.SignedString([]byte(os.Getenv("secret")))
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

		return []byte(os.Getenv("secret")), nil

	})

	if parseError != nil {
		return nil, parseError

	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		decodedtoken := make(map[string]string)
		tokenID := int(claims["ID"].(float64))

		decodedtoken["Username"] = claims["Username"].(string)
		decodedtoken["ID"] = strconv.Itoa(tokenID)

		return decodedtoken, nil
	}

	return nil, err.ErrTokenNotValid
}
