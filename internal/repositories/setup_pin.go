package repository

import (
	"context"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) SetUpPin(ctx context.Context, u *model.Pin) (err error) {
	query := `UPDATE banking.users SET hashed_pin=? WHERE uuid=?`

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.Pin, u.Tel)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {

		return model.ErrSetPin
	}

	return
}
