# Go Programming
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
go doc : find the documentation in the current directory 

go fmt : format a file 
go fmt ./... : format all files under the current directory 
golint : tells us mistakes that are not best practice 

go run <filename> : run file 
go build <filename> : build an executable file in the current directory
go install <filename> : build executable file on the bin path 
```

##### Packages 
```
1) Convention is to use the directory name as the package name 
2) Variables can only be exported if they start with upper case, otherwise private
```

##### Modules 
```
All dependency packages and installations are saved in GOPATH, 
meaning the PATH must change when switching projects

Modules are a way to manage different dependencies, and support the following
- ability to work from any directory, not just the GOPATH
- ability to install a precise version of a package
- ability to import multiple versions of the same package 
- ability to list all dependencies in a project  

Modules are a collection of packages stored in a file tree with go.mod as the root 

go.mod 
1) defines the module import path  
2) defines dependencies used 
3) go mod init github.com/example to initialize module and specify import path
```

##### Documentation
```
go doc commands 
1) go doc : documentation for current package 
2) go doc Example : documentation for Example in the current package
3) go doc fmt : high level documentation for the fmt package

godoc commands 
1) godoc -http=:8080 : see go documentations on local port without internet access 
2) godoc fmt : thorough documentation for the fmt package
3) godoc -src fmt Printf : see the implementation (not the documentation) of Printf

writing documentations 
1) comments without spaces above a package or function will generate documentations 
2) by convention, we start the sentence with Package <package name> ... for package documentations  
3) by convention, we start the sentence with capital letters for function documentations 

// Example is a function that does not do anything.
// It is simply meant to show an example of documenting.
func example() {

} 

For thorough examples, check out 
https://golang.org/pkg/errors/#pkg-overview
https://golang.org/src/errors/errors.go?s=1875:1902#L48
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
// var can be declared globally 
var x int8 = 1     // int8 can be used to save memory 

func main() {
    // short declaration operator cannot be declared globally 
    y := 3 

    // untyped variable 
    var z = 1 
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

##### Multiple Type Assignment 
```go
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


### "String"
---
##### String Types
```go
s := "hello world"           // must use double quotes 
s := `hello world`           // ` is equivalent of Python """

c := 'A'                     // ascii number for char A 
b := []byte(s)               // ascii number for each char in s
```

##### String Operations 
```go
for i, v := range s {
    fmt.Printf("%T\n", s[i])    // byte type 
    fmt.Printf("%T\n", v)       // rune type 
}

// methods 
strings.Repeat("string", 4)
strings.Join([]string{"My", "Name", "Is", "Jenny"}, " ")
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

// Slices 
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
map[string]int{"x1": 1}
Person{first: "Alex", last: "Junior"}
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

// Sort Packages  
sort.Ints(arr) 
sort.Strings(arr)
```


<br />


### {Map}
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
type ColoredRect struct {
    Rect          // anonymous field, may also do "varName Rect"
    color string 
}

cr := ColoredRect{
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

r := rect{width: 3, height: 4}
measure(r)
```


<br />


### Function
---
##### Functions 
```go
// functions 
func foo(x int, y string) (int, string) {
    // functions in Go can have multiple returns 
    return x+1, y
}

// anonymous functions 
func(x int) {
    fmt.Println(x + 2)
}(2)
```

##### First Class Citizen
```go
// functions in Go are First-Class Citizens
// they can be assigned to variables, returned, and passed

// 1) function expressions 
f := func(x int) int {
    return x + 2
}
fmt.Println(f(2))


// 2) returning a function 
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


// 3) callback functions 
func main() {
    run(hello)
}

func hello() {
    fmt.Println("Hello")
}

func run(f func()) {
    f()
}
```

##### Variadic Parameter
```go
// variadic means 0 or more int arguments can be passed 
func foo(x ...int) {
    fmt.Println(x)          // [2,3,4,5,6]
    fmt.Println("%T\n", x)  // slice type 
}

foo(2, 3, 4, 5, 6)
foo([]int{2, 3, 4, 5, 6}...)

// 0 params create a slice type pointing to nil 
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
// delays the execution of a function or method 
// great for closing files (readability, multiple returns exist etc.)

// for multiple defers, executes in LIFO order
sum(1,2,3,4)       
defer sum(2,3,4,5)
defer sum(1,2,3)
sum(5,6,7,8)

// --> 10, 26, 6, 14 
```


<br />


### Pointer*
---
##### Pointer
```go
// pointers are great for passing in addresses of large chunks of data 

&a     // address 
*int   // pointer to the address of an int 
*a     // dereference an address 
```

##### Dereferencing Structs
```go
type ListNode struct {
    Val int 
    Next *ListNode
}

var ptr1 *ListNode
(*ptr1).Next = p2    // this is what you would expect   
ptr1.Next = p2       // but you can also just do this 
```

##### Method Set 
```go
// type *T can call methods with receiver type T and *T
// type T can call methods with receiver type T, and if T is addressable, type *T
// if type T is not addressable, cannot call methods with receiver type *T 

// this is because 
// for *T, we know exactly what T it is pointing to, enabling us to call methods with receiver T 
// for T, we might have multiple *T's pointing to it, disabling us from calling methods with receiver *T 

type circle struct {
    radius float64
}

type shape interface {
    area() float64
}

func (c *circle) area() float64 {
    return 3.14 * c.radius * c.radius 
}

func getArea(s shape) {
    fmt.Println(s.area())
}

func main() {
    c := circle{5}
    // if type T is not addressable, cannot call methods with receiver type *T 
    getArea(c)

    // if type T is addressable, can call methods with receiver type *T       
    c.area()        
}
```


<br />


### Golang Things 
---
##### Creating Type
```go
type trythis int     // a type called "trythis"
var a trythis = 123
var b int = 123 

b == int(a)          // Conversion 
```

##### Printf 
```
%d    : integer 
%f    : float 
%v    : value 
%b    : binary 
%#x   : hexadecimal 
%T    : type 
```

##### Conversion 
```go
int(x)              // changing floats to int 
string(x)           // changing byte to string 

strconv.Itoa(32)    // changing int to string
strconv.Atoi("32")  // changing string to int (Atoi returns int, error)
```

##### Math
```go
math.Pow(2, 31)     // returns a float 
```

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

##### Custom Sorting / Comparator 
```go
type Element struct {
    Name string 
    Cost int 
}

func main() {
    e1 := Element{"Ruby", 32000}
    e2 := Element{"Sapphire", 29000}
    e3 := Element{"Silver", 5400}
    e4 := Element{"Jade", 29000}

    elements := []Element{e1, e2, e3, e4}
    sort.SliceStable(elements, func(i, j int) bool {
        if elements[i].Cost == elements[j].Cost {
            return elements[i].Name < elements[j].Name
        }
        return elements[i].Cost < elements[j].Cost
    })
    fmt.Println(elements)
}
```




