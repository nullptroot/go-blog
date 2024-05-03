package pwd

import (
	"golang.org/x/crypto/bcrypt"
)

/*
 * hash 加密密码
 * @param string pwd 待加密的明文密码
 */
func HashPwd(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

/*
 * 验证 hash 密码
 * @param string hashedPwd 已加密的hash密码
 * @param string sourcePwd 确认密码
 */
func CheckPwd(hashedPwd string, sourcePwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(sourcePwd))
	return err == nil
}
