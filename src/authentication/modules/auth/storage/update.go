package userstorage

import (
	"context"
	"github.com/phuthuyxam11/go-common-service/common"
	usermod "igbot.com/authentication/modules/auth/model"
)

func (s *sqlStore) UpdateUser(ctx context.Context, id int, data *usermod.UsersModelUpdate) error {
	db := s.db.Begin()

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
