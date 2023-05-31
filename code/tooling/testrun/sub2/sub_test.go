package sub

import "testing"

func TestSub(t *testing.T) {
    if try() == 1 {
        t.Fail()
    }
}
