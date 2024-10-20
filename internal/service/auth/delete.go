package authservice

import "context"

// Delete - удалить пользователя
func (s *service) Delete(ctx context.Context, id int64) error {
	return s.authRepository.DeleteUser(ctx, id)
}
