package usecase

import (
	"go-crud-learn/domain"
	"go-crud-learn/model"
)

type appDataUsecase struct {
	appDataRepo domain.AppDataRepository
}

func NewAppDataUsecase(appDataRepo domain.AppDataRepository) domain.AppDataUsecase {
	return &appDataUsecase{
		appDataRepo: appDataRepo,
	}
}

func (u appDataUsecase) GetListAppData() (res []model.AppData, err error) {
	res, err = u.appDataRepo.GetListAppData()
	if err != nil {
		return res, err
	}
	return res, nil
}
