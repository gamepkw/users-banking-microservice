package repository

import (
	"context"
	"log"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) GetHashedPasswordByUUID(ctx context.Context, uuid string) (res *model.UserResponse, err error) {

	userResponse, err := m.fetchPasswordFromDatabase(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return &userResponse, err
}

func (m *userRepository) fetchPasswordFromDatabase(ctx context.Context, uuid string) (res model.UserResponse, err error) {
	query := `SELECT hashed_password FROM banking.users WHERE uuid = ?`

	rows, err := m.conn.Query(query, uuid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	found := false

	for rows.Next() {
		found = true
		err := rows.Scan(&res.HashedPassword)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	if !found {
		return res, model.ErrUserNotFound
	}

	return
}
