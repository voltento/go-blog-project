package httperr

import "errors"

type httpError struct {
	error
	statusCode int
}

func WrapWithHttpCode(err error, code int) error {
	return &httpError{error: err, statusCode: code}
}

// HTTPStatusCode returns http status code from error or default error if not provided
func HTTPStatusCode(err error, defaultError int) int {
	var errorWithCode *httpError
	if errors.As(err, &errorWithCode) {
		return errorWithCode.statusCode
	}

	return defaultError
}
