package crypto_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMd5(t *testing.T) {
	result := GetMd5("abc123")
	assert.NotNil(t, result)
	assert.EqualValues(t, "e99a18c428cb38d5f260853678922e03", result)
}
