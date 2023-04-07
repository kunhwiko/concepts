package go

import (
    "fmt"
)

// Go does not support inheritance and focuses more on composition
// Go allows for the creation of new types
type names []string

func declareType() {
    var friends names = []string{"Alice", "Bob"}
    friends = names{"Cassandra", "Ellie"}
    friends.modifyFirst()
}

// This is a method that has a receiver of type 'names' called 'n'
// By convention, the instance of the type is represented as a single letter
func (n names) modifyFirst() {
    n[0] = "Alice"
}