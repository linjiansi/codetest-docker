package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/linjiansi/codetest-docker/src/usecase"
	"github.com/linjiansi/codetest-docker/src/util"
)

type AuthenticationMiddleware interface {
	Authentication(next http.Handler) http.Handler
}

type authenticationMiddleware struct {
	u usecase.UserUsecase
}

func NewAuthenticationMiddleware(u usecase.UserUsecase) AuthenticationMiddleware {
	return &authenticationMiddleware{u}
}

func (a *authenticationMiddleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Apikey")
		if apiKey == "" {
			appErr := util.NewAuthenticationError(errors.New("API key is required"))
			util.ReturnErrorResponse(w, appErr)
		}

		id, err := a.u.Authenticate(r.Context(), apiKey)
		if err != nil {
			util.ReturnErrorResponse(w, err)
		}

		ctx := context.WithValue(r.Context(), util.UserIdKey{}, id)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
