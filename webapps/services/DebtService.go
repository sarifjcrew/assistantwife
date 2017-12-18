package services

import (
	"pribadi/assistantwife/webapps/repositories"
)

type DebtService struct{}

func (s *DebtService) Get(repo repositories.Repository) (interface{}, error) {
	return repo.Get()
}

func (s *DebtService) Save(repo repositories.Repository) error {
	return repo.Save()
}
func (s *DebtService) DeleteById(repo repositories.Repository, id interface{}) error {
	return repo.DeleteById(id)
}
