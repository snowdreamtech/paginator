package gorm

import (
	"github.com/snowdreamtech/paginator"
	"gorm.io/gorm"
)

type gormDB struct {
	*gorm.DB
}

func (gormDB *gormDB) Paginate(page int, limit int) *gorm.DB {
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

func (gormDB *gormDB) PaginateWithResult(list *[]interface{}, page int, limit int) paginator.PageResult {
	var total int

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
