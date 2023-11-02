package repository

import (
	"context"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) RegisterUser(ctx context.Context, u *model.User) (res model.UserResponse, err error) {
	query := `INSERT INTO banking.users SET uuid=?, hashed_password=?, tel=?, created_at=?`
	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		if isDuplicateEntryError(err) {
			return res, model.ErrDuplicateUUID
		}
	}

	_, err = stmt.ExecContext(ctx, u.UUID, u.Password, u.Tel, time.Now())
	if err != nil {
		return
	}
	res.UUID = u.UUID
	res.HashedPassword = u.Password

	return
}
