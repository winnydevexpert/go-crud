package domain

import (
	"go-crud-learn/model"

	"gorm.io/gorm"
)

type AppDataUsecase interface {
	GetListAppData() (res []model.AppData, err error)
}

type AppDataRepository interface {
	GetBeginTransaction() *gorm.DB
	GetListAppData() (res []model.AppData, err error)
}
