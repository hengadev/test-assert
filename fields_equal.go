package assert

import (
	"fmt"
	"reflect"
	"testing"
)

type FieldStructComparer interface {
	AssertComparable()
}

// CompareStructFields compares two instances of a struct type T that implements
// the Comparable interface, based on the specified fields.
//
// Parameters:
//   - t: Testing object to report errors.
//   - a, b: Instances of the struct instance of type T to compare.
//   - fields: A slice of strings representing the field names to compare.
//
// Returns:
//   - bool: true if all specified fields are equal, false otherwise.
//   - error: An error if an issue occurs during the comparison. Returns nil if the comparison is successful.
func FieldsEqual[T FieldStructComparer](t testing.TB, a, b *T, fields []string) (bool, error) {
	t.Helper()
	if a == nil || b == nil {
		return false, fmt.Errorf("cannot compare nil values")
	}

	// NOTE: I am not going to do Validation things here, this is not the purpose of that function
	// if !(*a).IsComparable() || !(*b).IsComparable() {
	// 	return false, fmt.Errorf("one or both instances are invalid")
	// }

	valA := reflect.ValueOf(*a)
	valB := reflect.ValueOf(*b)

	// Ensure we're working with struct types
	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		return false, fmt.Errorf("both arguments must be structs")
	}

	// Compare each specified field
	for _, fieldName := range fields {
		fieldA := valA.FieldByName(fieldName)
		fieldB := valB.FieldByName(fieldName)

		// Check if the field exists
		if !fieldA.IsValid() || !fieldB.IsValid() {
			return false, fmt.Errorf("field %s does not exist", fieldName)
		}

		// Check if the fields are comparable
		if !fieldA.Type().Comparable() {
			return false, fmt.Errorf("field %s is not comparable", fieldName)
		}

		// Compare the field values
		if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
			return false, nil
		}
	}

	return true, nil
}
