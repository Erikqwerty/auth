package convertor

import (
	"github.com/erikqwerty/auth/internal/model"
	modelRepo "github.com/erikqwerty/auth/internal/repository/auth/model"
)

// ToAuthFromRepo - Конвертор преобразующие структуру repo слоя в структуру бизнес-логики
func ToReadUserFromRepo(modeldb *modelRepo.User) *model.ReadUser {
	createUser := model.CreateUser{
		Name:         modeldb.Name,
		Email:        modeldb.Email,
		PasswordHash: modeldb.PasswordHash,
		RoleID:       modeldb.RoleID,
		CreatedAt:    modeldb.CreatedAt,
	}
	return &model.ReadUser{
		ID:         modeldb.ID,
		CreateUser: createUser,
		UpdatedAt:  modeldb.UpdatedAt,
	}
}
