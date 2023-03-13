package postgres

import (
	Interface "GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Storage Хранилище данных.
type Storage struct {
	Db *pgxpool.Pool
}

// New Конструктор, принимает строку подключения к БД.
func New(ctx context.Context, constr string) (*Storage, error) {
	db, err := pgxpool.Connect(ctx, constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		Db: db,
	}
	return &s, nil
}

func (pg *Storage) Posts() ([]Interface.Post, error) {
	query := `SELECT id, title, content, author_id, authorName, created_at, publishedAt FROM posts;`
	rows, err := pg.Db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var posts []Interface.Post

	for rows.Next() {
		var p Interface.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.Created_at,
			&p.PublishedAt,
		)

		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	return posts, rows.Err()
}

func (pg *Storage) AddPost(p Interface.Post) (int, error) {
	var id int
	err := pg.Db.QueryRow(context.Background(), `
		INSERT INTO posts (author_id, title, content, authorName, created_at, publishedAt)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
		`,
		p.AuthorID,
		p.Title,
		p.Content,
		p.AuthorName,
		p.Created_at,
		p.PublishedAt,
	).Scan(&id)

	return id, err
}

func (pg *Storage) UpdatePost(p Interface.Post) error {
	query := `UPDATE posts set authorname = $4, content = $3, title = $2 WHERE author_id = $1;`
	_, err := pg.Db.Exec(context.Background(), query, p.AuthorID, p.Title, p.Content, p.AuthorName)
	return err
}

func (pg *Storage) DeletePost(p Interface.Post) error {
	delet := `DELETE FROM posts WHERE id = $1;`
	_, err := pg.Db.Exec(context.Background(), delet, p.ID)
	return err
}

//----------------------------------------------------------

func (pg *Storage) Postsuser() ([]Interface.Authors, error) {
	query := `SELECT id, name FROM authors;`
	rows, err := pg.Db.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	var user []Interface.Authors

	for rows.Next() {
		var p Interface.Authors
		err = rows.Scan(
			&p.Id,
			&p.Name,
		)

		if err != nil {
			return nil, err
		}

		user = append(user, p)
	}

	return user, rows.Err()
}

func (pg *Storage) AddPostuser(p Interface.Authors) (int, error) {
	var id int
	err := pg.Db.QueryRow(context.Background(), `
		INSERT INTO authors (name)
		VALUES ($1) RETURNING id;
		`,
		p.Name,
	).Scan(&id)

	return id, err
}

func (pg *Storage) UpdatePostuser(p Interface.Authors) error {
	query := `UPDATE authors set name = $2 WHERE id=$1;`
	_, err := pg.Db.Exec(context.Background(), query, p.Id, p.Name)
	return err
}

func (pg *Storage) DeletePostuser(p Interface.Authors) error {
	delet := `DELETE FROM authors WHERE id = $1;`
	_, err := pg.Db.Exec(context.Background(), delet, p.Id)
	return err
}
