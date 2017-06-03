package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/twalkapp/server/models/users"
)

type Claims struct {
	Profile	users.Profile `json:"user"`
	jwt.StandardClaims
}
