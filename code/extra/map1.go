package main

import "fmt"

func main() {
    var m1 map[string]int
    m1 = make(map[string]int) // comment and run
    m1["Mango"] = 1
    fmt.Printf("m1: %#v\n", m1)

    m2 := make(map[string]int)
    m2["Mango"] = 2
    fmt.Printf("m2: %#v\n", m2)

    m3 := map[string]int{}
    //m3 := map[string]int{"Apple": 5}
    m3["Mango"] = 3
    fmt.Printf("m3: %#v\n", m3)

    fmt.Printf("len(m3): %#v\n", len(m3))

    v3 := m3["Mango"]
    fmt.Printf("v3: %#v\n", v3)

    i3 := m3["Apple"]
    fmt.Printf("i3: %#v\n", i3)

    m3["Orange"] = 3
    fmt.Printf("m3: %#v\n", m3)
    delete(m3, "Orange")
    fmt.Printf("m3: %#v\n", m3)

    fmt.Printf("m3: %#v\n", m3)
    delete(m3, "Non-Existing")
    fmt.Printf("m3: %#v\n", m3)

    j31, ok31 := m3["Non-Existing"]
    fmt.Printf("j31: %#v ok31: %#v\n", j31, ok31)

    j32, ok32 := m3["Mango"]
    fmt.Printf("j32: %#v ok32: %#v\n", j32, ok32)

    m3["Orange"] = 3
    m3["Apple"] = 3
    fmt.Printf("m3: %#v\n", m3)

    for key, value := range m3 {
        fmt.Println("Key:", key, "Value:", value)
    }
}
