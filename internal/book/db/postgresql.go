package db

import (
	"RestApi/internal/author/model"
	"RestApi/internal/book"
	"RestApi/pkg/client/postgresql"
	"RestApi/pkg/logging"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, book *book.Book) error {
	panic("implement me")
}

func (r *repository) FindOne(ctx context.Context, id string) (book.Book, error) {
	panic("implement me")
}

func (r *repository) Update(ctx context.Context, book book.Book) error {
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context) ([]book.Book, error) {
	q := `select id, name from public.book;`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	books := make([]book.Book, 0)

	for rows.Next() {
		var b book.Book

		err = rows.Scan(&b.ID, &b.Name)
		if err != nil {
			return nil, err
		}

		sq := `select a.id, a.name
		from book_authors ba
		join public.author a on a.id = ba.author_id
		where book_id = $1;`

		authorsRows, err := r.client.Query(ctx, sq, b.ID)
		if err != nil {
			return nil, err
		}
		authors := make([]model.Author, 0)
		for authorsRows.Next() {
			var a model.Author

			err = authorsRows.Scan(&a.ID, &a.Name)
			if err != nil {
				return nil, err
			}
			authors = append(authors, a)
		}
		b.Authors = authors
		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) book.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
