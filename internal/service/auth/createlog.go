package auth

import (
	"context"

	"github.com/erikqwerty/auth/internal/model"
)

// createLog - записывает лог в базу даных
func (s *service) createLog(ctx context.Context, actionType string) error {
	errTx := s.authRepository.CreateLog(ctx, &model.Log{
		ActionType:      actionType,
		ActionDetails:   details(ctx),
		ActionTimestamp: timeNowUTC3(),
	})

	if errTx != nil {
		return errTx
	}

	return nil
}
