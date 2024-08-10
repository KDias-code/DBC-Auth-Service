package service

import (
	"auth-service/internal/model"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"unicode"
)

func (s *Service) SignUp(signModel model.SignUp) error {
	isPhone := s.phoneValidation(signModel.Phone)
	if !isPhone {
		return fmt.Errorf("not valid phone format")
	}

	err := s.passwordValidation(signModel.Password)
	if err != nil {
		return err
	}

	signModel.Password, err = s.hashPassword(signModel.Password)
	if err != nil {
		return err
	}

	ctx := context.Background()
	err = s.redis.CreateUser(signModel.Phone, signModel.Password, ctx)
	if err != nil {
		return err
	}

	err = s.postgres.SaveUser(signModel)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s *Service) passwordValidation(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("length of password is shortter than 8")
	}

	digitCounter := 0
	bigRegisterCounter := 0
	lowRigesterCounter := 0
	signCounter := 0

	for _, val := range password {
		if unicode.IsDigit(val) {
			digitCounter++
		} else if unicode.IsUpper(val) {
			bigRegisterCounter++
		} else if unicode.IsLower(val) {
			lowRigesterCounter++
		} else if unicode.IsPunct(val) {
			signCounter++
		}
	}

	if digitCounter < 0 {
		return fmt.Errorf("no digit in password")
	}
	if bigRegisterCounter < 0 {
		return fmt.Errorf("no upper register in password")
	}
	if lowRigesterCounter < 0 {
		return fmt.Errorf("no lower register in password")
	}
	if signCounter < 0 {
		return fmt.Errorf("no sign in password")
	}

	return nil
}

func (s *Service) phoneValidation(phone string) bool {
	if len(phone) < 11 || len(phone) > 11 {
		return false
	}

	if phone[0] != 7 || phone[1] != 7 {
		return false
	}

	return true
}
