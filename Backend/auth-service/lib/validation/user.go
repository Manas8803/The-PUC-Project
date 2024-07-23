package validation

import (
	"errors"
	"regexp"

	"github.com/Manas8803/The-PUC-Project__BackEnd/auth-service/db"
)

func UserValidator(user *db.User) error {

	if user.Email == "" {
		return errors.New("ERROR : Email field must not be empty")
	}

	if !isValidEmail(user.Email) {
		return errors.New("ERROR : Invalid email")
	}

	if user.Password == "" {
		return errors.New("ERROR : Password field must not be empty")
	}

	if !isValidPassword(user.Email) {
		return errors.New("ERROR : Invalid password")
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`)
	return emailRegex.MatchString(email)
}

// * Criterias for valid password :
// At least 8 characters long
// Contains at least one uppercase letter
// Contains at least one lowercase letter
// Contains at least one digit
// Contains at least one special character (symbol)
func isValidPassword(password string) bool {

	lengthRegex := regexp.MustCompile(`^.{8,}$`)
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	digitRegex := regexp.MustCompile(`[0-9]`)
	specialCharRegex := regexp.MustCompile(`[!@#$%^&*()_+{}[\]:;<>,.?/~\\-]`)

	hasLength := lengthRegex.MatchString(password)
	hasUppercase := uppercaseRegex.MatchString(password)
	hasLowercase := lowercaseRegex.MatchString(password)
	hasDigit := digitRegex.MatchString(password)
	hasSpecialChar := specialCharRegex.MatchString(password)

	return hasLength && hasUppercase && hasLowercase && hasDigit && hasSpecialChar
}
