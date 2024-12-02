package utils

import (
	"dolott_authentication/internal/types"

	"github.com/matthewhartstonge/argon2"
)

var argon = argon2.DefaultConfig()

func HashPassword(password []byte) (string, *types.Error) {
	encoded, err := argon.HashEncoded(password)
	if err != nil {
		return "", types.NewInternalError("internal issue, error code #9001")
	}
	return string(encoded), nil
}

func VerifyPassword(password []byte, encodedPassword []byte) (bool, *types.Error) {
	ok, err := argon2.VerifyEncoded(password, encodedPassword)
	if err != nil {
		return false, types.NewInternalError("internal issue, error code #9002")

	}
	return ok, nil
}
