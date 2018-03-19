package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAPIError = &APIError{
	Errors: []ErrorDetail{
		{Message: "Not found.", Code: 404},
	},
}
var errTestError = fmt.Errorf("unknown host")

func TestAPIError_ErrorString(t *testing.T) {
	err := &APIError{}
	assert.Equal(t, "", err.Error())
	assert.Equal(t, "waves-go-client: 404 Not found.", testAPIError.Error())
}

func TestAPIError_Empty(t *testing.T) {
	err := APIError{}
	assert.True(t, err.Empty())
	assert.False(t, testAPIError.Empty())
}

func TestFirstError(t *testing.T) {
	cases := []struct {
		httpError error
		apiError  *APIError
		expected  error
	}{
		{nil, nil, nil},
		{nil, &APIError{}, nil},
		{nil, testAPIError, testAPIError},
		{errTestError, &APIError{}, errTestError},
		{errTestError, testAPIError, errTestError},
	}
	for _, c := range cases {
		err := FirstError(c.httpError, c.apiError)
		assert.Equal(t, c.expected, err)
	}
}
