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

/*
 * Functions in Go are first class citizens, meaning they can be assigned to
 * variables, returned, and passed as parameters
 */
func firstClassCitizen() {
    // functions can be assigned to variables
	f := func(x int) int {
		return x + 2
	}
	fmt.Println(f(2))

	// functions can be returned and assigned
	f = bar()
	fmt.Println(f(2))

    // functions can be assigned to parameters - callback functions
    run(hello)

    // variadic parameters - reference the below
    declareVariadicParams("hello", []int{1, 2, 3, 4}...)
    declareVariadicParams("hello")
}

func foo(x int) int {
    return x
}

func bar() func(x int) int {
    return foo
}

func hello() {
    fmt.Println("Hello")
}

func run(f func()) {
    f()
}

/*
 * Variadic means 0 or more arguments can be passed. Note that variadic
 * parameters must be specified at the end. The variadic parameter here
 * represents a slice of integers.
 */
func declareVariadicParams(s string, i ...int) {
    // empty parameter would create an empty slice
    if len(i) == 0 {
        fmt.Println("Empty slice")
    }
    fmt.Println(s, i)
}