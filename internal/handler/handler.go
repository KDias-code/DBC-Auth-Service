package handler

import (
	"auth-service/internal/model"
	"github.com/hashicorp/go-hclog"
)

type SignUp interface {
	SignUp(signModel model.SignUp) error
}

type SignIn interface {
	SignIn(signModel model.SignIn) (string, error)
}

type Handler struct {
	logger hclog.Logger
	signUp SignUp
	signIn SignIn
}

func NewHandler(logger hclog.Logger, signUp SignUp, signIn SignIn) *Handler {
	return &Handler{
		logger: logger,
		signUp: signUp,
		signIn: signIn,
	}
}
