package errors

type GetApiArticleError struct {
	Message string
}

func Get_api_article_error(text string) error {
	return &DbConnectionError{text}
}

func (e *GetApiArticleError) Error() string {
	return e.Message
}
