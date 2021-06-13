package repository

import (
	"securities_admin/db"
	"securities_admin/internal/model"
)

const PerPage = 15

type IRequestRepository interface {
	GetRequests(intent int, page int, senderId string) ([]model.Request, error)
}

type RequestRepository struct {
}

func (r *RequestRepository) GetRequests(requestType int, page int, senderId string) ([]model.Request, error) {
	var result []model.Request
	tx := db.DB
	if requestType == model.RequestTypeUnClassified {
		tx = tx.Where(map[string]interface{}{"Intent": ""})
	} else if requestType == model.RequestTypeClassified {
		tx = tx.Not(map[string]interface{}{"Intent": ""})
	}

	if senderId != "" {
		tx = tx.Where(map[string]interface{}{"SenderID": senderId})
	}

	res := tx.Limit(PerPage).Offset((page - 1) * PerPage).Find(&result)
	return result, res.Error
}
