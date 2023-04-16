package go

import (
    "fmt"
)

var runIfStatement() {
    x := 43
    if x == 43 {
        fmt.Println(x)
    }

    // this limits x's scope only to the if loop 
    if x := 42; x == 42 {
        fmt.Println(x)
    }

    // returns 43
    fmt.Println(x)
}

var runSwitchStatement() {
    switch {                   // find the first true expression
    case true:
        fmt.Println("First")   // prints and breaks here   
    case 2==2:
        fmt.Println("Second")  // correct but does not print
    default:
        fmt.Println("Does not exist")
    }

    switch x := "Bond"; x {           // limits the scope of x to this switch loop
    case "James":
        fmt.Println("First")    
    case "Alice", "Mary":
        fmt.Println("Second")
    case "Bond":
        fmt.Println("Third")       
        fallthrough                   // fallthrough proceeds to next statement and does not break
    case "Ellie":
        fmt.Println("Fourth")         // previous fallthrough forces execution here although case is false
        fallthrough
    default:
        fmt.Println("Does not exist") // previous fallthrough forces execution here
    }
}

