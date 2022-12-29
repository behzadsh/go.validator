package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/behzadsh/go.validator/rules"
)

func TestSetDefaultLocale(t *testing.T) {
	SetDefaultLocale("es")

	assert.Equal(t, "es", defaultLocale)
}

func TestStopOnFirstFailure(t *testing.T) {
	StopOnFirstFailure()

	assert.True(t, stopOnFirstFailure)
}

func TestRegister(t *testing.T) {
	// empty registry
	registry = map[string]rules.Rule{}

	rule := &rules.After{}
	Register("after", rule)

	assert.Equal(t, rule, registry["after"])
	assert.NotEqual(t, rule, registry["something"])
}
