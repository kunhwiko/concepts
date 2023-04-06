package go

import (
    "fmt"
)

// variable can be declared globally
// we can assign as int8 to conserve memory
var x int8 = 1

func declareVariables() {
    // multiple declarations
    var i, j int = 1, 2

    // short variable declaration - not possible outside functions
    k := 3 

    // multiple short variable declaration
    a, b, c := true, 1, "hello"

    // untyped variable 
    var y = 1

    fmt.Println(x, i, j, k, a, b, c, y)
}