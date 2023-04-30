package go

import (
    "fmt"
    "sort"
)

type Element struct {
    Name string 
    Cost int 
}

/*
 * Go comparator.
 */
func sortComparator() {
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

/*
 * Defer functions delay execution until the end of the function. If multiple
 * defers exist, they will be executed in LIFO order.
 */
func defer() {
    defer func() {
        fmt.Println("Printing 1")
    }()

    defer func() {
        fmt.Println("Printing 2")
    }()
}