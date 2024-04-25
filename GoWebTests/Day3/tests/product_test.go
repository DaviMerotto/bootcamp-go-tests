package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davimerotto/web-server/cmd/server/handler"
	"github.com/davimerotto/web-server/internal/entities"
	"github.com/davimerotto/web-server/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) Read(data interface{}) ([]entities.Product, error) {
	args := m.Called(&data)
	return args.Get(0).([]entities.Product), args.Error(1)
}

func (m *MockDB) Write(data interface{}) error {
	return nil
}

func CreateServer() *gin.Engine {
	mockDB := new(MockDB)
	mockProducts := assembleProducts()
	mockDB.On("Read", mock.Anything).Return(mockProducts, nil)
	mockDB.On("Write", mock.Anything).Return(nil)
	rep := products.NewStoreRepository(mockDB)
	s := products.NewService(rep)
	handler := handler.NewProduct(s)
	r := gin.Default()
	rg := r.Group("/api/v1/products")

	rg.PATCH("/:id", handler.Update())
	rg.PUT("/:id", handler.UpdateFull())
	rg.POST("/", handler.Create())
	rg.DELETE("/:id", handler.Delete())
	rg.GET("/", handler.GetAll())
	return r
}

func CreateRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "1234")

	return req, httptest.NewRecorder()
}

func assembleProducts() []entities.Product {
	return []entities.Product{
		{
			Id:            1,
			Name:          "Product 1",
			Color:         "Red",
			Price:         10.0,
			Stock:         10,
			Code:          "123",
			Published:     true,
			Creation_date: "2021-09-01",
		},
		{
			Id:            2,
			Name:          "Product 2",
			Color:         "Red",
			Price:         11.0,
			Stock:         11,
			Code:          "1234",
			Published:     true,
			Creation_date: "2021-09-01",
		},
	}
}
func TestGetAllProducts(t *testing.T) {
	//products := assembleProducts()
	t.Run("Get all products", func(t *testing.T) {
		req, w := CreateRequestTest(http.MethodGet, "/api/v1/products/", "")
		r := CreateServer()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestCreateProduct(t *testing.T) {
	t.Run("Create product", func(t *testing.T) {
		req, w := CreateRequestTest(http.MethodPost, "/api/v1/products/", `{"name":"Product 1","color":"red","price":10.0,"stock":10,"code":"123","published":true,"creation_date":"2021-09-01"}`)
		r := CreateServer()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	})
}
func TestUpdateProduct(t *testing.T) {
	t.Run("Update product", func(t *testing.T) {
		req, w := CreateRequestTest(http.MethodPatch, "/api/v1/products/1", `{"name":"Product 1 New"}`)
		r := CreateServer()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUpdateFullProduct(t *testing.T) {
	t.Run("Update product", func(t *testing.T) {
		req, w := CreateRequestTest(http.MethodPatch, "/api/v1/products/1", `{"name":"Product 1 New", "color":"blue", "price":20.0, "stock":20, "code":"456", "published":true, "creation_date":"2021-09-02"}`)
		r := CreateServer()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Delete product", func(t *testing.T) {
		req, w := CreateRequestTest(http.MethodDelete, "/api/v1/products/1", "")
		r := CreateServer()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusAccepted, w.Code)
	})
}
