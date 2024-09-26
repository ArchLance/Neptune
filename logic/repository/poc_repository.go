package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"neptune/logic/model"
	myerrors "neptune/utils/errors"
)

type PocRepository struct {
	Db *gorm.DB
}

func NewPocRepository(Db *gorm.DB) *PocRepository {
	return &PocRepository{Db: Db}
}

func (r *PocRepository) Create(poc *model.Poc) error {
	if poc == nil || poc.Id != 0 {
		return myerrors.RequestErr{Err: errors.New("请求参数错误")}
	}
	if err := r.check(*poc); err != nil {
		return err
	}
	return r.Db.Create(poc).Error
}

func (r *PocRepository) Update(poc *model.Poc) error {
	if poc == nil || poc.Id == 0 {
		return myerrors.RequestErr{Err: errors.New("请求参数错误")}
	}
	if err := r.check(*poc); err != nil {
		return err
	}
	result := r.Db.Where("id = ?", poc.Id).Updates(poc)
	if result.RowsAffected == 0 {
		return myerrors.RequestErr{Err: errors.New("更新失败，请检查参数")}
	}
	return result.Error
}

func (r *PocRepository) Delete(id int) error {
	if id == 0 {
		return myerrors.RequestErr{Err: errors.New("请求参数错误")}
	}
	return r.Db.Where("id = ?", id).Delete(&model.Poc{}).Error
}

func (r *PocRepository) GetById(id int) (model.Poc, error) {
	var findPoc model.Poc
	if id == 0 {
		return findPoc, myerrors.RequestErr{Err: errors.New("请求参数错误")}
	}
	result := r.Db.Find(&findPoc, "id=?", id)
	if result.Error != nil {
		return findPoc, myerrors.DbErr{Err: result.Error}
	}
	return findPoc, nil
}

type PocFilter struct {
	Offset            int    `form:"offset"`
	Limit             int    `form:"limit"`
	AppName           string `form:"app_name"`
	VulnerabilityName string `form:"vulnerability_name"`
	VulnerabilityType []int  `form:"vulnerability_type[]"`
}

// PocFilter 根据UserFilter过滤用户
func (r *PocRepository) PocFilter(filter PocFilter) ([]model.Poc, error) {
	var pocs []model.Poc
	if err := r.setupFilterSession(filter).Find(&pocs).Error; err != nil {
		return nil, myerrors.DbErr{Err: err}
	}
	return pocs, nil
}

func (r *PocRepository) setupFilterSession(filter PocFilter) *gorm.DB {
	db := r.Db.Model(&model.Poc{})
	if filter.Offset > 0 {
		db = db.Offset(filter.Offset)
	}
	if filter.Limit > 0 {
		db = db.Limit(filter.Limit)
	}
	if len(filter.AppName) > 0 {
		db = db.Where("app_name LIKE ?", "%"+filter.AppName+"%")
	}
	if len(filter.VulnerabilityType) > 0 {
		db = db.Where("vulnerability_type IN ?", filter.VulnerabilityType)
	}
	if len(filter.VulnerabilityName) > 0 {
		db = db.Where("vulnerability_name = ?", filter.VulnerabilityName)
	}
	// 若没有排序要求，则按ID排序
	db = db.Order("id")
	return db
}

func (r *PocRepository) Count(filter PocFilter) int64 {
	filter.Offset = 0
	filter.Limit = 0
	var count int64
	if err := r.setupFilterSession(filter).Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (r *PocRepository) check(poc model.Poc) error {
	if poc.VulnerabilityType <= model.VulnerabilityTypeMin || poc.VulnerabilityType >= model.VulnerabilityTypeMax {
		return myerrors.RequestErr{Err: errors.New("请求参数错误")}
	}
	if len(poc.PocName) > 0 {
		var count int64
		r.Db.Model(&model.Poc{}).Where("id <> ? AND poc_name = ?", poc.Id, poc.PocName).Count(&count)
		if count > 0 { // 用户名已存在
			return myerrors.ExistErr{Err: errors.New("当前poc名称已经存在")}
		}
	}
	return nil
}
