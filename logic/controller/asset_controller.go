package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/rsp"
)

type AssetController struct {
	AssetService *service.AssetService
}

func NewAssetController(service *service.AssetService) *AssetController {
	return &AssetController{
		AssetService: service,
	}
}

//func (ac *AssetController) GetAssetList() {
//	log.Info("controller: 获取资产列表")
//	err := ac.AssetService.GetAssetList()
//}

func (ac *AssetController) Create(ctx *gin.Context) {
	log.Info("controller: 创建资产")
	createAssetRequest := service.CreateAssetRequest{}
	err := ctx.ShouldBindJSON(&createAssetRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 创建资产参数错误 -> %w", err)})
		return
	}
	err = ac.AssetService.Create(&createAssetRequest)
	if err != nil {
		log.Error("controller: 创建资产失败，", err)
	}
	rsp.SuccessRspWithNoData(ctx)
}
