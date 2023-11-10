package middleware

import (
	"net/http"
	"strings"

	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/pkg/token"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação ausente"})
			c.Abort()
			return
		}

		claims, err := token.ValidateToken(strings.Split(tokenString, " ")[1])
		if err != nil {
			logger.Logger.Sugar().Infof("Token Inválido: %s", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}

// func ProtectedHandler(c *gin.Context) {
// 	claims := c.MustGet("claims").(jwt.MapClaims)
// 	c.JSON(http.StatusOK, gin.H{"usuario": claims["sub"], "mensagem": "Rota protegida!"})
// }
