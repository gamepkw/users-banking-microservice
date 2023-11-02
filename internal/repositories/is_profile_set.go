package repository

import (
	"context"
	"database/sql"
	"log"
)

func (m *userRepository) IsProfileSet(ctx context.Context, uuid string) (bool, error) {

	query := `SELECT first_name, last_name, email, age FROM banking.users WHERE uuid = ?`

	rows, err := m.conn.Query(query, uuid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var firstName sql.NullString
	var lastName sql.NullString
	var email sql.NullString
	var age sql.NullInt64

	for rows.Next() {
		err := rows.Scan(&firstName, &lastName, &email, &age)
		if err != nil {
			log.Fatal(err)
		}
	}
	if firstName.Valid && lastName.Valid && email.Valid && age.Valid {
		return true, nil
	} else {
		return false, nil
	}
}
