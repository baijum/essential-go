package sum

import "testing"

func TestAdd(t *testing.T) {
    out := Add(2, 3)
    if out != 5 {
        t.Error("Sum of 2 and 3:", out)
    }
}
