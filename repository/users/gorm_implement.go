package users

import (
	"context"

	"github.com/eduardylopes/db-go/entity"
	"gorm.io/gorm"
)

type repoDbGorm struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepository(writer *gorm.DB, reader *gorm.DB) UserRepositoryInterface {
	return &repoDbGorm{
		writer: writer,
		reader: reader,
	}
}

func (u *repoDbGorm) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User

	result := u.reader.Table("users").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}
