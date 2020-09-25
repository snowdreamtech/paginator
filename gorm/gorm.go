package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/snowdreamtech/paginator"
)

type gormDB struct {
	*gorm.DB
}

func (gormDB *gormDB) Paginate(page uint, limit uint) *gorm.DB {
	db := gormDB.DB

	if page <= 0 {
		page = paginator.DefaultPage
	}

	if limit <= 0 {
		limit = paginator.DefaultLimit
	}

	offset := (page - 1) * limit

	return db.Offset(offset).Limit(limit)
}

func (gormDB *gormDB) PaginateWithResult(list *[]interface{}, page uint, limit uint) paginator.PageResult {
	var total uint

	db := gormDB.DB

	if page <= 0 {
		page = paginator.DefaultPage
	}

	if limit <= 0 {
		limit = paginator.DefaultLimit
	}

	offset := (page - 1) * limit

	err := db.Find(&list).Count(&total).Error

	if err != nil {
		return paginator.PageResult{
			Page:  page,
			Limit: limit,
			Total: total,
			List:  *list,
			Error: err,
		}
	}

	err = db.Offset(offset).Limit(limit).Find(list).Error

	return paginator.PageResult{
		Page:  page,
		Limit: limit,
		Total: total,
		List:  *list,
		Error: err,
	}
}
