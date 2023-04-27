 
package go

import (
    "fmt"
    "http"
)

// main goroutine executes the program line by line
// additional gooutines can be created to run functions concurrently
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

    // create new child goroutines
    // main function could terminate even if child routines are still running
    for _, link := range links {
        go checkLink(link, c)
    }

    // main routine is blocked until data is received from channel
    for i := 0; i < len(links); i++ {
        fmt.Println(<-c)
    }

    for _, link := range links {
        go checkLink(link, c)
    }

    // channels in a goroutine do not block the main routine
    // main function could terminate before all goroutines here finish
    for i := 0; i < len(links); i++ {
        go checkLink(<-c, c)
    }
}


func loopRoutine() {
    links := []string{
        "http://www.google.com",
        "http://www.amazon.com",
        "http://www.facebook.com",
        "http://www.workday.com",
        "http://www.netflix.com",
    }

    c := make(chan string)

    for _, link := range links {
        go checkLink(link, c)
    }

    // eternal loop that continues to listen on channel
    for l := range c {
        go checkLink(l, c)
    }
}


func checkLink(link string, c chan string) {
    _, err := http.Get(link)
    if err != nil {
        // send string information to a channel
        c <- err.Error()
        return
    }
    c <- "Link is up"
}
