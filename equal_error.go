package assert

import (
	"errors"
	"testing"
)

// EqualError checks if two errors are equal by comparing them with errors.Is.
//
// Parameters:
//   - t: A testing.TB instance to report errors.
//   - got: The actual error to compare.
//   - want: The expected error.
//
// Returns:
//   - None: If the errors are not equal, an error message is printed.
func EqualError(t testing.TB, got, want error) {
	switch {
	case got == nil && want == nil:
		return
	case got == nil || want == nil:
		t.Errorf("got '%v', want '%v'", got, want)
	case !errors.Is(got, want):
		t.Errorf("got '%v', want '%v'", got.Error(), want.Error())
	}
}
