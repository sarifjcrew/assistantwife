package services

import (
	"pribadi/assistantwife/webapps/models"
	"pribadi/assistantwife/webapps/repositories"
)

type ShopService struct{}

func (s *ShopService) Get(repo repositories.Repository) (interface{}, error) {
	return repo.Get()
}

func (s *ShopService) Save(data interface{}, repo repositories.Repository) error {
	err := repo.Save(data)

	if err == nil {
		model := data.(*models.ShopModel)
		saldoRepo := repositories.SaldoRepo{}
		saldoRows, err := saldoRepo.Get()
		if err == nil {
			saldoMs := saldoRows.([]models.SaldoModel)
			if len(saldoMs) > 0 {
				saldoM := saldoMs[0]
				saldoM.Saldo -= model.Harga
				saldoRepo.Save(&saldoM)
			}
		}
	}

	return err
}

func (s *ShopService) DeleteById(repo repositories.Repository, id interface{}) error {
	return repo.DeleteById(id)
}
