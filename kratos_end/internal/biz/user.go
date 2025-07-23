package biz

import (
	"context"
)

type WlUser struct {
	Id           int32
	UserName     string
	UserNickname string
	Department   int32
	Mobile       string
	Email        string
	Password     string
	Gender       string
	Role         int32
	UserStatus   string
	Comment      string
}

type UserRepo interface {
	Register(ctx context.Context, user *WlUser) (*WlUser, error)
	Login(ctx context.Context, userName, password string) (*WlUser, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Register(ctx context.Context, user *WlUser) (*WlUser, error) {
	return uc.repo.Register(ctx, user)
}

func (uc *UserUsecase) Login(ctx context.Context, userName, password string) (*WlUser, error) {
	return uc.repo.Login(ctx, userName, password)
}
