package store

import (
	"context"
	"database/sql"
)

type User struct{
Id int64 `json:"id"`
Email string `json:"email"`
UserName string `json:"username"`
Password string `json:"-"`
CreatedAt string  `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func NewUserStore(db *sql.DB) UserRepository {
	return &UserStore{db: db}
}

func (u *UserStore) CreateUser(ctx context.Context, user *User) error {
	Query := `
	INSERT INTO users(username, email, password)
	VALUES($1, $2, $3) RETURNING id, created_at
	`
	row := u.db.QueryRowContext(
		ctx,
		Query,
		user.UserName,
		user.Email,
		user.Password,
	)

	err := row.Scan(&user.Id, &user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}
	
