package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/micor-lzy/service-user/user"
	"time"
)

var (
	key = []byte("mySuperSecretKeyLol")
)

type CustomClaims struct {
	User *user.User
	jwt.StandardClaims
}
type Authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *user.User) (string, error)
}
type TokenService struct {
	repo Repository
}

func (srv *TokenService) Decode(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

	// Validate the token and return the custom claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func (srv *TokenService) Encode(user *user.User) (string, error) {
	expireToken := time.Now().Add(time.Hour * 72).Unix()
	claims := CustomClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "service.user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(key)
}
