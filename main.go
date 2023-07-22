package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eduardylopes/db-go/configs"
	"github.com/eduardylopes/db-go/repository"
)

func main() {

	repo := repository.New(repository.Options{
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterGorm: configs.GetWriterGorm(),
		ReaderGorm: configs.GetReaderGorm(),
	})

	users, err := repo.User.GetAll(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
