package util

type ErrorType interface {
}

type TError struct {
	Type    ErrorType
	Message string
	File    string
	Line    int
}

func (te *TError) Error() string {
	return te.Message
}
