package service

import (
	"auth-service/internal/model"
	"context"
)

type Postgres interface {
	SaveUser(signModel model.SignUp) error
	CreateToken(phone, token string) error
	GetToken(phone string) (string, error)
}

type Redis interface {
	CreateUser(phone, password string, ctx context.Context) error
	CheckUser(phone string, ctx context.Context) (string, error)
}

type Service struct {
	postgres Postgres
	redis    Redis
}

func NewService(postgres Postgres, redis Redis) *Service {
	return &Service{
		postgres: postgres,
		redis:    redis,
	}
}
