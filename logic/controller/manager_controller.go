package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
	"student_manage/logic/data/request"
	"student_manage/logic/service"
	myerrors "student_manage/utils/errors"
	"student_manage/utils/rsp"
)

// 这里以ManagerService接口当做参数，
type ManagerController struct {
	ManagerService service.ManagerService
}

func NewManagerController(service service.ManagerService) *ManagerController {
	return &ManagerController{
		ManagerService: service,
	}
}

func (controller *ManagerController) Create(ctx *gin.Context) {
	log.Info().Msg("Create manager")
	createManagerRequest := request.CreateManagerRequest{}
	err := ctx.ShouldBind(&createManagerRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 获取创建管理员参数失败 -> %w", err)})
		return
	}

	err = controller.ManagerService.Create(createManagerRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (controller *ManagerController) Update(ctx *gin.Context) {
	log.Info().Msg("Update manager")
	updateManagerRequest := request.UpdateManagerRequest{}
	err := ctx.ShouldBind(&updateManagerRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 获取更新管理员参数失败 -> %w", err)})
		return
	}

	err = controller.ManagerService.Update(updateManagerRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}
func (controller *ManagerController) Delete(ctx *gin.Context) {
	log.Info().Msg("Delete manager")
	managerId := ctx.Param("id")
	id, err := strconv.Atoi(managerId)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 删除时转换id %s 失败 -> %w", managerId, err)})
		return
	}

	err = controller.ManagerService.Delete(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}
func (controller *ManagerController) GetById(ctx *gin.Context) {
	log.Info().Msg("Get manager")
	managerId := ctx.Param("id")
	id, err := strconv.Atoi(managerId)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 查找时转换id %s 失败 -> %w", managerId, err)})
		return
	}
	manager, err := controller.ManagerService.GetById(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRsp(ctx, manager)
}
func (controller *ManagerController) GetAll(ctx *gin.Context) {
	log.Info().Msg("Get all manager")
	managers, err := controller.ManagerService.GetAll()
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRsp(ctx, managers)
}
