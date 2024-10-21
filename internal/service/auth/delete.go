package authservice

import "context"

// Delete - удалить пользователя
func (s *service) Delete(ctx context.Context, id int64) error {

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		errTX := s.authRepository.DeleteUser(ctx, id)
		if errTX != nil {
			return errTX
		}
		//
		return nil
	})

	return err
}
