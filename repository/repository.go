package repository

import (
	"github.com/eduardylopes/db-go/repository/users"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

type Container struct {
	User users.UserRepositoryInterface
}

func New(options Options) *Container {
	return &Container{
		User: users.NewGormRepository(options.WriterGorm, options.ReaderGorm),
		// User: users.NewSqlxRepository(options.WriterSqlx, options.ReaderSqlx),
	}
}
