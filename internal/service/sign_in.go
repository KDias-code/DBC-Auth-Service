package service

import (
	"auth-service/internal/model"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignIn(signModel model.SignIn) (string, error) {
	isPhone := s.phoneValidation(signModel.Phone)
	if !isPhone {
		return "", fmt.Errorf("incorrect phone number")
	}

	err := s.passwordValidation(signModel.Password)
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	password, err := s.redis.CheckUser(signModel.Phone, ctx)
	if err != nil {
		return "", err
	}

	err = s.compareHash(password, signModel.Password)
	if err != nil {
		return "", fmt.Errorf("phone or password is incorrect")
	}

	token, err := s.postgres.GetToken(signModel.Phone)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) compareHash(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
