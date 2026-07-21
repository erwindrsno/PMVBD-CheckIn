package auth

import (
	"fmt"
	"log/slog"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func VerifyPassword(password, storedPassword string) error {
	slog.Info("verify password", "password and stored password", fmt.Sprintf("%s\n%s", password, storedPassword))
	match, err := argon2id.ComparePasswordAndHash(password, storedPassword)
	if err != nil {
		return err
	}

	if !match {
		return ErrInvalidCredentials
	}
	return nil
}
