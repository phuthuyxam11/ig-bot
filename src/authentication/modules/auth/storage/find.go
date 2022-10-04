package userstorage

import (
	"context"
	"github.com/phuthuyxam11/go-common-service/common"
	"gorm.io/gorm"
	usermod "igbot.com/authentication/modules/auth/model"
)

func (s *sqlStore) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermod.UserLoginData, error) {
	db := s.db.Table(usermod.UserLoginData{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermod.UserLoginData

	if err := db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
