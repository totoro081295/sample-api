package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashPassword パスワードの照合
func CompareHashPassword(hash string, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return nil
	}
	return err
}

// ToHashFromBcrypt ハッシュ化のメソッド
func ToHashFromBcrypt(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), 10)
	return string(converted)
}
