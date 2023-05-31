package hello

import "fmt"

// Hello says "Hello" with name
func Hello(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
