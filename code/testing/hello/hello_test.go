package hello

import "testing"

func TestHello(t *testing.T) {
    out := Hello("Tom")
    if out != "Hello, Tom!" {
        t.Fail()
    }
}
