# Go Test Helpers


A Go package providing a set of helper functions to simplify common testing scenarios. It includes functions for equality checks, error comparisons, deep equality assertions, and more to make your tests more readable and concise.


## Features


- **Equal**: Check if two values are equal.
- **EqualError**: Check if two errors are equal (or one is `nil`).
- **NotEqual**: Ensure that two values are not equal.
- **NotNil**: Ensure that an error is not `nil`.
- **ReflectEqual**: Perform deep equality checks between two values.
- **ReflectNotEqual**: Ensure that two values are not deeply equal.

## Installation


You can install this package by running the following command:

```bash
go get github.com/yourusername/testhelpers
```
## Usage


Here's how you can use the functions in your test files:


### Equal

```go
package yourpackage

import (
	"testing"
	"github.com/GaryHY/test-assert"
)

func TestEqual(t *testing.T) {
	assert.Equal(t, 5, 5) // Passes
	assert.Equal(t, 5, 6) // Fails with "got '5', want '6'"
}
```

### EqualError

```go
func TestEqualError(t *testing.T) {
	assert.EqualError(t, nil, nil) // Passes
	assert.EqualError(t, fmt.Errorf("error"), fmt.Errorf("error")) // Passes
	assert.EqualError(t, fmt.Errorf("error1"), fmt.Errorf("error2")) // Fails
}
```

### ReflectEqual

```go
func TestReflectEqual(t *testing.T) {
	testhelpers.ReflectEqual(t, struct{ Name string }{"John"}, struct{ Name string }{"John"}) // Passes
	testhelpers.ReflectEqual(t, struct{ Name string }{"John"}, struct{ Name string }{"Doe"}) // Fails with "got '{Name:John}', want '{Name:Doe}'"
}
```


## Contributing


Feel free to fork the repository, open issues, and submit pull requests. Contributions are always welcome!

## License


This project is licensed under the MIT License - see the LICENSE file for details.

### Key Points in the README:


- **Overview**: A short explanation of the package.
- **Features**: The main functions of the package and their purpose.
- **Installation**: Instructions on how to install the package.
- **Usage**: Examples of how to use the provided functions in tests.
- **Contributing**: Encouragement to contribute with guidelines.
- **License**: Including the MIT license information (or adjust based on your licensing preference).


Feel free to modify the content to suit your needs.
