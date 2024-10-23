package convertor

import (
	"github.com/erikqwerty/auth/internal/model"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// ToModelUserFromCreateRequest - Конвертер для преобразования gRPC-запроса создания пользователя в модель бизнес-логики User
func ToModelUserFromCreateRequest(req *desc.CreateRequest) *model.CreateUser {
	return &model.CreateUser{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: req.Password,
		RoleID:       int32(req.Role),
	}
}

// ToModelUserFromUpdateRequest - Конвертор для преобразования gRPC-запроса обновления пользователя в модель бизнес-логики User
func ToModelUserFromUpdateRequest(req *desc.UpdateRequest) *model.UpdateUser {
	return &model.UpdateUser{
		Email:  req.Email,
		Name:   req.Name.Value,
		RoleID: int32(req.Role),
	}
}

// ToGetResponseFromModelUser - Конвертор для преобразования модели бизнес-логики User в gRPC-ответ
func ToGetResponseFromModelUser(user *model.ReadUser) *desc.GetResponse {
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(user.RoleID),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
