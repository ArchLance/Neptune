package repository

import (
	"student_manage/logic/model"
)

type ManagerRepository interface {
	Create(manager model.Manager) error
	Update(manager model.Manager) error
	Delete(id int) error
	GetById(id int) (model.Manager, error)
	GetAll() ([]model.Manager, error)
}
