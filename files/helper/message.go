package helper

import "strings"

func CheckPrefix(msg string) error {
	msg = strings.TrimSpace(msg)
	if msg[0] != '&' {
		return &WrongPrefixError{}
	}
	return nil
}
