package routes

import (
	"bytes"
	"encoding/json"
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

func TestStatusOneUser(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/user/1c9dcef3-98f7-4ec8-8a04-2ae953282615"
	r := SetupRoutes()
	r.GET("/user/:id", getOneUser)

	req, _ := http.NewRequest("GET", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status Error: Valor recebido foi %d, e valor esperado é %d em %s", res.Code, http.StatusOK, rName)
	// }
}

func TestUpdateOneUser(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/user/1c9dcef3-98f7-4ec8-8a04-2ae953282615"
	r := SetupRoutes()
	r.PUT("/user/:id", updateUser)

	userData := map[string]interface{}{
		"name": "Testando",
	}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("PUT", rName, bytes.NewBuffer(jsonData))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestCreateOneUser(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/user"
	r := SetupRoutes()
	r.POST("/user", postUser)

	userData := map[string]interface{}{
		"name":     "Testando",
		"email":    "test1@example.com",
		"password": "123456",
	}
	jsonData, err := json.Marshal(userData)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest("POST", rName, bytes.NewBuffer(jsonData))
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusCreated, res.Code)
}

func TestDeleteOneUser(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/user/1c9dcef3-98f7-4ec8-8a04-2ae953282615"
	r := SetupRoutes()
	r.DELETE("/user/:id", deleteUser)

	req, _ := http.NewRequest("DELETE", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
