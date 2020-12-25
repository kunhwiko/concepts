package go
import (
    "fmt"
    "encoding/json"
)

type Person struct {
    First string
    Last string 
    Age int 
}

func converter() {
    // Marshal: creating a JSON object from maps 
    p1 := Person {"Alex", "Junior", 12}
    p2 := Person {"Alex", "Senior", 32}
    people := []Person{p1, p2}

    data, err := json.Marshal(people)
    if err != nil {
        fmt.Println("Failed to convert")
    }
    // JSON object is array of bytes  
    fmt.Println(string(data))   

    // Unmarshal: moving a JSON object back to a slice 
    err = json.Unmarshal(data, &people)
    if err != nil {
        fmt.Println("Failed to convert")
    }
    fmt.Println(people)
}