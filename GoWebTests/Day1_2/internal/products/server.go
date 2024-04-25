package products

import "github.com/davimerotto/web-server/internal/entities"

type Service interface {
	GetAll() ([]entities.Product, error)
	Create(p entities.Product) (entities.Product, error)
	Delete(id uint) error
	UpdateFull(p entities.Product) (entities.Product, error)
	Update(id uint, p entities.Product) (entities.Product, error)
}

type service struct {
	repository Repository
}

func (s *service) GetAll() ([]entities.Product, error) {
	prods, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return prods, nil
}

func (s *service) Create(p entities.Product) (entities.Product, error) {
	product, err := s.repository.Create(p)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateFull(p entities.Product) (entities.Product, error) {
	product, err := s.repository.UpdateFull(p)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func (s *service) Update(id uint, p entities.Product) (entities.Product, error) {
	product, err := s.repository.Update(id, p)
	if err != nil {
		return entities.Product{}, err
	}
	return product, nil
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}
