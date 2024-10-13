package server

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/erikqwerty/auth/internal/db"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
	"golang.org/x/crypto/bcrypt"
)

// Create создание нового пользователя в системе
func (a *Auth) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {

	err := a.checkReqCreate(ctx, req)
	if err != nil {
		return &desc.CreateResponse{Id: 0}, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	location := time.FixedZone("UTC+3", 3*60*60)
	user := db.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		RoleID:       int32(req.Role),
		CreatedAt:    time.Now().In(location),
		UpdatedAt:    time.Now().In(location),
	}

	id, err := a.DB.InsertUser(ctx, user)
	if err != nil {
		return nil, err
	}

	log.Printf("Создание нового пользователя в системе: %v, %v, %v, %v, %v", req.Name, req.Email, req.Password, req.PasswordConfirm, req.Role)
	return &desc.CreateResponse{
		Id: id,
	}, nil
}

// checkReqCreate - проверяет полученный запрос клиента на возможность выполнения
func (a *Auth) checkReqCreate(ctx context.Context, req *desc.CreateRequest) error {
	if req.Password != req.PasswordConfirm {
		return fmt.Errorf("пароль не совпадает с подтверждением пароля, создание пользователя не возможно")
	}
	status, err := a.DB.СheckEmailExists(ctx, req.Email)
	if status {
		return fmt.Errorf("пользователь с таким email уже существует")
	}
	if err != nil {
		// TODO: когда будет логер надо будет отловить, ошибка в работе с бдшкой но не с запросом пользователя
		log.Println(err)
	}
	if !isValidEmail(req.Email) {
		return fmt.Errorf("email не валиден")
	}
	if req.Role == 0 {
		return fmt.Errorf("не указана роль пользователя")
	}

	return nil
}

// isValidEmail проверяет валидность email-адреса. Возвращает true если валидно.
func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
