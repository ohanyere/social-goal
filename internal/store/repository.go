package store

import "context"

type PostRepository interface {
	Create(ctx context.Context,  *Post) error
}

type UserRepository interface {
	CreateUser(ctx context.Context, *User) error
}
