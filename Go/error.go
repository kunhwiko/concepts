package Go 

import (
    "fmt"
    "os"
    "strings"
    "io"
    "log"
)

func createFile() {
    f, err := os.Create("newfiles.txt")

    // error checking using print 
    if err != nil {
        fmt.Println("cannot create", err)
        return 
    }
    defer f.Close()

    data := strings.NewReader("hello world")
    io.Copy(f, data)
}

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

func fatal() {
    f, err := os.Open("does-not-exist.txt")
    defer deferredFunction()

    // log.Fatalln prints and then calls os.Exit()
    if err != nil {
        log.Fatalln("cannot open", err)
    }
    defer f.Close()
}

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