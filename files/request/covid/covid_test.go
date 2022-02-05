package covid

import "testing"

func Test_GetCountryInfo(t *testing.T) {
	// Case 1: Error getting country
	_, err := GetCountryInfo("!%&#")
	if err == nil {
		t.Errorf("[TestGetCountryInfo] Incorrect results, want err: %v, got: %v", true, err)
	}

	// Case 2: Success
	_, err = GetCountryInfo("indonesia")
	if err != nil {
		t.Errorf("[TestGetCountryInfo] Incorrect results, want err: %v, got: %v", false, err)
	}
}
