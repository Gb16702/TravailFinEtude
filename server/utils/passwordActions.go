package utils

import (
	"github.com/alexedwards/argon2id"
	"errors"
)

/**
*	This function hashes a password
*
* 	@params password string
*	@returns string
* 	@returns error
*/
func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", err
	}

	return hash, nil
}

/**
*	This function compares the request body password with the password from the database
*
* 	@params requestPassword
* 	@params userPassword
*	@returns err error
*/
func ComparePassword(requestPassword, userPassword string) (err error) {
	match, _, err := argon2id.CheckHash(requestPassword, userPassword);
	if err != nil {
		return err;
	}

	if !match {
		return errors.New("Invalid credentials")
	}

	return nil;
}