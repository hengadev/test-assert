package assert

import (
	"reflect"
	"testing"
)

// ReflectNotEqual checks if two values of any type are not deeply equal using reflect.DeepEqual.
//
// Parameters:
//   - t: A testing.TB instance to report errors.
//   - got: The actual value to compare.
//   - want: The expected value.
//
// Returns:
//   - None: If the values are deeply equal, an error message is printed.
func ReflectNotEqual[T any](t testing.TB, got, want T) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
