package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/linjiansi/codetest-docker/src/domain/model"
	"github.com/linjiansi/codetest-docker/src/domain/repository"
	"github.com/linjiansi/codetest-docker/src/repository/data_model"
	"github.com/linjiansi/codetest-docker/src/util"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) repository.UserRepository {
	return &userRepository{db}
}

func (u *userRepository) FetchUser(ctx context.Context, apiKey string) (*model.User, error) {
	user := &data_model.User{}
	err := u.db.Get(user, "SELECT * FROM users WHERE api_key = ?", apiKey)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, util.NewDataNotFoundError(err)
		} else {
			return nil, util.NewInternalServerError(err)
		}
	}
	return model.NewUser(
		user.ID,
		user.Name,
		user.ApiKey,
	), nil
}
