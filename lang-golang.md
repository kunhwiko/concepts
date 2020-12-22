# Golang Tips 
### About Go 
---
##### Overview 
```
1) Strong / Static Typing
2) Compiled / Object Oriented Language* 
3) Mixed Cap Convention
4) Concurrency (built for multicores) + Compilation Efficient 
5) Pass by Value, allows for Pass by Pointer 
6) Automatic Garbage Collection 

* although OOP aspects exist, note the following : 
1) Go uses "types", not "classes"
2) Go does not "instantiate", you create a "value of a type"  
```

##### Workspace 
```
bin : holds compiled executable programs 
pkg : holds package objects so recompiling is unnecessary  
src : holds source files organized as packages  
```

##### Commands 
```
go version : version check
go env : environment check 
go help : help tool 
go fmt : format a file 
go fmt ./... : format all files under the current directory 
go run <filename> : run file 
go build <filename> : build an executable file in the current directory
go install <filename> : build executable file on the bin path 
```

##### Modules (learn more about modules, dependencies, packages)
```
1) a way to manage different dependencies 
2) collection of packages stored in a file tree with go.mod as the root 

go.mod 
1) defines the module path  
2) defines dependency requirements 
```

##### Identifiers 
```
name used to identify a variable, function, and any other user-defined item 
```


<br />


### Basics 
---
##### Declaring and Assigning 
```go
// short declaration operator cannot be declared globally 
func main() {
    x := 3
}

// var can be declared globally 
var y int8 = 1           // used to save memory space    

func main() {
    var z = 1            // untyped variable  
}
``` 

##### Numeric Types 
```
unsigned integers : uint8, 16, 32, 64
signed integers : int8, 16, 32, 64
float numbers: float32, 64
complex numbers : complex64, 128
byte = uint8
rune = int32
```

##### String Types
```go
s := "hello world"   // must use double quotes 
s := `
hello          // ` is equivalent of Python """
world
`

c := 'A'             // ascii number for char A 
b := []byte(s)       // ascii number for each char in s
```

##### Multiple Type Assignment 
```go
// untyped constants 
const (
    // untyped constants 
    a = 48 
    b = "Hello" 

    // typed constants 
    a int8 = 48 
    b string = "Hello" 
)

// alternatively 
var a, b int 
a = 30 
b = 30 
```

##### Looping
```go
for i := 0; i < 100; i++   // for loop 
for i < 10                 // while loop 
for                        // eternal loop 
```

##### If Statement 
```go
// this limits x's scope to the if loop 
if x := 42; x == 42 {
    fmt.Println(x)      // 42
}
fmt.Println(x)          // fails 
```

##### Switch Statement 
```go
switch {                   // find the first true expression
case true:
    fmt.Println("First")   // prints and breaks here   
case 2==2:
    fmt.Println("Second")  // correct but does not print
default:
    fmt.Println("Does not exist")
}

switch x := "Bond"; x {
case "James", "Bond", "Mark":   // James or Bond or Mark 
    fmt.Println("First")        // fallthrough proceeds to next statement 
    fallthrough           
case "Alice", "Mary":
    fmt.Println("Second")       // fallthrough prints even if statement is wrong 
    fallthrough           
default:
    fmt.Println("Does not exist")

}
```


<br />


### [Array/Slices]
---
##### Initialization
```go
// Array
var x [5]int   
var x [5]int{}
y := [5]int{4, 5, 6, 7, 8}
z := [5][5]int{}

// Slices (ArrayList)
var x []int 
y := []int{4, 5, 6, 7, 8}
z := [][]int{{1,2,3,4}, {5,6,7,8}}

// Slices with Make Function 
// if a slice is full, append operations copy existing elements to a larger slice 
// as this takes O(n) time, we can specify the capacity beforehand   
x := make([]int, 10, 100)
y := [][]int{x, x}

// Length and Capacity 
len(x)        // length is 10 
cap(x)        // total capacity is 100 
```

##### Composite Literal 
```
any array, slice, map, struct initialized with type and braces of elements 
[5]int{3,4,5,6}
map[string]int{"x1" : 1}
Person{first : "Alex", last : "Junior"}
```

##### Operations 
```go 
// Enhanced For Loop 
for i, v := range x {
    // i is index 
    // v is value 
}

// Others 
y := x[1:]
z := append(x, 77, 101, 123)
k := append(x, z...)
```


<br />


