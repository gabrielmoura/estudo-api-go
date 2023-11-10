package routes

import (
	"net/http"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gabrielmoura/estudo-api-go/pkg/token"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var LRequest LoginRequest
	if err := c.ShouldBindJSON(&LRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de login inválidos"})
		return
	}

	user := LRequest.checkCredentials()

	// Adicione aqui a lógica de autenticação, verificando o usuário e a senha.
	// Neste exemplo, estamos verificando um usuário estático.
	if user != nil {
		tokenString, claims, err := token.CreateToken(user.ID, []string{"ADMIN", "USER"})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o token", "stack": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":      tokenString,
			"plainToken": claims,
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
}

func (L LoginRequest) checkCredentials() *entity.User {
	user, err := db.GetOneUserByEmail(db.Con, L.Username)
	if err != nil {
		logger.Logger.Sugar().Warnf("Erro ao buscar usuário: %s", err)
		return nil
	}

	if !user.ValidatePassword(L.Password) {
		logger.Logger.Warn("senha incorreta")
		return nil
	}
	return user
}
