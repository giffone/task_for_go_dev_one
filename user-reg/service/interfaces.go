package service

import (
	"context"
	"user-reg/model"
)

type RestyCli interface {
	GetSalt(ctx context.Context) ([]byte, error)
}

type MongoCli interface {
	CreateUser(ctx context.Context, user *model.CreateUser) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type Service interface {
	CreateUser(ctx context.Context, user *model.CreateUser) error
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}
