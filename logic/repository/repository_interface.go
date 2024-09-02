package repository

import (
	"neptune/logic/model"
)

type ManagerRepository interface {
	Create(manager model.Manager) error
	Update(manager model.Manager) error
	Delete(id int) error
	GetById(id int) (model.Manager, error)
	GetAll() ([]model.Manager, error)
	ExistById(id int) (bool, error)
	ExistByAccount(account string) (bool, error)
}

type UserRepository interface {
	Update(user *model.User) error
	GetById(id int) (model.User, error)
	GetByAccount(account string) (model.User, error)
	GetByEmail(email string) (model.User, error)
}
