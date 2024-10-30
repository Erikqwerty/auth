package convertor

import (
	"time"

	"github.com/erikqwerty/auth/internal/model"
	modelCache "github.com/erikqwerty/auth/internal/repository/cache/model"
)

func ToUserCacheModelFromServiceUserCache(user *model.UserCache) *modelCache.UserCache {
	if user == nil {
		return nil
	}

	var updateAt int64

	if user.UpdatedAt != nil {
		updateAt = user.UpdatedAt.Unix()
	}

	return &modelCache.UserCache{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		RoleID:    user.RoleID,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: &updateAt,
	}
}

func ToServiceUserCacheFromUserCacheModel(user *modelCache.UserCache) *model.UserCache {
	if user == nil {
		return nil
	}

	var updateAt time.Time

	if user.UpdatedAt != nil {
		updateAt = time.Unix(0, *user.UpdatedAt)
	}

	return &model.UserCache{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		RoleID:    user.RoleID,
		CreatedAt: time.Unix(user.CreatedAt, 0),
		UpdatedAt: &updateAt,
	}
}
