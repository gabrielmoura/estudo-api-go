package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielmoura/estudo-api-go/infra/db"
	"github.com/stretchr/testify/assert"
)

func TestStatusAllProducts(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/product"
	r := SetupRoutes()
	r.GET(rName, getAllProduct)

	req, _ := http.NewRequest("GET", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status Error: Valor recebido foi %d, e valor esperado é %d em %s", res.Code, http.StatusOK, rName)
	// }
}

func TestStatusOneProduct(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/product/d1a58243-5c20-4296-b315-3677f0d2b0c1"
	r := SetupRoutes()
	r.GET("/product/:id", getOneProduct)

	req, _ := http.NewRequest("GET", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	// if res.Code != http.StatusOK {
	// 	t.Fatalf("Status Error: Valor recebido foi %d, e valor esperado é %d em %s", res.Code, http.StatusOK, rName)
	// }
}

func TestUpdateOneProduct(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/product/d1a58243-5c20-4296-b315-3677f0d2b0c1"
	r := SetupRoutes()
	r.PUT("/product/:id", updateProduct)

	userData := map[string]interface{}{
		"name": "Abacaxi",
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

func TestCreateOneProduct(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/product"
	r := SetupRoutes()
	r.POST("/product", postProduct)

	userData := map[string]interface{}{
		"name":  "Abacate",
		"price": 1.99,
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

func TestDeleteOneProduct(t *testing.T) {
	db.Conn("sqlite", "../../cmd/server/sqlite.db")
	// UserMock()

	rName := "/product/d1a58243-5c20-4296-b315-3677f0d2b0c1"
	r := SetupRoutes()
	r.DELETE("/product/:id", deleteProduct)

	req, _ := http.NewRequest("DELETE", rName, nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
