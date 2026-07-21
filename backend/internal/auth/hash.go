package auth

import "github.com/alexedwards/argon2id"

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

func VerifyPassword(password, storedPassword string) error {
	match, err := argon2id.ComparePasswordAndHash(password, storedPassword)
	if err != nil {
		return err
	}

	if !match {
		return ErrInvalidCredentials
	}
	return nil
}
