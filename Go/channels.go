package go
import (
	"fmt"
	"sync"
)

// Channels 
// channels are pipelines that help connect goroutines with each other 
// you can send values from one goroutine to another
// important : sends and receives will block until both the sender/receiver are ready

func bodyOne() {
    messages := make(chan string)

    // warning: this will fail
    // this is because the sender is ready and will block
    // but the receiver will never be ready, forcing a deadlock 
    messages <- "ping"

    go func() {
        // send a value into the channel
        // channel blocks until message is ready to be received 
        messages <- "ping"
    }()
    
    // receives a value from the channel
    // since the sender and receiver both block
    // channels allow us to wait without explicit synchronization  
    msg := <-messages 
    fmt.Println(msg)
}


// Buffered Channels 
// these are channels that allow pieces of data to sit in the channel

func bodyTwo() {
    messages := make(chan string, 1)

    // normally channels would block here and force a deadlock 
    // but our buffer of 1 allows one piece of data to sit in the channel 
    messages <- "hello"
    
    // although the sender and receiver can never be simultaneously ready 
    // our buffer allows the sent data to be received 
    msg := <-messages 
    fmt.Println(msg)
}

// Directional Channels 
// send-only channels : messages := make(chan<- string)
// receive-only channels : messages := make(<-chan string)

func bodyThree() {
    c := make(chan string)
    cs := make(chan<- string)

    // this is possible because a send-only channel is in fact a channel 
    cs = c 
    // this does not work because a channel is not a send-only channel 
    c = cs 
}


// Closing and Ranging Channels

func bodyFour() {
    c := make(chan int)

    go foo1(c) // send 

    // note that if bar was a goroutine, 
    // both routines might get terminated before executing
    bar1(c)  // receive 
}

func foo1(c chan<- int) {
    for i := 0; i < 10; i++ {
        // each send will block until bar() can receive 
        c <- i
    }
    // if we don't close, bar() will block until something is sent 
    // this forces a deadlock 
    close(c)
}

func bar1(c <-chan int) {
    // openly receive from the channel until it closes
    for v := range c {
        fmt.Println(v)
    }
}


// Select Keyword 

func bodyFive() {
    even := make(chan int)
    odd := make(chan int)
    quit := make(chan int)

    go send(even, odd, quit)
    receive(even, odd, quit)
}

func send(e, o, q chan<- int) {
    for i := 0; i < 100; i++ {
        if i % 2 == 0 {
            e <- i
        } else {
            o <- i
        }
    }
    q <- 0
}

func receive(e, o, q <-chan int) {
    for {
        select {
        // possible to use comma ok idiom 
        case v, ok := <-e:
            if ok {
                fmt.Println("From even channel:", v)  
            }   
        case v := <-o:
            fmt.Println("From odd channel:", v)
        // we don't have to close the channel because we return (terminate) here
        case v := <-q:
            fmt.Println("From quit channel", v)
            return 
        }
    }
}


// Design Patterns (Fan In & Fan Out)

// 1) Fan In: act of combining multiple outputs into one channel  

func receive1(even, odd <-chan int, fanin chan<- int) {
    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for v := range even {
            fanin <- v 
        }
        wg.Done()
    }()

    go func() {
        for v := range odd {
            fanin <- v 
        }
        wg.Done()
    }()
    wg.Wait()
} 

// 2) Fan-Out: act of dividing into more goroutines 

func fanOut(c1, c2 chan int) {
    var wg sync.WaitGroup
    for v := range c1 {
        wg.Add(1)
        go func(v2 int) {
            c2 <-work(v2)
            wg.Done()
        }(v)
    }
    wg.Wait()
    close(c2)
}