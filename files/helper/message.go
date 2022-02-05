package helper

import (
	"errors"
	"strings"
)

var (
	wrongPrefixError = errors.New("wrong prefix")
)

func CheckPrefix(msg string) error {
	msg = strings.TrimSpace(msg)
	if msg[0] != '&' {
		return wrongPrefixError
	}
	return nil
}
