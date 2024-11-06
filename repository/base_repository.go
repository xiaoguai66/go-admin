package repository

import (
	"admin-demo/global"
	"admin-demo/service/request"
	"gorm.io/gorm"
)

type BaseRepository struct {
	Orm *gorm.DB
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{
		Orm: global.DB,
	}
}

func Paginage(p request.Paginate) func(orm *gorm.DB) *gorm.DB {
	return func(orm *gorm.DB) *gorm.DB {
		return orm.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}
