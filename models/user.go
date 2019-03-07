package models

import (
	"regexp"

	"github.com/asaskevich/govalidator"
)

func init() {
	// custom validator for password
	govalidator.CustomTypeTagMap.Set("pass", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		password, ok := i.(string)
		if !ok {
			return false
		}

		// very strong password check
		// at least 1 uppercase letter, 1 lowercase, 1 numerical and 1 special
		regexp := regexp.MustCompile(
			"^[[[:alnum:]]*[[:graph:]]]*.{8,20}$")
		correct := regexp.MatchString(password)
		return correct
	}))
	// "?=.*[A-Z].)(?=.*[!@#$&*_])(?=.*[0-9])(?=.*[a-z]).{8}$"
	// "^[A-Z]+[0-9]+[a-zA-Z0-9_`!@#$%^&.?()-=+]*.{8,20}$"
	// custom validator for nickname
	govalidator.CustomTypeTagMap.Set("login", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		nickname, ok := i.(string)
		if !ok {
			return false
		}

		// only latin letters are allowed
		regexp := regexp.MustCompile(
			"^[a-zA-Z0-9_.]*.{3,20}$")
		correct := regexp.MatchString(nickname)
		if !correct {
			return false
		}
		if len(nickname) < 3 || len(nickname) > 16 {
			return false
		}
		return true
	}))

	// custom validator for fullname
	govalidator.CustomTypeTagMap.Set("name", govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
		fullname, ok := i.(string)
		if !ok {
			return false
		}

		// only latin letters are allowed
		regexp := regexp.MustCompile(
			"^[a-zA-Z]+$")
		correct := regexp.MatchString(fullname)
		if !correct {
			return false
		}

		if len(fullname) < 2 || len(fullname) > 32 {
			return false
		}
		return true
	}))
}
