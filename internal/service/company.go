package service

import (
	"securities_admin/internal/model"
	"securities_admin/internal/repository"
)

type ICompanyService interface {
	GetCompanies(page int, query string) ([]model.Company, error)
}

type CompanyService struct {
	repo repository.ICompanyRepository
}

func NewCompanyService() *RequestService{
	return &RequestService{
		repo: &repository.RequestRepository{},
	}
}

func (r *CompanyService)GetCompanies(page int, query string) ([]model.Company, error) {
	return r.repo.GetCompanies(page, query)
}
