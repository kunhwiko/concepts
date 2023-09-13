package go

import (
    "fmt"
    "strings"
)

/*
 * Arrays have a fixed length and are initialized with default values (e.g. 0) by default.
 */
func initArray() {
    var x [5]int   
    y := [5]int{4, 5, 6, 7, 8}
    z := [5][5]int{}
}

/*
 * Slices have dynamic length and when full, existing elements need to be copied over to a
 * larger slice (i.e. O(n) complexity). 
 */
func initSlice() {
    var x []int 
    y := []int{4, 5, 6, 7, 8}
    z := [][]int{{1,2,3,4}, {5,6,7,8}}

    // can optionally specify the initial length and capacity beforehand
    m := make([]int, 10, 100)
    fmt.Println(len(m))  // prints 10
    fmt.Println(cap(m))  // prints 100
    
    // slice ranges will create a new copy of the slice
    fmt.Println(y[1:4])
    fmt.Println(y[:2])
    fmt.Println(y[1:])
}

func runLoopForSlice() {
    sum := 0

    // basic for loop
    for i := 0; i < 10; i++ {
        sum += i
    }

    // basic while loop
    for sum < 100 {
        sum += sum
    }

    // eternal while loop
    for {
        if sum > 20 {
            break
        }
        sum++
    }
}

/*
 * Enhanced for loops come as (index, value) and are scoped only to that loop.
 */
func runEnhancedLoopForSlice() {
    x := []string{"hello", "there", "world"}

    for idx, val := range x {
        fmt.Println(idx, val)
    }

    // idx and val are unknown at this point!
}

/*
 * Adding an element to the slice does not modify the existing slice but creates a 
 * new instance (i.e. O(n) operation).
 */
func appendSlice() {
    x := []int{1, 2, 3}
    x = append(x, 4, 5, 6)
    y := []int{7, 8, 9}
    x = append(x, y...)
}

func joinSlice() {
    // join strings
    x := []string{"hello", "there", "world"}
    str := strings.Join(x, ", ")

    // split strings
    x = strings.Split(str, ", ")

    // repeat strings
    str = strings.Repeat(str, 4)
}

/*
 * Maps in Go are unsorted default maps.
 */
func initMap() {
    // empty is a map type that points to nil
    // assigning an entry will throw a nil exception
    var empty map[string]int
    empty["I"] = 1

    // initializes an empty map that we can now assign an entry
    empty = make(map[string]int)
    empty["I"] = 1

    numbers := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
    }
    
    numbers["I"]  // returns 1
    numbers["VX"] // returns 0 and does not return key error
    
    // put and delete items
    empty["C"] = 1000
    delete(empty, "C")
}

func runEnhancedLoopForMap() {
    numbers := map[string]int{
        "I": 1,
        "V": 5,
        "X": 10,
    }

    for key, val := range numbers {
        fmt.Println(key, val)
    }

    // comma ok idiom
    // ok is a boolean telling us whether or not the key is in the map
    if v, ok := numbers["C"]; ok {
        fmt.Println(v, "This should not print")
    } 
}