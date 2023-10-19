package helpers

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return ""
	}
    return string(bytes)
}

func VerifyHash(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}