{Map}
---
##### Initialization 
```go 
// maps in Go are unsorted, default maps  
hm := map[string]int{"x1":1, "x2":-1, "x3":4} 
hm["x1"]               // 1
hm["x4"]               // 0 (does not give key error) 

var hm2 map[string]int // hm2 is a type map that currently points to nil 
hm2["x1"] = 1          // entry to nil error 
hm2 = make(map[string]int)  
```

##### Comma Ok Idiom 
```go
v, ok := hm["x4"]
v     // the value mapping 
ok    // whether or not the key is in the map 

if v, ok := hm["x4"]; ok {
  // do something 
} 
```

##### Operations 
```go
for k, v := range hm {
    // do something 
}

delete(hm, "x1")
delete(hm, "x4")   // no error for deleting unknown key 
```


<br />


### {Structs / Interfaces} 
---
##### Reference Type 
```
Slices, Maps, Channels are reference types 
Structs are NOT reference types 
Try passing a struct into a function and change it 
```

##### Structs
```go
type Rect struct {
    width, height float64
}

r := Rect{width: 3, height: 4}
r.width
r.height
```

##### Embedded Structs 
```go
type Coloredrect struct {
    // anonymous field, may also do "variable_name Rect"
    Rect       
    color string 
}

cr := Coloredrect{
    Rect: Rect{2.3, 3.4},
    color: "red",
} 
cr.Rect.width    // returns 2.3 
cr.width         // also returns 2.3 
```

##### Anonymous Structs 
```go
p1 := struct{
    first string 
    last string 
}{
    first: "Alex",
    last: "Junior",
}
```

##### Interfaces
```
type geometry interface {
    area() float64
    perim() float64
}

// to implement an interface, we implement all the methods 
type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perim() float64 {
    return 2*r.width + 2*r.height
}

// polymorphism
func measure(g geometry) {
    fmt.Println(g, g.area(), g.perim())
}
measure(rect{width: 3, height: 4})
```


<br />


### Function
---
##### Functions 
```go
func foo(parameter1 type, parameter2 type) (return_type1, return_type2) {

}
x, y := foo(argument1, argument2)

// anonymous functions 
func(x int) {
    fmt.Println(x + 2)
}(2)
```

##### First Class Citizen
```go
// function expressions 
f := func(x int) int {
    return x + 2
}
fmt.Println(f(2))

// returning a function 
func main() {
    f := bar()
    fmt.Printf("%T\n", f)    // func() string type
    fmt.Printf("%T", f())    // string type 
}

func foo() string {
    return "Hello World"
}

func bar() func() string {
    return foo
}
```

##### Variadic Parameter
```go
foo(2, 3, 4, 5, 6)
foo([]int{2, 3, 4, 5, 6}...)

// variadic means 0 or more int arguments can be passed 
func foo(x ...int) {
    fmt.Println(x)          // [2,3,4,5,6]
    fmt.Println("%T\n", x)  // slice type 
}

// 0 params creates a slice type pointing to nil 
foo()

// variadic parameters must be at the end 
func bar(s string, x...int) 
```

##### Methods 
```go
// a method is a function associated with an object (str.lower())
func (receiver type) foo() {}

p1.foo()      // method
bar(p1)       // function

func (p Person) foo() int {
    return p.first + p.second + p.third 
}

func bar(p Person) int {
    return p.first + p.second + p.third 
}
```

##### Defer
```go
// delays the execution of a function, method, or anonymous method 
// great for closing files (readability, multiple returns exist etc.)
// for multiple defers, executes in LIFO order
sum(1,2,3,4)
defer sum(2,3,4,5)
defer sum(1,2,3)
sum(5,6,7,8)

// --> 10, 26, 6, 14 
```


<br />


### Golang Things 
---
##### Printf 
```
%d    : integer 
%f    : float 
%v    : value 
%b    : binary 
%#x   : hexadecimal 
%T    : type 
```


<br />


### Logic Operations / Bit Manipulation 
---
##### iota 
```go 
// iota is only possible for constants, and auto increments by 1 
const (
    a = iota           // a = 0 
    b = iota           // b = 1
    c                  // c = 2 (don't have to specify iota)
)

const (
    d = 2016 + iota   // back to 0 --> 2016 
    e                 // e = 2017 
)
```

##### Utilizing iota 
```go
const (
    _ = iota               // unassigned variable 
    kb = 1 << (iota * 10)
    mb = 1 << (iota * 10)
    gb = 1 << (iota * 10)
)
```

