package go

import (
    "fmt"
)

// Go does not support inheritance and focuses more on composition
// Go allows for the creation of new types
type names []string

// "functions" do not belong to a particular type
func declareType() {
    // declare names type
    var friends names = []string{"Alice", "Bob"}
    friends = names{"Cassandra", "Ellie"}
    friends.modifyFirst()

    x, y := ReturnMultiple(1, "hello")

    // anonymous function
    z := func(x int) int {
        return x + 2
    }(2)
}

// "methods" belong to a particular type where this one has a receiver of type 'names' called 'n'
// by convention, the instance of the type is represented as a single letter
func (n names) modifyFirst() {
    n[0] = "Alice"
}

// functions can have multiple returns
// capital case functions are exported to external packages
func ReturnMultiple(x int, y string) (int, string) {
    return x+1, y
}