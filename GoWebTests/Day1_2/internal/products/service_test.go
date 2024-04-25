package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/davimerotto/web-server/internal/entities"
	"github.com/davimerotto/web-server/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
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
	myService := NewService(myRepo)
	result, err := myService.GetAll()

	assert.Equal(t, products, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	expectedError := errors.New("error")
	dbMock := store.Mock{
		Err: expectedError,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewStoreRepository(&storeStub)
	myService := NewService(myRepo)
	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestServiceCreate(t *testing.T) {
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
	product := entities.Product{
		Id:            3,
		Name:          "Product 3",
		Color:         "Green",
		Price:         30.0,
		Stock:         30,
		Code:          "789",
		Published:     true,
		Creation_date: "2021-09-03",
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
	myService := NewService(myRepo)
	result, err := myService.Create(product)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestServiceUpdate(t *testing.T) {
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
	product := entities.Product{
		Name: "Product 1 New",
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
	myService := NewService(myRepo)
	result, err := myService.Update(1, product)
	assert.Nil(t, err)
	assert.Equal(t, product.Name, result.Name)
}

func TestServiceDelete(t *testing.T) {
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
	myService := NewService(myRepo)
	err := myService.Delete(1)
	assert.Nil(t, err)
	result, err := myService.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(result))
}
