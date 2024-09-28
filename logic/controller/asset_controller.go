package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/rsp"
	"strconv"
)

type AssetController struct {
	AssetService *service.AssetService
}

func NewAssetController(service *service.AssetService) *AssetController {
	return &AssetController{
		AssetService: service,
	}
}

func (ac *AssetController) GetAssetList(ctx *gin.Context) {
	log.Info("controller: 获取资产列表")
	assets, err := ac.AssetService.GetAll()
	if err != nil {
		rsp.ErrRsp(ctx, err)
	}
	rsp.SuccessRsp(ctx, gin.H{
		"assets": assets,
	})
}

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

func (ac *AssetController) Delete(ctx *gin.Context) {
	log.Info("controller: 删除资产")
	assetsId := ctx.Param("id")
	id, err := strconv.Atoi(assetsId)
	err = ac.AssetService.Delete(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (ac *AssetController) Update(ctx *gin.Context) {
	log.Info("controller: 更新资产")
	updateAssetRequest := service.UpdateAssetRequest{}
	err := ctx.ShouldBindJSON(&updateAssetRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: fmt.Errorf("controller: 更新资产参数错误 -> %w", err)})
		return
	}
	err = ac.AssetService.Update(&updateAssetRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (ac *AssetController) GetById(ctx *gin.Context) {
	log.Info("controller: 获取资产")
	// 获取资产id
	assetsId := ctx.Param("id")
	log.Info(assetsId)
	id, err := strconv.Atoi(assetsId)
	log.Info(id)
	asset, err := ac.AssetService.GetById(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRsp(ctx, asset)
}

func (ac *AssetController) BatchDelete(ctx *gin.Context) {
	log.Info("controller: 获取资产列表")
	idsReq := service.IdsReq{}
	err := ctx.ShouldBindJSON(&idsReq)
	err = ac.AssetService.DeleteByIds(&idsReq)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}
