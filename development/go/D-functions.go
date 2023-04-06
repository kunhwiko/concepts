package go

import (
    "fmt"
)

func main() {
    x, y := returnMultiple(1, "hello")

    // anonymous function
    z := func(x int) int {
        return x + 2
    }(2)
}

// functions can have multiple returns 
func returnMultiple(x int, y string) (int, string) {
    return x+1, y
}