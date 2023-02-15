package storage

import (
	"context"
	"fmt"
	"user-reg/model"
	"user-reg/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func New(cli *mongo.Client) service.MongoCli {
	return &storage{users: cli.Database("authdb").Collection("users")}
}

type storage struct {
	users *mongo.Collection
}

func (s *storage) CreateUser(ctx context.Context, user *model.CreateUser) error {
	_, err := s.users.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("storage: CreateUser: %w", err)
	}
	return nil
}

func (s *storage) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	filter := bson.D{
		primitive.E{
			Key:   "email",
			Value: email,
		},
	}
	if err := s.users.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
