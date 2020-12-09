# Golang Tips 
### About Go 
---
##### Overview 
```
1) Strong / Static Typing
2) Compiled / Object Oriented Language
3) Concurrency (built for multicores) + Compilation Efficient 
4) Pass by Value, allows for Pass by Pointer 
5) Automatic Garbage Collection 
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

    // untyped constants 
    a int8 = 48 
    b string = "Hello" 
)
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
len(x)

// Slices
x := []int{4,5,6,7,8}
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

