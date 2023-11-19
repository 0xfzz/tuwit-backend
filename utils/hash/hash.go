package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(message string) []byte {
	hashed, err := bcrypt.GenerateFromPassword([]byte(message), bcrypt.DefaultCost)
	if err != nil {
		panic("Error while hashing")
	}
	return hashed
}

func Compare(hashedPassword []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
