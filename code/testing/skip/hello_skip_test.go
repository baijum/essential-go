package hello

import (
    "runtime"
    "testing"
)

func TestHello(t *testing.T) {
    if runtime.GOOS == "linux" {
        t.SkipNow()
    }
    out := Hello("Tom")
    if out != "Hello, Tom!" {
        t.Fail()
    }
}
