package main

import (
    "encoding/json"
    "fmt"
    "strings"
)

// Person represents a person
type Person struct {
    Name  string
    Place string
}

// MarshalJSON implements the Marshaller interface
func (p Person) MarshalJSON() ([]byte, error) {
    name := strings.ToUpper(p.Name)
    place := strings.ToUpper(p.Place)
    s := []byte(`{"NAME":"` + name + `","PLACE":"` + place + `"}`)
    return s, nil
}

func main() {
    p := Person{Name: "Baiju", Place: "Bangalore"}
    o, err := json.Marshal(p)
    if err != nil {
        panic(err)
    }
    fmt.Println(string(o))
}
