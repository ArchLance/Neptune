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

type CreatePocRequest struct {
	VulnerabilityName string `json:"vulnerability_name"`
	PocName           string `validate:"required,max=64,min=1" json:"poc_name"`
	AppName           string `validate:"required,max=64,min=1" json:"app_name"`
	VulnerabilityType int    `validate:"required" json:"vulnerability_type"`
	PocContent        string `validate:"required" json:"poc_content"`
}

type UpdatePocRequest struct {
	Id                int    `json:"id"`
	VulnerabilityName string `json:"vulnerability_name"`
	PocName           string `validate:"required,max=64,min=1" json:"poc_name"`
	AppName           string `validate:"required,max=64,min=1" json:"app_name"`
	VulnerabilityType int    `validate:"required" json:"vulnerability_type"`
	PocContent        string `validate:"required" json:"poc_content"`
}

type PocResponse struct {
	Id                int    `json:"id"`
	VulnerabilityName string `json:"vulnerability_name"`
	PocName           string `json:"poc_name"`
	AppName           string `json:"app_name"`
	VulnerabilityType int    `json:"vulnerability_type"`
	AddTime           string `json:"add_time"`
	PocContent        string `json:"poc_content"`
}

func (s *PocService) Create(poc *CreatePocRequest) error {
	err := s.Validate.Struct(poc)
	if err != nil {
		return myerrors.ParamErr{Err: err}
	}
	pocModel := model.Poc{
		VulnerabilityName: poc.VulnerabilityName,
		PocName:           poc.PocName,
		AppName:           poc.AppName,
		VulnerabilityType: poc.VulnerabilityType,
		AddTime:           mytime.Now(),
		PocContent:        poc.PocContent,
	}
	err = s.PocRepository.Create(&pocModel)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (s *PocService) Update(poc *UpdatePocRequest) error {
	err := s.Validate.Struct(poc)
	if err != nil {
		return myerrors.ParamErr{Err: err}
	}
	pocData, err := s.PocRepository.GetById(poc.Id)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	pocData.VulnerabilityName = poc.VulnerabilityName
	pocData.PocName = poc.PocName
	pocData.AppName = poc.AppName
	pocData.VulnerabilityType = poc.VulnerabilityType
	pocData.AddTime = mytime.Now()
	pocData.PocContent = poc.PocContent
	err = s.PocRepository.Update(&pocData)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (s *PocService) Delete(id int) error {
	err := s.PocRepository.Delete(id)
	if err != nil {
		return myerrors.DbErr{Err: err}
	}
	return nil
}

func (s *PocService) GetById(id int) (PocResponse, error) {
	pocData, err := s.PocRepository.GetById(id)
	if err != nil {
		return PocResponse{}, myerrors.DbErr{Err: err}
	}
	pocResponse := PocResponse{
		Id:                pocData.Id,
		VulnerabilityName: pocData.VulnerabilityName,
		PocName:           pocData.PocName,
		AppName:           pocData.AppName,
		VulnerabilityType: pocData.VulnerabilityType,
		AddTime:           pocData.AddTime,
		PocContent:        pocData.PocContent,
	}
	return pocResponse, nil
}
func (s *PocService) GetByAppName(appName string) ([]PocResponse, error) {
	result, err := s.PocRepository.GetByAppName(appName)
	if err != nil {
		return []PocResponse{}, myerrors.DbErr{Err: err}
	}
	var pocs []PocResponse
	for _, poc := range result {
		poc := PocResponse{
			Id:                poc.Id,
			VulnerabilityName: poc.VulnerabilityName,
			PocName:           poc.PocName,
			AppName:           poc.AppName,
			VulnerabilityType: poc.VulnerabilityType,
			AddTime:           poc.AddTime,
			PocContent:        poc.PocContent,
		}
		pocs = append(pocs, poc)
	}
	return pocs, nil
}

func (s *PocService) GetByVulnerabilityType(vulnerabilityType int) ([]PocResponse, error) {
	result, err := s.PocRepository.GetByVulnerabilityType(vulnerabilityType)
	if err != nil {
		return []PocResponse{}, myerrors.DbErr{Err: err}
	}
	var pocs []PocResponse
	for _, poc := range result {
		poc := PocResponse{
			Id:                poc.Id,
			VulnerabilityName: poc.VulnerabilityName,
			PocName:           poc.PocName,
			AppName:           poc.AppName,
			VulnerabilityType: poc.VulnerabilityType,
			AddTime:           poc.AddTime,
			PocContent:        poc.PocContent,
		}
		pocs = append(pocs, poc)
	}
	return pocs, nil
}

func (s *PocService) GetAll() ([]PocResponse, error) {
	result, err := s.PocRepository.GetAll()
	if err != nil {
		return []PocResponse{}, myerrors.DbErr{Err: err}
	}
	var pocs []PocResponse
	for _, poc := range result {
		poc := PocResponse{
			Id:                poc.Id,
			VulnerabilityName: poc.VulnerabilityName,
			PocName:           poc.PocName,
			AppName:           poc.AppName,
			VulnerabilityType: poc.VulnerabilityType,
			AddTime:           poc.AddTime,
			PocContent:        poc.PocContent,
		}
		pocs = append(pocs, poc)
	}
	return pocs, nil
}
