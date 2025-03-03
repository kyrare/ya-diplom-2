package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func LoginValidator(login string) error {
	if login == "" {
		return errors.New("login must not be empty")
	}

	if utf8.RuneCountInString(login) < 3 {
		return errors.New("login must be 6 characters long")
	}

	if utf8.RuneCountInString(login) > 255 {
		return errors.New("login must be shorter than 255 characters")
	}

	return nil
}

func PasswordValidator(password string) error {
	if password == "" {
		return errors.New("password must not be empty")
	}

	if utf8.RuneCountInString(password) < 6 {
		return errors.New("password must be 6 characters long")
	}

	if utf8.RuneCountInString(password) > 255 {
		return errors.New("password must be shorter than 255 characters")
	}

	return nil
}

// Validator functions to ensure valid input
func CcnValidator(s string) error {
	// Credit Card Number should a string less than 20 digits
	// It should include 16 integers and 3 spaces
	if len(s) > 16+3 {
		return fmt.Errorf("CCN is too long")
	}

	if len(s) == 0 || len(s)%5 != 0 && (s[len(s)-1] < '0' || s[len(s)-1] > '9') {
		return fmt.Errorf("CCN is invalid")
	}

	// The last digit should be a number unless it is a multiple of 4 in which
	// case it should be a space
	if len(s)%5 == 0 && s[len(s)-1] != ' ' {
		return fmt.Errorf("CCN must separate groups with spaces")
	}

	// The remaining digits should be integers
	c := strings.ReplaceAll(s, " ", "")
	_, err := strconv.ParseInt(c, 10, 64)

	return err
}

func ExpValidator(s string) error {
	// The 3 character should be a slash (/)
	// The rest should be numbers
	e := strings.ReplaceAll(s, "/", "")
	_, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return fmt.Errorf("EXP is invalid")
	}

	// There should be only one slash and it should be in the 2nd index (3rd character)
	if len(s) >= 3 && (strings.Index(s, "/") != 2 || strings.LastIndex(s, "/") != 2) {
		return fmt.Errorf("EXP is invalid")
	}

	return nil
}

func CvvValidator(s string) error {
	// The CVV should be a number of 3 digits
	// Since the input will already ensure that the CVV is a string of length 3,
	// All we need to do is check that it is a number
	_, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("CVV is invalid")
	}
	return nil
}
