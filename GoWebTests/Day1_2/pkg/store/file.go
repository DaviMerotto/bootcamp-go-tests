package store

import (
	"encoding/json"
	"os"

	"github.com/davimerotto/web-server/internal/entities"
)

type Store interface {
	Read(data interface{}) ([]entities.Product, error)
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
)

type Mock struct {
	Data       []byte
	ReadAsUsed bool
	Err        error
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

func (f *FileStore) AddMock(mock *Mock) {
	f.Mock = mock
}

func (f *FileStore) ClearMock() {
	f.Mock = nil
}

func (f *FileStore) Read(data interface{}) ([]entities.Product, error) {
	if f.Mock != nil {
		f.Mock.ReadAsUsed = true
		if f.Mock.Err != nil {
			return []entities.Product{}, f.Mock.Err
		}
		return []entities.Product{}, json.Unmarshal(f.Mock.Data, &data)
	}
	file, err := os.ReadFile(f.FileName)
	if err != nil {
		return []entities.Product{}, err
	}
	return []entities.Product{}, json.Unmarshal(file, &data)
}

func (f *FileStore) Write(data interface{}) error {
	if f.Mock != nil {
		if f.Mock.Err != nil {
			return f.Mock.Err
		}
		var err error
		f.Mock.Data, err = json.Marshal(&data)
		if err != nil {
			return err
		}
		return nil
	}
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.FileName, fileData, 0644)
}

func NewFileStore(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{fileName, &Mock{}}
	}
	return nil
}
