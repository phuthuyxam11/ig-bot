package userstorage

import (
	"200lab/struct/common"
	usermod "200lab/struct/modules/auth/model"
	"context"
)

func (s *sqlStore) CreateUser(ctx context.Context, data *usermod.UsersModelCreate) error {
	db := s.db.Begin()

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
