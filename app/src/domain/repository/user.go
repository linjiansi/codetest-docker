package repository

import (
	"context"

	"github.com/linjiansi/codetest-docker/src/domain/model"
)

type UserRepository interface {
	FetchUser(ctx context.Context, apiKey string) (*model.User, error)
}
