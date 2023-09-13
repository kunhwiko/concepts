### Basics
---
##### Overview
```
a) Go is a statically typed, strongly typed, compiled language.
b) Go uses the concepts of types and structs instead of classes. Go also heavily uses composition for OOP and does not 
   use inheritance.
c) Go is pass by value, refer to https://neilalexander.dev/2021/08/29/go-pass-by-value. Slices, maps, and channels are 
   reference types but structs are not.  
d) Go supports automated garbage collection.
```

##### Numeric Types
```
unsigned integers : uint8 (byte), uint16, uint32, uint64
signed integers   : int8, int16, int32 (rune), int64
float numbers     : float32, float64
complex numbers   : complex64, complex128
```

### Modules
---
##### Packages
```
Packages organize Go source files into a unit that is modular, reusable, and maintainable. As a convention, a package is 
represented in its own directory and the directory name is used as the package name. 
```

##### Modules
```
Modules are a collection of packages stored in a file tree with go.mod as the root. The go.mod file defines the import 
path for the module and all dependency requirements. The go.sum file records checksums of dependencies represented in 
hashes that are used to confirm that dependencies were not tampered. 
```

##### Module Cache
```
Packages that are downloaded are saved and cached in $GOPATH/pkg/cache. When a Go program requires a package, it will 
read from the cache instead of having to go through the Internet.   
```

##### Module Commands
```go
// check Go environment variables (e.g. $GOPATH)
go env

// initializes a new module
go mod init

// list all modules and whether they need to be upgraded
go list -m -u all

// displays a graph of dependencies
go mod graph

// add missing and remove unused modules
go mod tidy

// checks go.sum to verify 
go mod verify
```

##### Exportability
```
Variables and functions are available in other packages if they start with a capital letter. Otherwise, variables and 
functions are considered private to the current package.
```

### Goroutines
---
##### Goroutines vs Threads
```
a) Goroutines are managed by the Go runtime (i.e. scheduler) rather than by the OS. 
b) Goroutines consume a smaller stack size (~2KB) than threads (1MB+), making context switching faster. The stack is 
   used to save local variables of function calls and has the capability to grow dynamically by allocating heap storage.
c) Goroutines can communicate and share information efficiently with other goroutines through channels.
```

##### Main Goroutines
```
The main goroutine is created when a program is launched. Any subsequently created goroutines are child routines.
```

##### Concurrency and Parallelism
```
Golang can achieve concurrency through goroutines (i.e. if a goroutine is blocked, another can be scheduled to execute 
work). Golang can achieve parallelism through multiple CPU cores by configuring GOMAXPROCS.
```

##### Go Scheduler
```
Responsible for scheduling and monitoring code running in goroutines. Although a CPU core can only execute a single 
goroutine at a time, multiple goroutines can still run concurrently.
```
