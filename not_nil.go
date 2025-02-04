package assert

import "testing"

// NotNil checks if the given error is not nil.
//
// Parameters:
//   - t: A testing.TB instance to report errors.
//   - got: The error to check.
//
// Returns:
//   - None: If the error is not nil, an error message is printed.
func NotNil(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got %q, want nil error", got)
	}
}
