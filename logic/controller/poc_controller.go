package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"neptune/logic/repository"
	"neptune/logic/service"
	myerrors "neptune/utils/errors"
	"neptune/utils/rsp"
	"strconv"
)

type PocController struct {
	PocService *service.PocService
}

func NewPocController(service *service.PocService) *PocController {
	return &PocController{
		PocService: service,
	}
}

func validateAndFormatYAML(input string) (string, error) {
	// 解析 YAML
	var data interface{}
	err := yaml.Unmarshal([]byte(input), &data)
	if err != nil {
		return "", fmt.Errorf("无效的 YAML 格式: %v", err)
	}

	// 格式化 YAML
	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2) // 设置缩进为 2 个空格

	err = encoder.Encode(data)
	if err != nil {
		return "", fmt.Errorf("格式化 YAML 时出错: %v", err)
	}

	return buf.String(), nil
}

func (c *PocController) Create(ctx *gin.Context) {
	log.Info("controller: 创建poc")
	createPocRequest := service.PocRequest{}
	err := ctx.ShouldBind(&createPocRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	createPocRequest.PocContent, err = validateAndFormatYAML(createPocRequest.PocContent)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	err = c.PocService.Create(&createPocRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *PocController) Update(ctx *gin.Context) {
	log.Info("controller: 更新poc")
	updatePocRequest := service.PocRequest{}
	err := ctx.ShouldBind(&updatePocRequest)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	updatePocRequest.PocContent, err = validateAndFormatYAML(updatePocRequest.PocContent)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	err = c.PocService.Update(&updatePocRequest)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *PocController) Delete(ctx *gin.Context) {
	log.Info("controller: 删除poc")

	pocId := ctx.Query("id")
	id, err := strconv.Atoi(pocId)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	err = c.PocService.Delete(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRspWithNoData(ctx)
}

func (c *PocController) GetById(ctx *gin.Context) {
	log.Info("controller: poc")
	managerId := ctx.Query("id")
	id, err := strconv.Atoi(managerId)
	if err != nil {
		rsp.ErrRsp(ctx, myerrors.ParamErr{Err: err})
		return
	}
	poc, err := c.PocService.GetById(id)
	if err != nil {
		rsp.ErrRsp(ctx, err)
		return
	}
	rsp.SuccessRsp(ctx, poc)
}

func (c *PocController) PocFilter(ctx *gin.Context) {
	log.Info("controller: poc")
	var filter repository.PocFilter
	if ctx.ShouldBindQuery(&filter) != nil {
		rsp.ErrRsp(ctx, myerrors.RequestErr{Err: errors.New("绑定参数失败")})
		return
	}
	// 筛选
	pocs, err := c.PocService.Filter(filter)
	if err != nil {
		log.Errorf("PocFilter error: %v", err)
		rsp.ErrRsp(ctx, err)
		return
	}
	count := c.PocService.Count(filter)
	rsp.SuccessRsp(ctx, gin.H{
		"pocs":  pocs,
		"count": count,
	})
}
