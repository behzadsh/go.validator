package bag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorBag_IsEmpty(t *testing.T) {
	errorBag := make(ErrorBag)

	assert.True(t, errorBag.IsEmpty())
	assert.Len(t, errorBag, 0)

	errorBag["test"] = []string{"error 1"}

	assert.False(t, errorBag.IsEmpty())
	assert.Len(t, errorBag, 1)
}

func TestErrorBag_Add(t *testing.T) {
	errorBag := make(ErrorBag)
	selector := "selector1"
	msg := "error 1 for selector 1"

	errorBag.Add(selector, msg)

	assert.Len(t, errorBag[selector], 1)
	assert.Equal(t, msg, errorBag[selector][0])
}

func TestErrorBag_All(t *testing.T) {
	errorBag := make(ErrorBag)
	errorBag["test"] = []string{"error 1", "error 2", "error 3"}
	errorBag["anotherTest"] = []string{"error 1"}

	assert.Len(t, errorBag.All(), 2)
	assert.Len(t, errorBag.All()["test"], 3)
	assert.Len(t, errorBag.All()["anotherTest"], 1)
}

func TestErrorBag_FirstOf(t *testing.T) {
	errorBag := make(ErrorBag)
	errorBag["test"] = []string{"error 1", "error 2", "error 3"}

	assert.Equal(t, "error 1", errorBag.FirstOf("test"))
	assert.Equal(t, "", errorBag.FirstOf("notExists"))
}

func TestErrorBag_Has(t *testing.T) {
	errorBag := make(ErrorBag)
	errorBag["test"] = []string{"error 1", "error 2", "error 3"}
	errorBag["anotherTest"] = []string{"error 1"}
	errorBag["emptyTest"] = []string{}
	errorBag["nilTest"] = nil

	assert.True(t, errorBag.Has("test"))
	assert.True(t, errorBag.Has("anotherTest"))
	assert.False(t, errorBag.Has("emptyTest"))
	assert.False(t, errorBag.Has("emptyTest"))
	assert.False(t, errorBag.Has("motExistsTest"))
}
