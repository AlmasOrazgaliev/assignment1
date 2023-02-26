package model

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password"`
	EncryptedPassword string `json:"encryptedPassword"`
	IsSeller          bool   `json:"isSeller"`
	IsAdmin           bool   `json:"isAdmin"`
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
		u.Password = ""
	}
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
