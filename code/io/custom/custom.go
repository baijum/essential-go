package main

import (
    "fmt"
    "strconv"
)

// Temperature repesent air temperature
type Temperature struct {
    Value float64
    Unit  string
}

func (t Temperature) String() string {
    f := strconv.FormatFloat(t.Value, 'f', 2, 64)
    return f + " degree " + t.Unit
}

func main() {
    temp := Temperature{30.456, "Celsius"}
    fmt.Println(temp)
    fmt.Printf("%v\n", temp)
    fmt.Printf("%+v\n", temp)
    fmt.Printf("%#v\n", temp)
}
