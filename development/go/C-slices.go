package go

import (
    "fmt"
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

func runLoop() {
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


    x := []string{"hello", "there", "world"}

    // enhanced for loop comes as (index, value) 
    // idx and val are scoped to this loop and declaration will be thrown away
    for idx, val := range x {
        fmt.Println(idx, val)
    }
}

// enhanced for loop comes as (index, value) 
// idx and val are scoped to this loop and declaration will be thrown away
func runEnhancedLoop() {
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
