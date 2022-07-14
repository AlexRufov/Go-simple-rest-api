package db

import (
	"Go-simple-rest-api/internal/book"
	"Go-simple-rest-api/pkg/client/postgresql"
	"Go-simple-rest-api/pkg/logging"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, book book.Book) (string, error) {
	q := `insert into book
    (name, age, created_at)
	SELECT $1::varchar(100), $2, $3
	where NOT EXISTS (
        	SELECT name, age, created_at FROM book WHERE name = $1 and age = $2 and created_at = $3
   		)
	RETURNING book.id;`

	booksRows, err := r.client.Query(ctx, q, &book.Name, &book.Age, &book.CreatedAt)
	if err != nil {
		return "", err
	}
	var bookID string
	for booksRows.Next() {
		err = booksRows.Scan(&book.ID)
		if err != nil {
			return "", err
		}
		bookID = book.ID
	}
	r.logger.Infof("Book created id = %s", bookID)
	return bookID, nil
}

func (r *repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	q := `select id, name, age, created_at from book where id = $1;`
	var b book.Book
	rows, err := r.client.Query(ctx, q, id)
	if err != nil {
		return b, err
	}

	for rows.Next() {
		err = rows.Scan(&b.ID, &b.Name, &b.Age, &b.CreatedAt)
		if err != nil {
			return b, err
		}
	}

	if err = rows.Err(); err != nil {
		return b, err
	}
	r.logger.Infof("Book with id = %s found", id)
	return b, nil
}

func (r *repository) Update(ctx context.Context, book book.Book) error {
	q := `update book
	set name = $1::varchar(100), age = $2
	where id = $3;`
	_, err := r.client.Exec(ctx, q, &book.Name, &book.Age, &book.ID)
	if err != nil {
		return err
	}
	r.logger.Infof("Book with id = %s updated", book.ID)
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	q := `delete from book where id = $1;`
	_, err := r.client.Exec(ctx, q, id)
	if err != nil {
		return err
	}
	r.logger.Infof("Book with id = %s deleted", id)
	return nil
}

func (r *repository) FindAll(ctx context.Context) ([]book.Book, error) {
	q := `select id, name, age, created_at from book;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	books := make([]book.Book, 0)

	for rows.Next() {
		var b book.Book

		err = rows.Scan(&b.ID, &b.Name, &b.Age, &b.CreatedAt)
		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	r.logger.Infof("Found %d books", len(books))
	return books, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
