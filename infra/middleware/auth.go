package middleware

import (
	"github.com/gabrielmoura/estudo-api-go/internal/dto"
	"github.com/gabrielmoura/estudo-api-go/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"

	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Token de autenticação ausente"})
			c.Abort()
			return
		}
		strToken := strings.Split(tokenString, " ")
		if len(strToken) < 2 {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Bearer token ausente"})
		}
		claims, err := token.ValidateToken(strToken[1])

		// Verificar se o token é válido
		if err != nil {
			logger.Logger.Sugar().Infof("Token Inválido: %s", err)
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Token inválido", Stack: err.Error()})
			c.Abort()
			return
		}

		// Verificar se o IP é o mesmo
		if claims["ip"] != "" && claims["ip"] != c.ClientIP() {
			c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Token inválido"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
func checkIpOrigin(claims jwt.MapClaims, c *gin.Context) bool {
	if claims["ip"] != c.ClientIP() {
		return false
	}
	return true
}

// func ProtectedHandler(c *gin.Context) {
// 	claims := c.MustGet("claims").(jwt.MapClaims)
// 	c.JSON(http.StatusOK, gin.H{"usuario": claims["sub"], "mensagem": "Rota protegida!"})
// }
