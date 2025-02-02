package usecase

import (
	"context"
	"github.com/linjiansi/codetest-docker/src/domain/repository"
)

type UserUsecase interface {
	Authenticate(ctx context.Context, apiKey string) (int, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo}
}

func (u *userUsecase) Authenticate(ctx context.Context, apiKey string) (int, error) {
	user, err := u.repo.FetchUser(ctx, apiKey)
	if err != nil {
		return 0, err
	}
	return user.Id(), nil
}
