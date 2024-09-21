package repository

import (
	"gorm.io/gorm"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
)

type PocRepository struct {
	Db *gorm.DB
}

func NewPocRepository(Db *gorm.DB) *PocRepository {
	return &PocRepository{Db: Db}
}

func (r *PocRepository) Create(poc *model.Poc) error {
	result := r.Db.Create(poc)
	if result.Error != nil {
		return myerrors.DbErr{Err: result.Error}
	}
	return nil
}

func (r *PocRepository) Update(poc *model.Poc) error {
	result := r.Db.Updates(poc)
	if result.Error != nil {
		return myerrors.DbErr{Err: result.Error}
	}
	return nil
}

func (r *PocRepository) Delete(id int) error {
	var poc model.Poc
	result := r.Db.Where("id = ?", id).Delete(&poc)
	if result.Error != nil {
		return myerrors.DbErr{Err: result.Error}
	}
	return nil
}

func (r *PocRepository) GetById(id int) (model.Poc, error) {
	var findPoc model.Poc
	result := r.Db.Find(&findPoc, "id=?", id)
	if result.Error != nil {
		return findPoc, myerrors.DbErr{Err: result.Error}
	}
	return findPoc, nil
}

func (r *PocRepository) GetByAppName(appName string) ([]model.Poc, error) {
	var findPocs []model.Poc
	result := r.Db.Find(&findPocs, "app_name=?", appName)
	if result.Error != nil {
		return findPocs, myerrors.DbErr{Err: result.Error}
	}
	return findPocs, nil
}

func (r *PocRepository) GetByVulnerabilityType(vulnerabilityType int) ([]model.Poc, error) {
	var findPocs []model.Poc
	result := r.Db.Find(&findPocs, "vulnerability_type=?", vulnerabilityType)
	if result.Error != nil {
		return findPocs, myerrors.DbErr{Err: result.Error}
	}
	return findPocs, nil
}

func (r *PocRepository) GetAll() ([]model.Poc, error) {
	var pocs []model.Poc
	result := r.Db.Find(&pocs)
	if result.Error != nil {
		return pocs, myerrors.DbErr{Err: result.Error}
	}
	return pocs, nil
}
