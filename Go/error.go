package go
import (
    "fmt"
    "os"
    "strings"
    "io"
    "log"
    "errors"
)

// error checking using print 

func createFile() {
    f, err := os.Create("newfiles.txt")

    if err != nil {
        fmt.Println("cannot create", err)
        return 
    }
    defer f.Close()

    data := strings.NewReader("hello world")
    io.Copy(f, data)
}

// error checking using logging 

func logging() {
    f, err := os.Create("log.txt")

    if err != nil {
        log.Println("cannot create", err)
    }
    defer f.Close()

    // for logging, you can send errors into a specified text file 
    log.SetOutput(f)

    f2, err := os.Open("does-not-exist.txt")
    if err != nil {
        log.Println("cannot open", err)
    }
    defer f2.Close()
}

// error checking using fatal 

func fatal() {
    f, err := os.Open("does-not-exist.txt")
    defer deferredFunction()

    // log.Fatalln prints and then calls os.Exit()
    if err != nil {
        log.Fatalln("cannot open", err)
    }
    defer f.Close()
}

// error checking using panic 

func panic() {
    f, err := os.Open("does-not-exist.txt")
    defer deferredFunction()

    // log.Panicln prints and then calls panic(err)
    // panic stops the execution of the current goroutine 
    // however, unlike fatal, deferred executions get executed
    if err != nil {
        log.Panicln("cannot open", err)
        // panic(err)
    }
    defer f.Close()   
}

func deferredFunction() {
    fmt.Println("hello there")
}


// using recover to come back from a panic 

func recoverExample() {
    f()
    fmt.Println("Returned from function f")
}

func f() {
    defer func() {
        // r will not be nil if we are currently in a panic 
        if r := recover(); r != nil {
            // r will call the panic statement 
            fmt.Println("Recovered in function f", r)
        }
    }()
    fmt.Println("Calling function g")
    g(0)
}

func g(i int) {
    // once i == 4, panic will terminate the current routine and run deferred calls 
    if i > 3 {
        log.Panicln("Caused panic with value : ", i)
    }
    defer fmt.Println("Deferring in function g, current value is : ", i)
    fmt.Println("Printing in function g", i)
    g(i + 1)
}

// RESULT 
// Calling function g 
// Printing in function g 0
// Printing in function g 1
// Printing in function g 2
// Printing in function g 3
// Deferring in function g, current value is : 3
// Deferring in function g, current value is : 2
// Deferring in function g, current value is : 1
// Deferring in function g, current value is : 0
// Recovered in function f Cause panic with value : 4


// throwing errors 

var customError = errors.New("That is not a valid operation")

func createErrors() {
    _, err := func(x int) (int, error) {
        if x < 0 {
            return 0, customError
        } 
        return x * x, nil
    }(-2)

    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Success without fatal")
}