package assert

import (
	"reflect"
	"testing"
)

// ReflectEqual checks if two values of any type are deeply equal using reflect.DeepEqual.
//
// Parameters:
//   - t: A testing.TB instance to report errors.
//   - got: The actual value to compare.
//   - want: The expected value.
//
// Returns:
//   - None: If the values are not deeply equal, an error message is printed.
func ReflectEqual[T any](t testing.TB, got, want T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
