package model

// UserCache - структура для кеширования пользователя в redis
type UserCache struct {
	ID        string `redis:"id"`
	Email     string `redis:"email"`
	Name      string `redis:"name"`
	RoleID    int32  `redis:"role_id"`
	CreatedAt int64  `redis:"created_at"`
	UpdatedAt *int64 `redis:"updated_at"`
}
