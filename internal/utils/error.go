package utils

type Error struct {
	message string
}

func NewError(msg string) error {
	return &Error{message: msg}
}

func (e *Error) Error() string {
	return e.message
}
