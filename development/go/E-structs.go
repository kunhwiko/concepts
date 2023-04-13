package go

import (
    "fmt"
)

// Go does not support inheritance and focuses more on composition
// Go allows for the creation of new types
type names []string

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
// "address" will be translated as a field name "address" of type address
type person struct {
    firstName string
    lastName  string
    contact   contactInfo
    address
}

// "functions" do not belong to a particular type
func declareType() {
    // declare names type
    var friends names = []string{"Alice", "Bob"}
    friends = names{"Cassandra", "Ellie"}
    friends.modifyFirst()
    
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
        alice.address = address{"Wonderland", "Wonderland"}
    }
}

// "methods" belong to a particular type where this one has a receiver of type 'names' called 'n'
// by convention, the instance of the type is represented as a single letter
func (n names) modifyFirst() {
    n[0] = "Alice"
}

// structs are "pass by value" (i.e. copy is created) and the original reference is not modified
func modifyStruct() {
    var ellie person
    ellie.wrongModifyFirstName("alice")
    elliePtr := &ellie
    elliePtr.modifyFirstName("alice")
}

func (p person) wrongModifyFirstName(newName string) {
    p.firstName = newName
}

func (p *person) modifyFirstName(newName string) {
    (*p).firstName = newName
}
