package database

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (p *PostRepository) Create(postCreate CreatePost) error {
	query := sq.Insert("public.post").
		Columns("post_image", "post_description", "create_time", "user_id").
		Values(postCreate.PostImage, postCreate.PostDescription, postCreate.CreateTime, postCreate.UserId).
		RunWith(p.db).
		PlaceholderFormat(sq.Dollar)
	_, err := query.Query()

	return err

}
