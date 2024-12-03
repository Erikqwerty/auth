package interceptor

import (
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	"github.com/erikqwerty/auth/internal/autherrors"
	desc "github.com/erikqwerty/auth/pkg/userapi_v1"
)

type TestRequest struct {
	Email           string
	Name            string
	Password        string
	PasswordConfirm string
	Role            int32
	ID              int32
}

func TestValidateRequest(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		req     interface{}
		wantErr error
	}{
		{
			name: "api валидный запрос",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Name:            gofakeit.Name(),
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              1},
			wantErr: nil,
		},
		{
			name: "api пропущен email",
			req: TestRequest{
				Name:            gofakeit.Name(),
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              1},
			wantErr: autherrors.ErrEmailNotSpecified,
		},
		{
			name: "api невалидный email",
			req: TestRequest{
				Email:           gofakeit.Name(),
				Name:            gofakeit.Name(),
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              1},
			wantErr: autherrors.ErrInvalidEmail,
		},
		{
			name: "api пропущено имя",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              1},
			wantErr: autherrors.ErrNameNotSpecified,
		},
		{
			name: "api пропущен пароль",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Name:            "Test",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              1},
			wantErr: autherrors.ErrPasswordNotSpecified,
		},
		{
			name: "api пропущено подтверждение пароля",
			req: TestRequest{
				Email:    gofakeit.Email(),
				Name:     "Test",
				Password: "1234",
				Role:     int32(desc.Role_ROLE_USER),
				ID:       1},
			wantErr: autherrors.ErrPasswordConfirmNotSpecified,
		},
		{
			name: "api не указана роль",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Name:            "Test",
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_UNSPECIFIED),
				ID:              1},
			wantErr: autherrors.ErrRoleNotSpecified,
		},
		{
			name: "api не существующая роль",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Name:            "Test",
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            99,
				ID:              1},
			wantErr: autherrors.ErrInvalidRole,
		},
		{
			name: "api ошибочный id",
			req: TestRequest{
				Email:           gofakeit.Email(),
				Name:            "Test",
				Password:        "1234",
				PasswordConfirm: "1234",
				Role:            int32(desc.Role_ROLE_USER),
				ID:              0},
			wantErr: autherrors.ErrInvalidID,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := ValidateRequest(tt.req)
			if tt.wantErr != nil {
				require.Error(t, err)
				require.Equal(t, tt.wantErr, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
