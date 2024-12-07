package utils

import "golang.org/x/crypto/bcrypt"


// func HashPassword(password string) (string, error) {
// 	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// }

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

