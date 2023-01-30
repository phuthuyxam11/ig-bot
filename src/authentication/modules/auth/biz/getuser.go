package userbiz

import (
	"context"
	"github.com/phuthuyxam11/go-common-service/common"
	usermod "igbot.com/authentication/modules/auth/model"
)

type GetUserStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, modeKeys ...string) (*usermod.UserLoginData, error)
}
type GetUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *GetUserBiz {
	return &GetUserBiz{store: store}
}

func (biz *GetUserBiz) GetUserBiz(ctx context.Context, id int) (*usermod.UserLoginData, error) {
	data, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(usermod.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(usermod.EntityName, err)
	}

	if data.DeletedAt.Valid {
		return nil, common.ErrCannotDeleteEntity(usermod.EntityName, nil)
	}

	return data, err
}
