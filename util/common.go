/*
Copyright © 2022 niuzhiqiang <niuzhiqiang90@foxmail.com>

*/
package util

import "regexp"

func VerifyEmailFormat(email string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	emailReg := regexp.MustCompile(pattern)
	return emailReg.MatchString(email)
}
