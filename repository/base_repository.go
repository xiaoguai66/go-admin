package repository

import (
	"admin-demo/global"
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
