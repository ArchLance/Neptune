package controller

import "neptune/logic/service"

type PocController struct {
	PocService *service.PocService
}

func NewPocController(service *service.PocService) *PocController {
	return &PocController{
		PocService: service,
	}
}
