package go

import (
    "fmt"
)

// interfaces are not "concrete types", meaning we cannot directly create a value of an interface type
type geometry interface {
    area() float64
}

// interface allows us to use polymorphism to reuse code
// in other words, there is no need to write the same methods for different types 
func isValid(g geometry) bool {
    return g.area() > 0
}

// by implementing all the methods, we have implicitly inherited the interface 
type rectangle struct {
    width, height float64
}

func (r rectangle) area() float64 {
    return r.width * r.height
}

// by implementing all the methods, we have implicitly inherited the interface 
type triangle struct {
    width, height float64
}

func (t triangle) area() float64 {
    return 0.5 * t.width * t.height
}

func checkValid() {
    rect := rectangle{width: 4, height: 2}
	tri := triangle{width: 4, height: 3}
	fmt.Println(isValid(rect), isValid(tri))
}