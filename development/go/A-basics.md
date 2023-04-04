### Modules
---
##### Packages
```
Packages organize Go source files into a unit that is modular, reusable, and maintainable.
As a convention, a package is represented in its own directory and the directory name is used as the package name. 
```

##### Modules
```
Modules are a collection of packages stored in a file tree with go.mod as the root.
The go.mod file defines the import path for the module and dependency requirements required for a successful build.
The go.sum file records checksums of dependencies represented in hashes that are used to confirm that dependencies were not tampered. 
```

##### Module Cache
```
Packages that are downloaded are saved and cached in $GOPATH/pkg/cache.
When a Go program requires a package, it will read from the cache instead of having to go through the Internet.   
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
Variables and functions are available in other packages if they start with a capital letter.
Otherwise, variables and functions are considered private to the current package.
```