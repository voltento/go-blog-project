package httperr

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var originalError = errors.New("original error")

func TestWrapWithHttpCode(t *testing.T) {
	httpCode := 404

	wrappedError := WrapWithHttpCode(originalError, httpCode)
	assert.NotNil(t, wrappedError, "expected non-nil wrapped error")

	httpErr, ok := wrappedError.(*httpError)
	assert.True(t, ok, "expected *httpError type, got %T", wrappedError)

	assert.Equal(t, httpCode, httpErr.statusCode, "expected status code to match")
	assert.Equal(t, originalError, httpErr.error, "expected original error to match")
}

func TestHTTPStatusCodeWithWrappedError(t *testing.T) {
	httpCode := 400

	wrappedError := WrapWithHttpCode(originalError, httpCode)

	result := HTTPStatusCode(wrappedError, 500)
	assert.Equal(t, httpCode, result, "expected status code to match wrapped error's status code")
}

func TestHTTPStatusCodeWithNonWrappedError(t *testing.T) {
	defaultCode := 500

	result := HTTPStatusCode(originalError, defaultCode)
	assert.Equal(t, defaultCode, result, "expected default status code when error is not wrapped")
}

func TestHTTPStatusCodeWithNilError(t *testing.T) {
	defaultCode := 500

	result := HTTPStatusCode(nil, defaultCode)
	assert.Equal(t, defaultCode, result, "expected default status code when error is nil")
}

func TestHTTPStatusCodeUnknownErrorType(t *testing.T) {
	defaultCode := 500

	result := HTTPStatusCode(originalError, defaultCode)
	assert.Equal(t, defaultCode, result, "expected default status code when error is nil")
}
