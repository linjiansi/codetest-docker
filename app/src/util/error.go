package util

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorCode string
type ErrorMessage string

type AppError struct {
	Code      ErrorCode
	Message   ErrorMessage
	BaseError error
}

const (
	CodeAuthenticationError ErrorCode = "AUTHENTICATION_ERROR"
	CodePaymentError        ErrorCode = "PAYMENT_ERROR"
	CodeDataNotFound        ErrorCode = "DATA_NOT_FOUND"
	CodeUnexpectedError     ErrorCode = "UNEXPECTED_ERROR"
)

const (
	MessageAuthenticationError ErrorMessage = "認証に失敗しました。"
	MessagePaymentError        ErrorMessage = "支払いに失敗しました。"
	MessageDataNotFound        ErrorMessage = "データが見つかりませんでした。"
	MessageUnexpectedError     ErrorMessage = "不明なエラーが発生しました。"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (e *AppError) Error() string {
	return string(e.Message)
}

func NewAppError(code ErrorCode, message ErrorMessage, err error) AppError {
	return AppError{
		Code:      code,
		Message:   message,
		BaseError: err,
	}
}

func NewAuthenticationError(err error) *AppError {
	return &AppError{
		Code:      CodeAuthenticationError,
		Message:   MessageAuthenticationError,
		BaseError: err,
	}
}

func NewPaymentError(err error) *AppError {
	return &AppError{
		Code:      CodePaymentError,
		Message:   MessagePaymentError,
		BaseError: err,
	}
}

func NewInternalServerError(err error) *AppError {
	return &AppError{
		Code:      CodeUnexpectedError,
		Message:   MessageUnexpectedError,
		BaseError: err,
	}
}

func NewDataNotFoundError(err error) *AppError {
	return &AppError{
		Code:      CodeDataNotFound,
		Message:   MessageDataNotFound,
		BaseError: err,
	}
}

func ReturnErrorResponse(w http.ResponseWriter, err error) {
	var appErr *AppError
	ok := errors.As(err, &appErr)
	if !ok {
		appErr = NewInternalServerError(err)
	}

	var detail string
	if appErr.BaseError != nil {
		detail = appErr.BaseError.Error()
	}

	var statusCode int
	switch appErr.Code {
	case CodeAuthenticationError:
		statusCode = http.StatusUnauthorized
	case CodePaymentError:
		statusCode = http.StatusPaymentRequired
	case CodeDataNotFound:
		statusCode = http.StatusNotFound
	case CodeUnexpectedError:
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Message: string(appErr.Message),
		Detail:  detail,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
