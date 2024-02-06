package go

import (
    "fmt"
)

// "structs" are a collection of fields
type contactInfo struct {
    email string
    phone int
}

type address struct {
    country string
    city    string
}

// structs can be embedded
type person struct {
    firstName string
    lastName  string
    contact   contactInfo
    address                 // "address" will be translated as a field name "address" of type address
}

func declareType() {
    // declare person type 
    var alice person
    ellie := person{"Ellie", "Williams", contactInfo{"ellie@gmail.com", 123456}, address{"United States", "Boston"}}
    cassandra := person{
        lastName:  "Alexandra", 
        firstName: "Cassandra",
        contact: contactInfo{
            email: "cass@gmail.com",
            phone: 987654321,
        },
        address: address{
            country: "Greece",
            city:    "Athens",
        },
    }
    
    // embedded structs for alice should have been initialized with defaults
    if alice.firstName == "" || alice.lastName == "" {
        alice.firstName = "alice"
        alice.lastName = "wonderland"
        alice.contact.email = "aliceinwonderland@gmail.com"
        alice.address = address{"Wonderland", "Tea Garden"}
    }
}

/*
 * Structs are "pass by value" (i.e. copy of value is created) and the original value is not modified
 * Note that slices, maps, channels, pointers, functions are reference types
 * For reference types, the reference to the value is copied instead of the actual value 
 */
func modifyStruct() {
    var ellie person
    ellie.wrongModifyFirstName("alice")
    
    elliePtr := &ellie
    elliePtr.modifyFirstName("alice")

    // Go gives us a shortcut that auto inferences the "ellie" struct into a pointer
    ellie.modifyFirstName("ellie")
}

func (p person) wrongModifyFirstName(newName string) {
    p.firstName = newName
}

func (p *person) modifyFirstName(newName string) {
    // Go automatically dereferences the pointer for us
    p.firstName = newName
}

// anonymous structs
func initAnonymousStruct() {
    p1 := struct{
        first string 
        last string 
    }{
        first: "Ellie",
        last: "Williams",
    }
}
