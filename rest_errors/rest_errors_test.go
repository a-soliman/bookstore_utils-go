package rest_errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test message", errors.New("database error"))
	assert.NotNil(t, err)
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, 500, err.Status())

	assert.NotNil(t, 1, len(err.Causes()))
	assert.EqualValues(t, "database error", err.Causes()[0])
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, 400, err.Status())
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, 404, err.Status())
}

func TestNewUnAuthorizedError(t *testing.T) {
	err := NewUnAuthorizedError("test message")
	assert.NotNil(t, err)
	assert.EqualValues(t, "test message", err.Message())
	assert.EqualValues(t, 401, err.Status())
}
