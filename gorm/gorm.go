package gorm

import (
	"github.com/jinzhu/gorm"
)

type gormDB struct {
	*gorm.DB
}

func (db *gormDB) Paginate(page int, limit int) *gorm.DB {

	offset := (page - 1) * limit

	return db.DB.Offset(offset).Limit(limit)
}
