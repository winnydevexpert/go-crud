package repository

import (
	"go-crud-learn/domain"
	"go-crud-learn/model"

	"gorm.io/gorm"
)

type appDataRepository struct {
	DB *gorm.DB
}

func NewAppDataRepository(db *gorm.DB) domain.AppDataRepository {
	return &appDataRepository{DB: db}
}

func (a appDataRepository) GetBeginTransaction() *gorm.DB {
	return a.DB.Begin()
}

func (q appDataRepository) GetListAppData() ([]model.AppData, error) {
	var appDatas []model.AppData
	return appDatas, nil
}
