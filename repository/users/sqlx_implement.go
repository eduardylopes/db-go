package users

import (
	"context"

	"github.com/eduardylopes/db-go/entity"
	"github.com/jmoiron/sqlx"
)

type repoDbSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepository(writer *sqlx.DB, reader *sqlx.DB) UserRepositoryInterface {
	return &repoDbSqlx{
		writer: writer,
		reader: reader,
	}
}

func (u *repoDbSqlx) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	query := "select * from users"

	if err := u.reader.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}
