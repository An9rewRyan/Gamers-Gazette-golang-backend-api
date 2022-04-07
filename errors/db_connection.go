package errors

type DbConnectionError struct {
	Message string
}

func New_db_connection_error(text string) error {
	return &DbConnectionError{text}
}

func (e *DbConnectionError) Error() string {
	return e.Message
}
