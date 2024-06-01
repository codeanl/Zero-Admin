package MD5

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func SetPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return ""
	}
	//admin.Password = string(bytes)
	return string(bytes)
}

// CheckPassword 校验密码
func CheckPassword(md5password, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(md5password), []byte(password))
	return err == nil
}
