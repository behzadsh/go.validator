package rules

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResult(t *testing.T) {
	res := NewResult(false, "failed")

	assert.True(t, res.Failed())
	assert.Equal(t, "failed", res.Message())
}

func TestIndirectValue(t *testing.T) {
	s := "string"
	data := struct {
		S *string
	}{
		S: &s,
	}

	v := indirectValue(data.S)
	assert.Equal(t, reflect.String, v.Kind())
}
