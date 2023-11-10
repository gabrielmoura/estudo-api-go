package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/gabrielmoura/estudo-api-go/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	return routes
}
func UserMock() {
	user, _ := entity.NewUser("Test", "test@example.com", "123456")
	db.InsertUser(db.Con, user)
}

func TestStatusAllUser(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/user"
	r := SetupRoutes()
	r.GET(rName, getAllUser)

	req, _ := http.NewRequest("GET", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status Error: Valor recebido foi %d, e valor esperado é %d em %s", res.Code, http.StatusOK, rName)
	// }
}
