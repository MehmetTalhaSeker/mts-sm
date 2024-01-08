package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenPayload struct {
	ID       string
	Username string
}

func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration("10h") // TODO

	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(v).Unix(),
		"username": payload.Username,
		"id":       payload.ID,
	})

	token, err := t.SignedString([]byte("bla bla")) // TODO
	if err != nil {
		panic(err)
	}

	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("bla bla"), nil // TODO
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("something went wrong")
	}

	return &TokenPayload{
		Username: username,
		ID:       id,
	}, nil
}
