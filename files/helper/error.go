package helper

type error interface {
	Error() string
}

type WrongPrefixError struct {
	Message string
}

func (m *WrongPrefixError) Error() string {
	return "Wrong message prefix!"
}
