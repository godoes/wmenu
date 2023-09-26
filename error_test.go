package wmenu

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testInvalid    = newMenuError(ErrInvalid, "1", 0)
	testDuplicate  = newMenuError(ErrDuplicate, "2", 0)
	testTooMany    = newMenuError(ErrTooMany, "", 0)
	testNoResponse = newMenuError(ErrNoResponse, "", 0)
	errNormal      = errors.New(" Oops")
	errMenu        = newMenuError(errors.New("general"), "", 0)
)

func TestIsInvalidErr(t *testing.T) {
	a := assert.New(t)
	a.True(IsInvalidErr(testInvalid))
	a.False(IsInvalidErr(testDuplicate))
	a.False(IsInvalidErr(errNormal))
}

func TestIsDuplicateErr(t *testing.T) {
	a := assert.New(t)
	a.True(IsDuplicateErr(testDuplicate))
	a.False(IsDuplicateErr(testInvalid))
	a.False(IsDuplicateErr(errNormal))
}

func TestIsTooManyErr(t *testing.T) {
	a := assert.New(t)
	a.True(IsTooManyErr(testTooMany))
	a.False(IsTooManyErr(testDuplicate))
	a.False(IsTooManyErr(errNormal))
}

func TestIsNoResponseErr(t *testing.T) {
	a := assert.New(t)
	a.True(IsNoResponseErr(testNoResponse))
	a.False(IsNoResponseErr(testDuplicate))
	a.False(IsNoResponseErr(errNormal))
}

func TestIsMenuError(t *testing.T) {
	a := assert.New(t)
	a.True(IsMenuErr(testNoResponse))
	a.True(IsMenuErr(testDuplicate))
	a.True(IsMenuErr(testInvalid))
	a.True(IsMenuErr(testTooMany))
	a.True(IsMenuErr(testNoResponse))
	a.True(IsMenuErr(errMenu))
	a.False(IsMenuErr(errNormal))
}
