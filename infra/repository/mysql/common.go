package mysql

import (
	"go_web/domain/entity"
	"go_web/errors"
	"go_web/infra/repository/mysql/model"
	"strings"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var invalidParam = errors.InternalServerErrFromStr("invalid query pram")

type baseRepo[M model.Model, E entity.Entity] struct {
	db *gorm.DB
}

func (e *baseRepo[M, E]) create(params E) (*E, error) {
	model := M{}
	if err := copy(&model, &params); err != nil {
		return nil, err
	}

	if err := e.db.Create(&model).Error; err != nil {
		return nil, err
	}

	result := E{}
	if err := copier.Copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) update(ID uint32, params E) (*E, error) {
	model := M{}
	if err := e.db.First(&model, ID).Error; err != nil {
		return nil, err
	}

	if err := copyIgnoreEmpty(&model, &params); err != nil {
		return nil, err
	}

	if err := e.db.Save(model).Error; err != nil {
		return nil, err
	}

	result := E{}
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) findByID(ID uint32) (*E, error) {
	model := M{}
	if err := e.db.First(&model, ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := E{}
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) findByIDs(IDs []uint32) ([]E, error) {
	models := []M{}
	if err := e.db.Find(&models, IDs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := []E{}
	if err := copy(&result, &models); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *baseRepo[M, E]) findOneByConds(conds string, params ...interface{}) (*E, error) {
	if strings.Count(conds, "?") != len(params) {
		return nil, invalidParam
	}
	model := M{}
	if err := e.db.First(&model, conds, params).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := E{}
	if err := copy(&result, &model); err != nil {
		return nil, err
	}

	return &result, nil
}

func (e *baseRepo[M, E]) findAllByConds(conds string, params ...interface{}) ([]E, error) {
	if strings.Count(conds, "?") != len(params) {
		return nil, invalidParam
	}

	models := []M{}
	if err := e.db.Find(&models, conds, params).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.NotFound(err)
		}
		return nil, err
	}

	result := []E{}
	if err := copy(&result, &models); err != nil {
		return nil, err
	}

	return result, nil
}

func (e *baseRepo[M, E]) countByConds(conds string, params ...interface{}) (int64, error) {
	if strings.Count(conds, "?") != len(params) {
		return 0, invalidParam
	}
	var count int64

	if err := e.db.Model(&M{}).Where(conds, params...).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (e *baseRepo[M, E]) deleteByID(ID uint32) error {
	return e.db.Where("ID = ?", ID).Delete(&M{}).Error
}

func (e *baseRepo[M, E]) deleteByIDs(IDs []uint32) (int64, error) {
	result := e.db.Where("ID IN ?", IDs).Delete(&M{})
	return result.RowsAffected, result.Error
}

func copy(toValue interface{}, fromValue interface{}) (err error) {
	return copier.Copy(toValue, fromValue)
}

func copyIgnoreEmpty(toValue interface{}, fromValue interface{}) (err error) {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
	})
}
