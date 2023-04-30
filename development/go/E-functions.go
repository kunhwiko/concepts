package go

import (
    "fmt"
)

/*
 * Go does not support inheritance and focuses more on composition. Go also focuses on
 * the creation of new "types".
 */
type names []string

/*
 * This is a "function" that declares a variable of type "names"
 */
func declareType() {
    var friends names = []string{"Alice", "Bob"}
    friends = names{"Cassandra", "Ellie"}
    friends.modifyFirst()
}

/*
 * This is a "method" for type "names" with a receiver "n". By convention, the instance
 * of the type is represented as a single letter.
 */
func (n names) modifyFirst() {
    n[0] = "Alice"
}

/*
 * Go is pass by value, so a copy of "x" is created when passing to the function.
 */
func declareFunction() {
    x, y := ReturnMultiple(1, "hello")
    
    // x will stay as 1
    add(x)

    // function literal - anonymous function
    // x is the original reference and not a copy
    func(z int) {
        x += z
    }(2)
}

/*
 * Functions can have multiple returns and capital case functions represent
 * that the function can be exported to external packages.
 */
func ReturnMultiple(x int, y string) (int, string) {
    return x, y
}

func add(x int) {
    x += 10
}
