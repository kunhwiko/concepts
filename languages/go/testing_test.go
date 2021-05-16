// a test file must have a suffix of _test.go, like this file
// conventionally you will put the test file in the same package as the file tested

// for test, we must have a function parameter of type *testing.T

// use "go run main.go" to execute
// use "go test main.go" to test 
// use "go test -bench ." to benchmark 

// use "go test -cover" to test for coverage 
// use "go test -coverprofile c.out" to create a file c.out with coverage info
// use "go tool cover -html=c.out" to take c.out and see what was covered on a webpage



package go 

import "testing"

// assume we have func mySum(x ...int) int {}
// conventionally, we name the function as TestMySum() 

func TestMySum(t *testing.T) {
    if mySum(2, 3) != 5 {
        t.Error("Expected", 5, "Got", 6)
    }
}


// table tests to do multiple assertions at once 

func TestMySum2(t *testing T) {
    type test struct {
        data []int
        answer int
    }

    tests := []test {
        test{[]int{9, 10}, 19},
        test{[]int{3, 4, 5}, 12},
        test{[]int{-1, -3}, -4},
    }

    for _, v := range test {
        sum := mySum(v.data...)
        if x != v.answer {
            t.Error("Expected", v.answer, "Got", x)
        }
    }
}


// example tests 
// this adds an "Example" into Go Documentation 
// reference : https://godoc.org/errors 
// write the function, and then in COMMENTS, write the expected output 

func ExampleSum() {
    fmt.Println(mySum(2, 3))
    // Output:
    // 5
}


// benchmarking 

func BenchmarkSum(b *testing.B) {
    for i := 0; i < b.N; i++ {
        mySum(2, 3)
    }
}