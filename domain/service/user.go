package service

import (
	"go_web/domain/entity"

	"gorm.io/gorm"
)

type GetListOption struct {
	PageIndex   int
	ItemPerPage int
}

type UserService interface {
	WithTx(db *gorm.DB) UserService
	GetByID(ID uint32) (result *entity.User, err error)
	GetByEmail(email string) (result *entity.User, err error)
	DeleteByID(ID uint32) error
	Create(user entity.User) (result *entity.User, err error)
	Update(ID uint32, user entity.User) (result *entity.User, err error)
	ValidateCreate(user entity.User) error
	ValidateUpdate(user entity.User) error
	GetList(GetListOption) ([]*entity.User, int64, error)
}
