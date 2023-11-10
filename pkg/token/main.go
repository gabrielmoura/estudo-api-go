package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/gabrielmoura/estudo-api-go/configs"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id string, permission []string) (string, *jwt.MapClaims, error) {
	expiresIn := time.Second * time.Duration(configs.Conf.JwtExperesIn)

	// Criar as claims
	claims := jwt.MapClaims{
		"sub": id,
		"iss": configs.Conf.AppName,
		"exp": time.Now().Add(expiresIn).Unix(),
		"iat": time.Now().Unix(),
		"p":   strings.Join(permission, ","),
	}

	// Criar e assinar o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	nToken, err := token.SignedString([]byte(configs.Conf.JWTSecret))

	// Retornar o token e as claims
	return nToken, &claims, err
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica o algoritmo de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inválido: %v", token.Header["alg"])
		}

		return []byte(configs.Conf.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("token inválido")
}
