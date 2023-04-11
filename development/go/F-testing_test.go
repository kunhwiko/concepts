// Go test files must have a suffix of _test.go
// Go test files conventionally are put into the same package as the code to be tested

// use "go test" to execute tests or "go test -cover" to test for coverage 
package go

import "testing"

// tests have a function parameter of type *testing.T
func TestArrayLength(t *testing.T) {
    type friends []string
    arr := friends{"Alice", "Ellie"}
    if len(arr) != 2 {
        t.Error("Expected length 2, but got ", len(arr))
    }
}