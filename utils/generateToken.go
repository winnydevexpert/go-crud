package utils

import (
	"go-crud-learn/config"
	"go-crud-learn/model"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtToken struct {
	jwt.RegisteredClaims
}

func GenerateAccessToken(payload model.User, jwtType string) (*string, error) {
	tokenDuration, _ := time.ParseDuration(config.AppConfig.AccessTokenExpire)

	claim := JwtToken{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	signedToken, errSignedToken := token.SignedString([]byte(os.Getenv(jwtType)))
	if errSignedToken != nil {
		return nil, errSignedToken
	}
	return &signedToken, nil
}

func GenerateRefreshToken(payload model.User, jwtType string) (*string, error) {
	tokenDuration, _ := time.ParseDuration(config.AppConfig.RefreshTokenExpire)

	claim := JwtToken{
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	signedToken, errSignedToken := token.SignedString([]byte(os.Getenv(jwtType)))
	if errSignedToken != nil {
		return nil, errSignedToken
	}
	return &signedToken, nil
}
