package users

import (
	"context"

	"github.com/eduardylopes/db-go/entity"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]entity.User, error)
}
