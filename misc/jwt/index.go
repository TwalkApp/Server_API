package jwt

import (
	"time"
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/twalkapp/server/models/users"
	"github.com/twalkapp/server/models/auth"
	"github.com/twalkapp/server/misc/config"
)

func GenerateToken(user users.Profile) string{
	expire := time.Now().Add(time.Hour * time.Duration(config.Conf.JWT.Duration)).Unix()
	claims := auth.Claims {
		Profile: user,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expire,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.Conf.JWT.Secret))
	return signedToken
}

func ValidateToken(value string) (bool, *users.Profile, error){
	token, err := jwt.ParseWithClaims(value, &auth.Claims{}, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected siging method")
		}
		return []byte(config.Conf.JWT.Secret), nil
	})
	if err != nil {
		return false, nil, err
	}
	if clams, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return true, &clams.Profile, nil
	} else {
		return false, nil, nil
	}
}