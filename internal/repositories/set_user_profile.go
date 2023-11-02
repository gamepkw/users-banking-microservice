package repository

import (
	"context"
	"fmt"
	"time"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) SetUserProfile(ctx context.Context, u model.UserProfile) (err error) {
	query := `UPDATE banking.users SET first_name=?, last_name=?, email=?, age=?, updated_at=? WHERE uuid=?`

	fmt.Println(u)

	stmt, err := m.conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, u.FirstName, u.LastName, u.Email, u.Age, time.Now(), u.UUID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		return err
	}
	return
}
