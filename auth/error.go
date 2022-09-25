package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserDuplicate      = errors.New("username sudah digunakan")
	ErrEmailDuplicate     = errors.New("email sudah digunakan")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrUnknown            = errors.New("unknown error")
	ErrBadRequest         = errors.New("bad request bro")
	ErrUnauthorized       = errors.New("user unauthorized")
	ErrDataTidakLengkap   = errors.New("data tidak lengkap")
	ErrPasswordSame       = errors.New("password baru tidak boleh sama dengan password lama")
)
