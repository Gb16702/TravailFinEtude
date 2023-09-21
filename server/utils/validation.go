package utils

import (
	"regexp"
	"reflect"
	"strings"
	"errors"
	"fmt"
)

type EmailValidationError struct {
	Message string
}

/**
*	This function handles the error
*	@param string message
*   @implements error
*	@returns string
*/
func (e *EmailValidationError) Error() string {
	return e.Message
}

/**
*	This function checks if the value is nil
*
*	@param x interface{}
*	@returns bool
*/
func IsNilOrEmpty(x interface{}) bool {
	if x == nil {
		return true;
	}

	t := reflect.TypeOf(x);
	if t.Kind() == reflect.String {
		return len(strings.TrimSpace(x.(string))) == 0;
	}

	return false;
}

/**
*	This function checks if the value is a valid email
*
*	@param x interface{}
*	@param email string
*	@returns bool
*   @returns error
*/
func IsValidEmail(x interface{},) (bool, error) {
	const r = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$`;
	if matched, _ := regexp.MatchString(r, x.(string)); !matched {
		return false,
		&EmailValidationError{Message: fmt.Sprintf("Adresse mail invalide: %s", x)}};

	return true, nil;
}

/**
*	This function checks if the value is a valid password
*
*	@param password string
*	@returns bool
* 	@returns error
*/
func IsValidPassword(password string) (bool, error) {
	const minLength = 4;
	const maxLength = 32;

	if (len(password) < minLength) {
		return false, errors.New("Le mot de passe doit contenir au moins 4 caractères");
	}

	if (len(password) > maxLength) {
		return false, errors.New("Le mot de passe doit contenir au plus 20 caractères");
	}

	return true, nil;
}

/**
*	This function checks if the value is a valid password
*
*	@param password string
*	@param password_confirm string
*	@returns bool
*   @returns error
*/
func ArePasswordMatching(password, password_confirm string) (bool, error) {
	if(password != password_confirm) {
		return false, errors.New("Les mots de passe ne correspondent pas");
	}

	return true, nil;
}
