package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	res := NewResult()

	assert.NotNil(t, res)
	assert.Empty(t, res.Errors)
}

func TestResult_Failed(t *testing.T) {
	res := NewResult()
	assert.False(t, res.Failed())

	res.addError("something", "error")
	assert.True(t, res.Failed())
}
