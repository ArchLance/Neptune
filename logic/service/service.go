package service

import (
	"student_manage/logic/data/request"
	"student_manage/logic/data/response"
)

type ManagerService interface {
	Create(manager request.CreateManagerRequest) error
	Update(manager request.UpdateManagerRequest) error
	Delete(id int) error
	GetById(id int) (response.ManagerResponse, error)
	GetAll() ([]response.ManagerResponse, error)
}
