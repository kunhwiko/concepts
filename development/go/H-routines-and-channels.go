package go

import (
    "fmt"
    "net/http"
    "time"
)

/*
 * Main goroutine executes the program line by line and additional goroutines 
 * can be created to run functions concurrently. Note that the main routine 
 * does not wait for all the child routines to finish executing.
 */
func initRoutine() {
    links := []string{
        "http://www.google.com",
        "http://www.amazon.com",
        "http://www.facebook.com",
        "http://www.workday.com",
        "http://www.netflix.com",
    }

    // create a channel for goroutines to communicate
    c := make(chan string)

    // sends data to a channel
    // this will deadlock as a successful send requires a ready receiver
    c <- "hello world"

    // create new child goroutines
    for _, link := range links {
        go checkLink(link, c)
    }

    // this blocks main routine from terminating until data is received from channel
    for i := 0; i < len(links); i++ {
        fmt.Println(<-c)
    }

    for _, link := range links {
        go checkLink(link, c)
    }

    // channels used in a goroutine do not block the main routine
    // main routine could terminate before all goroutines here finish
    for i := 0; i < len(links); i++ {
        go checkLink(<-c, c)
    }
}

func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        c <- err.Error()
        return
    }
    c <- "Link is up"
}

func dataRace() {
    // goroutine here uses the same reference to i that the main routine is using
    // Go disallows routines to reference the same non-reference type variables to prevent unexpected behaviors
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println(i)
        }()
    }

    // this is the correct way
    for i := 0; i < 10; i++ {
        go func(i int) {
            i += 1
        }(i)
    }
}


func loopRoutine() {
    links := []string{
        "http://www.google.com",
        "http://www.amazon.com",
    }

    c := make(chan string)

    for _, link := range links {
        go checkLink(link, c)
    }

    // eternal loop that continues to listen on channel
    for l := range c {
        go func() {
            time.Sleep(5 * time.Second)
            checkLink(l, c)
        }()
    }
}
