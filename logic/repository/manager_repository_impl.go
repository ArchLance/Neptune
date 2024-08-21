package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"student_manage/logic/model"
	myerrors "student_manage/utils/errors"
)

type ManagerRepositoryImpl struct {
	Db *gorm.DB
}

func NewManagerRepositoryImpl(Db *gorm.DB) *ManagerRepositoryImpl {
	return &ManagerRepositoryImpl{Db: Db}
}
func (r *ManagerRepositoryImpl) Create(manager model.Manager) error {
	result := r.Db.Create(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 创建管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepositoryImpl) Update(manager model.Manager) error {
	result := r.Db.Updates(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 更新管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepositoryImpl) Delete(id int) error {
	var manager model.Manager
	result := r.Db.Where("id = ?", id).Delete(&manager)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 删除管理员失败 -> %w", result.Error)}
	}
	return nil
}
func (r *ManagerRepositoryImpl) GetById(id int) (model.Manager, error) {
	var findManager model.Manager
	result := r.Db.First(&findManager, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return findManager, myerrors.DbErr{Err: fmt.Errorf("repository: 查找管理员id %d失败 -> %w", id, result.Error)}
	}
	return findManager, nil
}
func (r *ManagerRepositoryImpl) GetAll() ([]model.Manager, error) {
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
func (r *ManagerRepositoryImpl) ExistById(id int) (bool, error) {
	var manager model.Manager
	result := r.Db.First(&manager, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, nil
}
func (r *ManagerRepositoryImpl) ExistByAccount(account string) (bool, error) {
	var manager model.Manager
	result := r.Db.First(&manager, "account = ?", account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, nil
}
