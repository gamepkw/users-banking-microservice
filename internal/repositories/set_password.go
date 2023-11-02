package repository

import (
	"context"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) SetPassword(ctx context.Context, u *model.UpdatePassword) (res model.UserResponse, err error) {
	query := `UPDATE banking.users SET hashed_password=?, updated_at=? WHERE uuid=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	_, err = stmt.ExecContext(ctx, u.NewPassword, time.Now(), u.Tel)
	if err != nil {
		return
	}
	res.UUID = u.Tel
	res.HashedPassword = u.NewPassword

	return
}
