package repository

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"

	model "github.com/gamepkw/users-banking-microservice/internal/models"
)

type userRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) UserRepository {
	return &userRepository{
		conn: conn,
	}
}

type UserRepository interface {
	RegisterUser(ctx context.Context, u *model.User) (res model.UserResponse, err error)
	GetHashedPasswordByUUID(ctx context.Context, uuid string) (res *model.UserResponse, err error)
	SetUpPin(ctx context.Context, u *model.Pin) (err error)
	SetNewPin(ctx context.Context, u *model.SetNewPin) (err error)
	GetHashedPinByUUID(ctx context.Context, uuid string) (res *model.Pin, err error)
	SetPassword(ctx context.Context, u *model.UpdatePassword) (res model.UserResponse, err error)
	SetUserProfile(ctx context.Context, u model.UserProfile) (err error)
	IsProfileSet(ctx context.Context, uuid string) (bool, error)
}

func isDuplicateEntryError(err error) bool {
	mysqlErr, ok := err.(*mysql.MySQLError)
	return ok && mysqlErr.Number == 1062
}
