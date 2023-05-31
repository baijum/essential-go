package hello

import "testing"

func TestSayHello(t *testing.T) {
     out := SayHello()
     if out != "Hello, world!" {
        t.Error("Incorrect message", out)
     }
}
