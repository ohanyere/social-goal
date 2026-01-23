package store

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
)

type PostStore struct {
	db *sql.DB
}

type Post struct {
	Id int64 `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	UserID int64 `json:"user_id"`
	Tags  []string `json:"tags"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}




func NewPostStore(db *sql.DB) PostRepository {
	return &PostStore{db: db}
}


func (p *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`

	row := p.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	)

	err := row.Scan(&post.Id, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}
