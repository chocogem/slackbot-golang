package error


type ErrorAuthentication struct {
	Err error
}

func (e *ErrorAuthentication) Error() string {
	return e.Err.Error()
}

func NewErrorAuthentication(err error) error {
	return &ErrorAuthentication{Err: err}
}
