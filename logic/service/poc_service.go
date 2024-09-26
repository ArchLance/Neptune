package service

import (
	"github.com/go-playground/validator/v10"
	"neptune/logic/model"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
	mytime "neptune/utils/time"
)

type PocService struct {
	PocRepository *repository.PocRepository
	Validate      *validator.Validate
}

func NewPocService(repository *repository.PocRepository, validate *validator.Validate) *PocService {
	return &PocService{
		PocRepository: repository,
		Validate:      validate,
	}
}

type PocRequest struct {
	Id                int    `json:"id"`
	VulnerabilityName string `json:"vulnerability_name"`
	PocName           string `validate:"required,max=64,min=1" json:"poc_name"`
	AppName           string `validate:"required,max=64,min=1" json:"app_name"`
	VulnerabilityType int    `validate:"required" json:"vulnerability_type"`
	PocContent        string `validate:"required" json:"poc_content"`
}

func (p PocRequest) toModel() *model.Poc {
	return &model.Poc{
		Id:                p.Id,
		VulnerabilityName: p.VulnerabilityName,
		PocName:           p.PocName,
		AppName:           p.AppName,
		VulnerabilityType: p.VulnerabilityType,
		AddTime:           mytime.Now(),
		PocContent:        p.PocContent,
	}
}

func (s *PocService) Create(poc *PocRequest) error {
	err := s.Validate.Struct(poc)
	if err != nil {
		return myerrors.ParamErr{Err: err}
	}
	return s.PocRepository.Create(poc.toModel())
}

func (s *PocService) Update(poc *PocRequest) error {
	err := s.Validate.Struct(poc)
	if err != nil {
		return myerrors.ParamErr{Err: err}
	}
	return s.PocRepository.Update(poc.toModel())
}

func (s *PocService) Delete(id int) error {
	return s.PocRepository.Delete(id)
}

func (s *PocService) GetById(id int) (model.Poc, error) {
	return s.PocRepository.GetById(id)
}

func (s *PocService) Filter(filter repository.PocFilter) ([]model.Poc, error) {
	return s.PocRepository.PocFilter(filter)
}

func (s *PocService) Count(filter repository.PocFilter) int64 {
	return s.PocRepository.Count(filter)
}
