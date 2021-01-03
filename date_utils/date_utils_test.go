package date_utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.EqualValues(t, "2006-01-02T15:04:05Z", apiDateLayout)
	assert.EqualValues(t, "2006-01-02 15:04:05", apiDBLayout)
}

func TestGetNow(t *testing.T) {
	res := GetNow().Unix()
	assert.EqualValues(t, res, time.Now().UTC().Unix())
}

func TestGetNowString(t *testing.T) {
	nowString := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	res := GetNowString()
	assert.EqualValues(t, res, nowString)
}

func TestGetNowDBFormat(t *testing.T) {
	nowString := time.Now().UTC().Format("2006-01-02 15:04:05")
	res := GetNowDBFormat()
	assert.EqualValues(t, res, nowString)
}
