package helpers

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func HassPass(p string) string {
	salt := 8
	password := []byte(p)

	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}
func ComparePass(h, p []byte) bool {
	hash, pass := []byte(h), []byte(p)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}

func IsEmailValid(email string) bool {
	// regular expression untuk pengecekan format email yang valid
	// pola ini akan memastikan bahwa email memiliki format yang benar, yaitu: username@domain.tld
	regex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	// membuat object regexp dari pola yang telah ditentukan
	re := regexp.MustCompile(regex)

	// melakukan pencocokan string email dengan pola yang telah ditentukan
	return re.MatchString(email)
}
