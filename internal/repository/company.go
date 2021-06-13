package repository

import (
	"securities_admin/db"
	"securities_admin/internal/model"
)

type ICompanyRepository interface {
	GetCompanies(page int, query string) ([]model.Company, error)
}

type CompanyRepository struct {
}

func (r *RequestRepository) GetCompanies(page int, query string) ([]model.Company, error) {
	var result []model.Company
	tx := db.DB
	if query != ""{
		// TODO: fulltext search
	}

	res := tx.Limit(PerPage).Offset((page - 1) * PerPage).Find(&result)
	return result, res.Error
}

