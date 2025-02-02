package util

import (
	"context"
	"errors"
)

type UserIdKey struct{}

func GetUserId(ctx context.Context) (int, error) {
	id, ok := ctx.Value(UserIdKey{}).(int)
	if !ok {
		return 0, NewAuthenticationError(errors.New("user id not found in context"))
	}
	return id, nil
}
