package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"neptune/logic/model"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
)

type ManagerService struct {
	ManagerRepository *repository.ManagerRepository
	Validate          *validator.Validate
}

func NewManagerService(repository *repository.ManagerRepository, validate *validator.Validate) *ManagerService {
	return &ManagerService{
		ManagerRepository: repository,
		Validate:          validate,
	}
}

type CreateManagerRequest struct {
	Level    int    `json:"level"`
	Name     string `validate:"required,max=255,min=1" json:"name"`
	Account  string `validate:"required,max=255,min=1" json:"account"`
	Password string `validate:"required,max=255,min=1" json:"password"`
}
type UpdateManagerRequest struct {
	Id       int    `json:"id"`
	Level    int    `json:"level"`
	Name     string `validate:"required,max=255,min=1" json:"name"`
	Account  string `validate:"required,max=255,min=1" json:"account"`
	Password string `validate:"required,max=255,min=1" json:"password"`
}
type ManagerResponse struct {
	Id       int
	Level    int    `json:"level"`
	Name     string `json:"name"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (r *ManagerService) Create(manager CreateManagerRequest) error {
	err := r.Validate.Struct(manager)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 创建管理员参数校验失败 -> %w", err)}
	}
	exist, err := r.ManagerRepository.ExistByAccount(manager.Account)
	if exist {
		return myerrors.ExistErr{Err: fmt.Errorf("service: 管理员账号已经存在 -> %w", err)}
	}
	managerModel := model.Manager{
		Level:    manager.Level,
		Name:     manager.Name,
		Account:  manager.Account,
		Password: manager.Password,
	}
	err = r.ManagerRepository.Create(&managerModel)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (r *ManagerService) Update(manager UpdateManagerRequest) error {
	err := r.Validate.Struct(manager)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 更新管理员参数校验失败 -> %w", err)}
	}
	managerData, err := r.ManagerRepository.GetById(manager.Id)
	if err != nil {
		return myerrors.NotFoundErr{Err: err}
	}
	managerData.Level = manager.Level
	managerData.Name = manager.Name
	managerData.Account = manager.Account
	managerData.Password = manager.Password
	err = r.ManagerRepository.Update(&managerData)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (r *ManagerService) Delete(id int) error {
	err := r.ManagerRepository.Delete(id)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (r *ManagerService) GetById(id int) (ManagerResponse, error) {
	managerData, err := r.ManagerRepository.GetById(id)
	if err != nil {
		return ManagerResponse{}, myerrors.NotFoundErr{Err: err}
	}
	managerResponse := ManagerResponse{
		Level:    managerData.Level,
		Name:     managerData.Name,
		Account:  managerData.Account,
		Password: managerData.Password,
	}
	return managerResponse, nil
}

func (r *ManagerService) GetAll() ([]ManagerResponse, error) {
	result, err := r.ManagerRepository.GetAll()
	if err != nil {
		return []ManagerResponse{}, myerrors.NotFoundErr{Err: err}
	}
	var managers []ManagerResponse
	for _, manager := range result {
		manager := ManagerResponse{
			Id:       manager.Id,
			Level:    manager.Level,
			Name:     manager.Name,
			Account:  manager.Account,
			Password: manager.Password,
		}
		managers = append(managers, manager)
	}
	return managers, nil
}
