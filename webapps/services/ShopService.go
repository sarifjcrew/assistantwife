package services

import (
	"pribadi/assistantwife/webapps/repositories"
)

type ShopService struct{}

func (s *ShopService) Get(repo repositories.Repository) (interface{}, error) {
	return repo.Get()
}

func (s *ShopService) DeleteById(repo repositories.Repository, id interface{}) error {
	return repo.DeleteById(id)
}
