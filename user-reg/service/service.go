package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"
	"user-reg/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func New(storage MongoCli, cli RestyCli) Service {
	return &service{
		m: storage,
		r: cli,
	}
}

type service struct {
	m MongoCli
	r RestyCli
}

func (s *service) CreateUser(ctx context.Context, user *model.CreateUser) error {
	ctxT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// check exist
	existUser, err := s.m.GetUserByEmail(ctxT, user.Email)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	if existUser != nil {
		return model.ErrUserExist
	}
	// call salt service
	b, err := s.r.GetSalt(ctxT)
	if err != nil {
		return err
	}

	var salt struct {
		Salt string `json:"salt"`
	}

	if err = json.Unmarshal(b, &salt); err != nil {
		return err
	}

	h := md5.New()
	if _, err = h.Write([]byte(salt.Salt+user.Password)); err != nil {
		return err
	}

	user.Password = hex.EncodeToString(h.Sum(nil))
	user.Salt = salt.Salt

	return s.m.CreateUser(ctx, user)
}

func (s *service) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	ctxT, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	// check exist
	existUser, err := s.m.GetUserByEmail(ctxT, email)
	if err != nil {
		return nil, err
	}
	return existUser, nil
}
