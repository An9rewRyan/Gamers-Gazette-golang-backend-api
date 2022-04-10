package errors

type GetSessionError struct {
	Message string
}

func New_get_session_error(text string) error {
	return &DbConnectionError{text}
}

func (e *GetSessionError) Error() string {
	return e.Message
}
