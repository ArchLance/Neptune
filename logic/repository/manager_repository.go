package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
)

type ManagerRepository struct {
	Db *gorm.DB
}

func NewManagerRepository(Db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{Db: Db}
}
func (r *ManagerRepository) Create(manager model.Manager) error {
	result := r.Db.Create(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 创建管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepository) Update(manager model.Manager) error {
	result := r.Db.Updates(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 更新管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepository) Delete(id int) error {
	var manager model.Manager
	result := r.Db.Where("id = ?", id).Delete(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 删除管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepository) GetById(id int) (model.Manager, error) {
	var findManager model.Manager
	result := r.Db.First(&findManager, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return findManager, myerrors.DbErr{Err: fmt.Errorf("repository: 查找管理员id %d失败 -> %w", id, result.Error)}
	}
	return findManager, nil
}
func (r *ManagerRepository) GetAll() ([]model.Manager, error) {
	var managers []model.Manager
	result := r.Db.Find(&managers)
	if result.Error != nil {
		return managers, myerrors.DbErr{Err: fmt.Errorf("repository: 查找全部管理员失败 -> %w", result.Error)}
	}
	if len(managers) == 0 {
		return managers, myerrors.DbErr{Err: errors.New("repository: 管理员不存在")}
	}
	return managers, nil
}
func (r *ManagerRepository) ExistById(id int) (bool, error) {
	var manager model.Manager
	result := r.Db.First(&manager, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, nil
}
func (r *ManagerRepository) ExistByAccount(account string) (bool, error) {
	var manager model.Manager
	result := r.Db.First(&manager, "account = ?", account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, nil
}
