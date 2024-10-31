package convertor

import (
	"strconv"
	"time"

	"github.com/erikqwerty/auth/internal/model"
	modelCache "github.com/erikqwerty/auth/internal/repository/cache/model"
)

func ToUserCacheModelFromServiceUserCache(user *model.UserInfo) *modelCache.UserCache {
	if user == nil {
		return nil
	}

	var updateAt int64

	if user.UpdatedAt != nil {
		updateAt = user.UpdatedAt.Unix()
	}

	return &modelCache.UserCache{
		ID:           strconv.Itoa(int(user.ID)),
		Email:        user.Email,
		Name:         user.Name,
		PasswordHash: user.PasswordHash,
		RoleID:       user.RoleID,
		CreatedAt:    user.CreatedAt.Unix(),
		UpdatedAt:    &updateAt,
	}
}

func ToServiceUserCacheFromUserCacheModel(user *modelCache.UserCache) *model.UserInfo {
	if user == nil {
		return nil
	}

	var updateAt time.Time

	if user.UpdatedAt != nil {
		updateAt = time.Unix(0, *user.UpdatedAt)
	}
	id, err := strconv.Atoi(user.ID)
	if err != nil {
		return nil
	}

	return &model.UserInfo{
		ID: int64(id),
		CreateUser: model.CreateUser{
			Email:        user.Email,
			Name:         user.Name,
			PasswordHash: user.PasswordHash,
			RoleID:       user.RoleID,
		},
		CreatedAt: time.Unix(0, user.CreatedAt),
		UpdatedAt: &updateAt,
	}
}
