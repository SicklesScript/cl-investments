package auth

import (
	"github.com/alexedwards/argon2id"
)

/*
Encrypts passwords for storage in database
*/
func HashPassword(password string) (string, error) {
	// Created hashed password
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}
	return hash, nil
}

/*
Check that hashed password matches normal password for login
*/
func CheckPasswordHash(password, hash string) (bool, error) {
	// Check that entered password matches hashed password stored in db
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return match, err
	}
	return match, err
}
