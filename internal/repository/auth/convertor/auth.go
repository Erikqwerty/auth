package convertor

import (
	"github.com/erikqwerty/auth/internal/model"
	modelRepo "github.com/erikqwerty/auth/internal/repository/auth/model"
)

// ToUserInfoFromRepo - Конвертор преобразующие структуру repo слоя в структуру бизнес-логики
func ToUserInfoFromRepo(modeldb *modelRepo.User) *model.UserInfo {
	createUser := model.CreateUser{
		Name:         modeldb.Name,
		Email:        modeldb.Email,
		PasswordHash: modeldb.PasswordHash,
		RoleID:       modeldb.RoleID,
	}
	return &model.UserInfo{
		ID:         modeldb.ID,
		CreateUser: createUser,
		UpdatedAt:  modeldb.UpdatedAt,
	}
}
