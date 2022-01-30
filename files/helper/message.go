package helper

import (
	"errors"
	"strings"
)

func CheckPrefix(msg string) error {
	msg = strings.TrimSpace(msg)
	if msg[0] != '&' {
		return errors.New("wrong prefix")
	}
	return nil
}
