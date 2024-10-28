package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// writeLog - записывает лог в базу даных
func (s *service) writeLog(ctx context.Context, actionType string) error {
	err := s.authRepository.CreateLog(ctx, &model.Log{
		ActionType:    actionType,
		ActionDetails: details(ctx),
	})

	if err != nil {
		return err
	}

	return nil
}
