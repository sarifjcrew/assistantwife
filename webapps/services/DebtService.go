package services

import (
	"pribadi/assistantwife/webapps/models"
	"pribadi/assistantwife/webapps/repositories"
)

type DebtService struct{}

func (s *DebtService) Get(repo repositories.Repository) (interface{}, error) {
	result := map[string]interface{}{}
	rows, err := repo.Get()

	result["Rows"] = rows
	result["TotalDebt"] = s.totalDebt(rows)

	return result, err
}

func (s *DebtService) Save(data interface{}, repo repositories.Repository) error {
	return repo.Save(data)
}
func (s *DebtService) DeleteById(repo repositories.Repository, id interface{}) error {
	return repo.DeleteById(id)
}

func (s *DebtService) totalDebt(rows interface{}) int64 {
	datas := rows.([]models.DebtModel)

	var debt int64
	for _, v := range datas {
		debt += v.Debt
	}

	return debt
}
