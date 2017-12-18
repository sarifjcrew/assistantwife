package services

import (
	"pribadi/assistantwife/webapps/repositories"
)

type SaldoService struct{}

func (s *SaldoService) Get(repo repositories.Repository) (interface{}, error) {
	return repo.Get()
}

func (s *SaldoService) Save(data interface{}, repo repositories.Repository) error {
	return repo.Save(data)
}
