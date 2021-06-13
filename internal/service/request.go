package service

import (
	"securities_admin/internal/model"
	"securities_admin/internal/repository"
)

type IRequestService interface {
	GetRequests(requestType int, page int, senderId string) ([]model.Request, error)
}

type RequestService struct {
	repo repository.IRequestRepository
}

func NewRequestService() *RequestService{
	return &RequestService{
		repo: &repository.RequestRepository{},
	}
}

func (r *RequestService)GetRequests(requestType int, page int, senderId string) ([]model.Request, error) {
	return r.repo.GetRequests(requestType, page, senderId)
}
