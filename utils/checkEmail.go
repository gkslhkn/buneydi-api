package utils

import "regexp"

func CheckEmail(email string) bool {
	matched, err := regexp.Match(`^[\w\.]+@([\w-]+\.)+[\w-]{2,4}$`, []byte(email))
	if err != nil {
		matched = false
	}
	return matched
}
