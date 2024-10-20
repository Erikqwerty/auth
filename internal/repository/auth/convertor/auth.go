package convertor

import (
	"github.com/erikqwerty/auth/internal/model"
	"github.com/erikqwerty/auth/internal/repository/auth/modelrepo"
)

func ToAuthFromRepo(modeldb *modelrepo.User) *model.User {
	return &model.User{
		ID:           modeldb.ID,
		Name:         modeldb.Name,
		Email:        modeldb.Email,
		PasswordHash: modeldb.PasswordHash,
		RoleID:       modeldb.RoleID,
		CreatedAt:    modeldb.CreatedAt,
		UpdatedAt:    modeldb.UpdatedAt,
	}
}
