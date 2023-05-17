package go

import (
    "fmt"
)

// interfaces are not "concrete types", meaning we cannot directly create a value of an interface type
type geometry interface {
    area() float64
}

func isValid(g geometry) bool {
    return g.area() > 0
}
 
type rectangle struct {
    width, height float64
}

type triangle struct {
    width, height float64
}

// by implementing all the methods, we have implicitly inherited the interface 
func (r rectangle) area() float64 {
    return r.width * r.height
}

func (t triangle) area() float64 {
    return 0.5 * t.width * t.height
}

// interface allows us to use polymorphism to reuse code
// in other words, there is no need to write the same methods for different types 
func checkValid() {
    rect := rectangle{width: 4, height: 2}
    tri := triangle{width: 4, height: 3}
    fmt.Println(isValid(rect), isValid(tri))
}

// if an interface is specified as a field type, any concrete type that implements this interface can be used
type Response struct {
    data ReaderClosable
}

// to implement this interface, both Reader and Closable interface must also be satisfied 
type ReaderClosable interface {
    Reader
    Closable
}

type Reader interface {
    read() string
}

type Closable interface {
    close() string
}

// satisfies the Reader, Closable, and ReaderClosable interface
type ImageReader struct {
    fileExtension string
}

func (i ImageReader) read() string {
    return "Reading from " + i.fileExtension + " file extension" 
}

func (i ImageReader) close() string {
    return "Closing " + i.fileExtension + " file extension"
}

func checkExampleResponse() {
    imageReader := ImageReader{fileExtension: "jpg"}
    resp := Response{data: imageReader}
}