package products

import (
	"encoding/json"
	"testing"

	"github.com/davimerotto/web-server/internal/entities"
	"github.com/davimerotto/web-server/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	products := []entities.Product{
		{
			Name:          "Product 1",
			Color:         "Red",
			Price:         10.0,
			Stock:         10,
			Code:          "123",
			Published:     true,
			Creation_date: "2021-09-01",
		}, {
			Name:          "Product 2",
			Color:         "Blue",
			Price:         20.0,
			Stock:         20,
			Code:          "456",
			Published:     true,
			Creation_date: "2021-09-02",
		},
	}
	dataJson, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "products.json",
		Mock:     &dbMock,
	}

	myRepo := NewStoreRepository(&storeStub)
	resp, _ := myRepo.GetAll()

	assert.Equal(t, products, resp)
}

func TestUpdate(t *testing.T) {
	products := []entities.Product{
		{
			Id:            1,
			Name:          "Product 1",
			Color:         "Red",
			Price:         10.0,
			Stock:         10,
			Code:          "123",
			Published:     true,
			Creation_date: "2021-09-01",
		}, {
			Id:            2,
			Name:          "Product 2",
			Color:         "Blue",
			Price:         20.0,
			Stock:         20,
			Code:          "456",
			Published:     true,
			Creation_date: "2021-09-02",
		},
	}
	dataJson, _ := json.Marshal(products)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "products.json",
		Mock:     &dbMock,
	}

	myRepo := NewStoreRepository(&storeStub)
	r, err := myRepo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, len(r), 2)
	products[0].Name = "Product 1 new"
	result, err := myRepo.UpdateFull(products[0])
	assert.Nil(t, err)
	assert.Equal(t, products[0].Name, result.Name)
	assert.Equal(t, uint(1), result.Id)
	assert.Equal(t, true, storeStub.Mock.ReadAsUsed)
}

func TestStore(t *testing.T) {
	product := entities.Product{
		Name:  "Product 1 new",
		Color: "Red",
		Price: 10.0,
		Stock: 10,
		Code:  "123",
	}
	dbMock := store.Mock{
		Data: []byte("[]"),
	}
	storeStub := store.FileStore{
		FileName: "products.json",
		Mock:     &dbMock,
	}
	myRepo := NewStoreRepository(&storeStub)
	result, _ := myRepo.Create(product)

	assert.Equal(t, product.Name, result.Name)
	assert.Equal(t, product.Color, result.Color)
	assert.Equal(t, product.Price, result.Price)
	assert.Equal(t, product.Stock, result.Stock)
	assert.Equal(t, product.Code, result.Code)
	assert.Equal(t, uint(1), result.Id)
}
