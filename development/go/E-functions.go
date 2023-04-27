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
}

// "methods" belong to a particular type where this one has a receiver of type 'names' called 'n'
// by convention, the instance of the type is represented as a single letter
func (n names) modifyFirst() {
    n[0] = "Alice"
}

func declareFunction() {
    x, y := ReturnMultiple(1, "hello")
    
    // Go is pass by value, so a copy of "x" is created when passing to function
    // x will stay as 1
    add(x)

    // function literal - anonymous function
    // x is the original reference and not a copy
    func(z int) {
        x += z
    }(2)
}

// functions can have multiple returns
// capital case functions are exported to external packages
func ReturnMultiple(x int, y string) (int, string) {
    return x, y
}

func add(x int) {
    x += 10
}
