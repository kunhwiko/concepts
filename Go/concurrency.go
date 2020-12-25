package Go

import (
	"fmt"
	"sync"
)

// Parallelism vs Concurrency
// parallelism is the simultaneous execution of computation 
// concurrency is a design pattern that potentially enables programs to run in parallel  


// Goroutines 
// much like threads, allow functions or methods to execute independently
func sum(x ...int) int {
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum 		
}

func body1() {
	// these will run concurrently 
	go sum([]int{1, 2, 3, 4, 5}...)
	go sum([]int{6, 7, 8, 9, 1}...)
}


// WaitGroups 
// the code above will likely not get executed
// this is because the program will likely reach the bottom of main() first
// when it does, all goroutines are also terminated, so how do we fix this?
func sum2(wg *sync.WaitGroup, x ...int) int {
    // Done() decrements the number of waitgroups 
    defer wg.Done()
    sum := 0
    for _, v := range x {
        sum += v
    }
    return sum 	
}

func body2() {
	// create a WaitGroup
    var wg sync.WaitGroup

    for i := 0; i <= 5; i++ {
        // Add() increments the number of waitgroups 
        wg.Add(1) 
        go sum2(&wg, []int{1,2,3,4,5}...) 
    } 
    // wait until WaitGroup count drops to 0 
    wg.Wait()
}


// Race Conditions and Mutex Locks 
// race conditions mean different routines are accessing the same resource
// "go run -race main.go" command can find race conditions
// locks can lock out certain resources 

var num = 0

func body3() {
    var wg sync.WaitGroup 
    var mu sync.Mutex 

    wg.Add(100)
    for i := 1; i <= 100; i++ {
        go func(val int) {
            mu.Lock()
            num += val
            fmt.Println(num)
            mu.Unlock()
            wg.Done()
        }(i)
    }
    wg.Wait()
}

