package repository

import (
	"context"
	"database/sql"
	"log"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

func (m *userRepository) GetHashedPinByUUID(ctx context.Context, uuid string) (res *model.Pin, err error) {
	pinDB, err := m.fetchPinFromDatabase(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return &pinDB, err

}

func (m *userRepository) fetchPinFromDatabase(ctx context.Context, uuid string) (res model.Pin, err error) {
	query := `SELECT hashed_pin FROM banking.users WHERE uuid = ?`

	rows, err := m.conn.Query(query, uuid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	found := false

	var pinValue sql.NullString
	for rows.Next() {
		found = true
		if err := rows.Scan(&pinValue); err != nil {
			log.Fatal(err)
		}
		if pinValue.Valid {
			res.Pin = pinValue.String
		} else {
			res.Pin = ""
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
