package go

import (
    "fmt"
    "strings"
)

// arrays have a fixed length
func initArray() {
    // initializes with 0s by default
    var x [5]int   
    y := [5]int{4, 5, 6, 7, 8}
    z := [5][5]int{}
}

// slices have dynamic length
func initSlice() {
    var x []int 
    y := []int{4, 5, 6, 7, 8}
    z := [][]int{{1,2,3,4}, {5,6,7,8}}
    
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

// enhanced for loop comes as (index, value) 
// idx and val are scoped to this loop and declaration will be thrown away
func runEnhancedLoopForSlice() {
    x := []string{"hello", "there", "world"}

    for idx, val := range x {
        fmt.Println(idx, val)
    }

    // idx and val are unknown at this point!
}

// adds an element to the slice
// this does not modify the existing slice but copies a new instance (i.e. O(n) operation)
func appendSlice() {
    x := []int{1, 2, 3}
    x = append(x, 4, 5, 6)
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

func initMap() {
    // empty is a type map that points to nil
    // assigning an entry will throw a nil exception
    var empty map[string]int
    empty["I"] = 1

    // initializes an empty map and we can now assign an entry
    empty = make(map[string]int)
    empty["I"] = 1

    // maps in Go are unsorted default maps
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