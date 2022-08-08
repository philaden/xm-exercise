package helpers

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 4)
	return string(hash), err
}

func HashPasswordAdmin(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 8)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ContainsSpecialCharacter(s string) bool {
	for i := 0; i < len(s); i++ {
		switch b := s[i]; {
		case b >= 'a' && b <= 'z':
			continue
		case b >= 'A' && b <= 'Z':
			continue
		case b >= '0' && b <= '9':
			continue
		default:
			return true
		}
	}
	return false
}

func ContainsCapitalLetter(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			return true
		}
	}
	return false

}

func ValidateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}
