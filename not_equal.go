package assert

import "testing"

// NotEqual checks if two values of type T are not equal.
//
// Parameters:
//   - t: A testing.TB instance to report errors.
//   - got: The actual value to compare.
//   - want: The expected value.
//
// Returns:
//   - None: If the values are equal, an error message is printed.
func NotEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got '%v', want '%v'", got, want)
	}
}
