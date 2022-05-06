package db

import (
	"RestApi/internal/author/model"
	"RestApi/internal/author/storage"
	"RestApi/pkg/client/postgresql"
	"RestApi/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r *repository) Create(ctx context.Context, author *model.Author) error {
	q := `insert into author (name) values ($1) returning id`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	if err := r.client.QueryRow(ctx, q, author.Name).Scan(&author.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			sqlErr := fmt.Errorf("SQL Error: %s, Detail: %s, Where: %s, Code:%s, SQLState = %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			r.logger.Error(sqlErr)
			return sqlErr
		}
		return err
	}
	return nil
}

func (r *repository) FindOne(ctx context.Context, id string) (model.Author, error) {
	q := `select id, name from author where id = $1`

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	var auth model.Author
	if err := r.client.QueryRow(ctx, q, id).Scan(&auth.ID, &auth.Name); err != nil {
		return model.Author{}, err
	}
	return auth, nil
}

func (r *repository) Update(ctx context.Context, user model.Author) error {
	panic("implement me")
}

func (r *repository) Delete(ctx context.Context, id string) error {
	panic("implement me")
}

func (r *repository) FindAll(ctx context.Context, sortOptions storage.SortOptions) ([]model.Author, error) {
	q := squirrel.Select("id, name, age, is_alive, created_at").From("public.author")
	if sortOptions != nil {
		q = q.OrderBy(sortOptions.GetOrderBy())
	}

	sql, i, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	r.logger.Trace(fmt.Sprintf("SQL Query: %s", q))

	rows, err := r.client.Query(ctx, sql, i...)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)

	for rows.Next() {
		var auth model.Author

		err = rows.Scan(&auth.ID, &auth.Name, &auth.Age, &auth.IsAlive, &auth.CreatedAt)
		if err != nil {
			return nil, err
		}
		authors = append(authors, auth)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func NewRepository(client postgresql.Client, logger *logging.Logger) storage.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
