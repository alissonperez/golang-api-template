package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/config"
	"go.uber.org/dig"
	"net/http"
	"strings"

	"fmt"
)

type Claims struct {
	ClientId int `json:"clientId"`
	jwt.StandardClaims
}

type Auth interface {
	FromRequest(r *http.Request) (*Claims, error)
}

type authJwt struct {
	config config.Config
}

func (a authJwt) FromRequest(r *http.Request) (*Claims, error) {
	authToken, err := getAuthToken(r.Header.Get("Authorization"))
	if err != nil {
		return nil, err
	}

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(authToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.config.GetString("jwt_key")), nil
	})

	if token != nil && token.Valid {
		return claims, nil
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("Malformed token")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("Expired token")
			} else {
				return nil, fmt.Errorf("Unexpected token")
			}
		}
	}

	return nil, fmt.Errorf("Unexpected token")
}

func getAuthToken(authHeader string) (string, error) {
	if len(authHeader) < 8 {
		return "", fmt.Errorf("Invalid Authorization header")
	}

	beginning := strings.ToLower(strings.Trim(authHeader[0:6], " "))
	if beginning != "bearer" {
		return "", fmt.Errorf("Authorization header must begging with 'Bearer'")
	}

	return strings.Trim(authHeader[7:], " "), nil
}

func CreateAuth(config config.Config) Auth {
	return &authJwt{config: config}
}

func Provide(container *dig.Container) {
	container.Provide(func(config config.Config) Auth {
		return CreateAuth(config)
	})
}
