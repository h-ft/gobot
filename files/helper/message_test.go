package helper

import (
	"testing"
)

func Test_CheckPrefix(t *testing.T) {
	// Case 1: Wrong prefix
	err := CheckPrefix("!play")
	if err != wrongPrefixError {
		t.Errorf("[TestCheckPrefixC1] Incorrect results, want: %s, got: %s", wrongPrefixError, err)
	}

	// Case 2: Success
	err = CheckPrefix("&play")
	if err != nil {
		t.Errorf("[TestCheckPrefixC2] Incorrect results, want: %v, got: %s", nil, err)
	}
}
