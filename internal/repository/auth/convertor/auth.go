package convertor

import (
	"github.com/erikqwerty/auth/internal/model"
	modelRepo "github.com/erikqwerty/auth/internal/repository/auth/model"
)

// ToAuthFromRepo - Конвертор преобразующие структуру repo слоя в структуру бизнес-логики
func ToAuthFromRepo(modeldb *modelRepo.User) *model.User {
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
