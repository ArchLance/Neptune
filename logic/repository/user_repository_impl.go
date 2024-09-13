package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{Db: Db}
}
func (r *UserRepositoryImpl) Update(user *model.User) error {
	result := r.Db.Updates(&user)
	if result.Error != nil {
		return myerrors.DbErr{Err: fmt.Errorf("repository: 更新用户失败 -> %w", result.Error)}
	}
	return nil
}
func (r *UserRepositoryImpl) GetById(id uint) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, "id=?", id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByAccount(account string) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, "account=?", account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, result.Error
	}
	return user, nil
}

func (r *UserRepositoryImpl) GetByEmail(email string) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, "email=?", email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, result.Error
	}
	return user, nil
}
