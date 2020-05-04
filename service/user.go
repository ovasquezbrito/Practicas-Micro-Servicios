package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User contains user's information
type User struct {
	Username       string
	HashedPassword string
	Role           string
}

// NewUser return a new user
func NewUser(username string, password string, role string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %v", err)
	}

	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}
	return user, nil
}

// IsCorrectPassword checks if the provided password id correcto or ni!
func (u *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}

// Clone returns a clone of this user
func (u *User) Clone() *User {
	return &User{
		Username:       u.Username,
		HashedPassword: u.HashedPassword,
		Role:           u.Role,
	}
}
