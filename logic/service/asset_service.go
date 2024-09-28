package service

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"neptune/logic/model"
	"neptune/logic/repository"
	myerrors "neptune/utils/errors"
	"net"
	"strings"
)

type AssetService struct {
	AssetRepository *repository.AssetRepository
	Validate        *validator.Validate
}

func NewAssetService(repository *repository.AssetRepository, validate *validator.Validate) *AssetService {
	return &AssetService{
		AssetRepository: repository,
		Validate:        validate,
	}
}

type UpdateAssetRequest struct {
	AssetId     int    `json:"asset_id" validate:"required"`
	AssetName   string `json:"asset_name" validate:"required"`
	ProductName string `json:"product_name" validate:"required"`
	IpList      string `json:"ip_list" validate:"required"`
}

type CreateAssetRequest struct {
	AssetName   string `json:"asset_name" validate:"required"`
	ProductName string `json:"product_name" validate:"required"`
	IpList      string `json:"ip_list" validate:"required"`
}

type AssetResponse struct {
	AssetId     int    `json:"asset_id" validate:"required"`
	AssetName   string `json:"asset_name" validate:"required"`
	ProductName string `json:"product_name" validate:"required"`
	IpList      string `json:"ip_list" validate:"required"`
	IpNumber    int    `json:"ip_number" validate:"required"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

func CalculateIpNumber(ipList string) int {
	return len(strings.Split(ipList, ","))
}

// FilterIpList TODO: 处理IPList
// 两条规则：1. 将输入的ipList中的分隔符进行统一化处理，统一处理为英文逗号，可接受的分隔符为换行符和英文逗号
// 2. 去掉ipList中的重复资产，分为两步，首先将所有域名资产转换为ip地址 ,然后为没有scheme的ip添加http://和https://，最后去除重复行
func FilterIpList(ipList string) string {
	// Step 1: Normalize separators
	normalized := strings.ReplaceAll(ipList, "\n", ",")
	items := strings.Split(normalized, ",")
	// Step 2: Convert domains to IPs and add schemes
	uniqueItems := make(map[string]bool)
	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		// Convert domain to IP if possible
		if ip := net.ParseIP(item); ip == nil {
			ips, err := net.LookupIP(item)
			if err == nil && len(ips) > 0 {
				item = ips[0].String()
			}
		}
		// Add schemes if missing
		if !strings.Contains(item, "://") {
			uniqueItems["http://"+item] = true
			uniqueItems["https://"+item] = true
		} else {
			uniqueItems[item] = true
		}
	}
	// Convert unique items back to a comma-separated string
	result := []string{}
	for item := range uniqueItems {
		result = append(result, item)
	}
	return strings.Join(result, ",")
}

func (a CreateAssetRequest) toModel() *model.Asset {
	newIpList := FilterIpList(a.IpList)
	return &model.Asset{
		AssetName:   a.AssetName,
		ProductName: a.ProductName,
		IpList:      newIpList,
		IpNumber:    CalculateIpNumber(newIpList),
	}
}

func (a UpdateAssetRequest) toModel() *model.Asset {
	return &model.Asset{
		AssetName:   a.AssetName,
		ProductName: a.ProductName,
		IpList:      a.IpList,
		IpNumber:    CalculateIpNumber(a.IpList),
	}
}

func (r *AssetService) GetById(id int) (AssetResponse, error) {
	assetData, err := r.AssetRepository.GetById(id)
	if err != nil {
		return AssetResponse{}, myerrors.NotFoundErr{Err: err}
	}
	assetResponse := AssetResponse{
		AssetId:     assetData.AssetId,
		AssetName:   assetData.AssetName,
		ProductName: assetData.ProductName,
		IpList:      assetData.IpList,
		IpNumber:    assetData.IpNumber,
	}
	return assetResponse, nil

}

func (r *AssetService) Create(assetRequest *CreateAssetRequest) error {
	err := r.Validate.Struct(assetRequest)
	if err != nil {
		return myerrors.ParamErr{Err: fmt.Errorf("service: 创建资产参数校验失败 -> %w", err)}
	}
	err = r.AssetRepository.Create(assetRequest.toModel())
	if err != nil {
		return myerrors.DbErr{Err: fmt.Errorf("service: 创建资产失败 -> %w", err)}
	}
	return nil
}

func (r *AssetService) Delete(id int) error {
	return r.AssetRepository.Delete(id)
}

func (r *AssetService) Update(assetRequest *UpdateAssetRequest) error {
	return r.AssetRepository.Update(assetRequest.toModel())

}

func (r *AssetService) GetAll() ([]model.Asset, error) {
	return r.AssetRepository.GetAll()
}

func (r *AssetService) DeleteByIds(ids *IdsReq) error {
	return r.AssetRepository.DeleteByIds(ids.Ids)
}
