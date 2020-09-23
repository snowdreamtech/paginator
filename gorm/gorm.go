package gorm

import (
	"github.com/jinzhu/gorm"
	"github.com/snowdreamtech/paginator"
)

type gormDB struct {
	*gorm.DB
}

func (gormDB *gormDB) Paginate(page int, limit int) *gorm.DB {
	db := gormDB.DB

	offset := (page - 1) * limit

	return db.Offset(offset).Limit(limit)
}

func (gormDB *gormDB) PaginateWithResult(query paginator.Query, options paginator.Options) Result {
	db := gormDB.
	
	return db.Offset(offset).Limit(limit)
}
