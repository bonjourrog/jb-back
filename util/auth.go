package util

import (
	"github.com/bonjourrog/jb/entity"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}

// VerifyRole returns true if the role given is a valid predefinied role.
func VerifyRole(role entity.Role) bool {
	switch role {
	case entity.RoleCompany, entity.RoleUser:
		return true
	}
	return false
}
