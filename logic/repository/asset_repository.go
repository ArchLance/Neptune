package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
)

type AssetRepository struct {
	Db *gorm.DB
}

func NewAssetRepository(Db *gorm.DB) *AssetRepository {
	return &AssetRepository{Db: Db}
}

func (r *AssetRepository) Update(asset *model.Asset) error {
	result := r.Db.Updates(&asset)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 更新资产失败 -> %w", result.Error)}
	}
	return nil
}
func (r *AssetRepository) GetById(id int) (model.Asset, error) {
	var asset model.Asset
	result := r.Db.First(&asset, "asset_id=?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return asset, result.Error
	}
	return asset, nil
}

func (r *AssetRepository) GetAll() ([]model.Asset, error) {
	var assets []model.Asset
	result := r.Db.Find(&assets)
	if result.Error != nil {
		return assets, myerrors.DbErr{Err: fmt.Errorf("repository: 获取所有资产失败 -> %w", result.Error)}
	}
	return assets, nil
}

func (r *AssetRepository) Create(asset *model.Asset) error {
	result := r.Db.Create(&asset)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 创建资产失败 -> %w", result.Error)}
	}
	return nil
}

func (r *AssetRepository) Delete(id int) error {
	result := r.Db.Delete(&model.Asset{}, "asset_id = ?", id)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 删除资产失败 -> %w", result.Error)}
	}
	return nil
}

func (r *AssetRepository) DeleteByIds(ids []int) error {
	result := r.Db.Delete(&model.Asset{}, "asset_id IN (?)", ids)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 批量删除资产失败 -> %w", result.Error)}
	}
	return nil
}
