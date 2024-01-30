package routes

import (
	"net/http"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/infra/logger"
	"github.com/gabrielmoura/estudo-api-go/internal/dto"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gabrielmoura/estudo-api-go/pkg/token"
	"github.com/gin-gonic/gin"
)

// LoginHandler godoc
// @Summary Login
// @Description Faz login na aplicação
// @Tags login
// @Accept  json
// @Produce  json
// @Param login body dto.LoginRequest true "Login"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var LRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&LRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de login inválidos"})
		return
	}

	user := checkCredentials(LRequest)

	// Adicione aqui a lógica de autenticação, verificando o usuário e a senha.
	// Neste exemplo, estamos verificando um usuário estático.
	if user != nil {
		tokenString, claims, err := token.CreateToken(user.ID, []string{"ADMIN", "USER"}, c.ClientIP())
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: "Erro ao criar o token", Stack: err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.LoginResponse{Token: tokenString, PlainToken: claims})
		return
	}

	c.JSON(http.StatusUnauthorized, dto.ErrorResponse{Error: "Credenciais inválidas"})
}

// RegisterHandler godoc
// @Summary Registra um novo usuário
// @Description Registra um novo usuário
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body entity.User true "User"
// @Success 200 {object} dto.MessageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var RRequest dto.RegisterRequest
	if err := c.ShouldBindJSON(&RRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Dados inválidos"})
		return
	}
	newUser, _ := entity.NewUser(RRequest.Name, RRequest.Email, RRequest.Password)

	if _, err := db.InsertUser(db.Con, newUser); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Erro ao cadastrar o usuário"})
		return
	}
	c.JSON(http.StatusOK, dto.MessageResponse{Message: "Usuário cadastrado com sucesso"})
}

// CheckCredentials
// @Description Verifica as credenciais do usuário
func checkCredentials(L dto.LoginRequest) *entity.User {
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
