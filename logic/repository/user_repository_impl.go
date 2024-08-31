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

func (r *UserRepositoryImpl) ExistById(id int) (bool, error) {
	var user model.User
	result := r.Db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, result.Error
	}
	return true, nil
}

func (r *UserRepositoryImpl) GetById(id int) (model.User, error) {
	var user model.User
	result := r.Db.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, myerrors.DbErr{Err: fmt.Errorf("repository: 查找用户id %d失败 -> %w", id, result.Error)}
	}
	return user, nil
}
