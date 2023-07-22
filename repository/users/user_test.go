package users_test

import (
	"context"
	"testing"

	"github.com/eduardylopes/db-go/configs"
	"github.com/eduardylopes/db-go/repository/users"
)

func BenchmarkGetAllGorm(b *testing.B) {
	ctx := context.Background()
	connection := users.NewGormRepository(configs.GetReaderGorm(), configs.GetWriterGorm())
	for i := 0; i < b.N; i++ {
		connection.GetAll(ctx)
	}
}

func BenchmarkGetAllSqlx(b *testing.B) {
	ctx := context.Background()
	connection := users.NewSqlxRepository(configs.GetReaderSqlx(), configs.GetWriterSqlx())
	for i := 0; i < b.N; i++ {
		connection.GetAll(ctx)
	}
}
