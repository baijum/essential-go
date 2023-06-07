Appendix A: Answers
===================

Chapter 2: Quickstart
---------------------

**Problem 1:** Write a function to check whether the first letter in a
given string is capital letters in English (A,B,C,D etc).

**Solution:**

::

   package main

   import "fmt"

   func StartsCapital(s string) bool {
       for _, v := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
           if string(s[0]) == string(v) {
               return true
           }
       }
       return false
   }

   func main() {
       h := StartsCapital("Hello")
       fmt.Println(h)
       w := StartsCapital("world")
       fmt.Println(w)
   }

**Problem 2:** Write a function to generate Fibonacci numbers below a
given value.

**Solution:**

::

   package main

   import "fmt"

   func Fib(n int) {
       for i, j := 0, 1; i < n; i, j = i+j, i {
           fmt.Println(i)
       }

   }

   func main() {
       Fib(200)
   }

Chapter 3: Control Structures
-----------------------------

**Problem 1:** Write a program to print greetings based on time.
Possible greetings are Good morning, Good afternoon and Good evening.

**Solution:**

::

   package main

   import (
       "fmt"
       "time"
   )

   func main() {
       t := time.Now()
       switch {
       case t.Hour() < 12:
           fmt.Println("Good morning!")
       case t.Hour() < 17:
           fmt.Println("Good afternoon.")
       default:
           fmt.Println("Good evening.")
       }
   }

**Problem 2:** Write a program to check if a given number is a multiple
of 2, 3, or 5.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import (
   	"fmt"
   )

   func main() {
   	n := 7
   	found := false
   	for _, j := range []int{2, 3, 5} {
   		if n%j == 0 {
   			fmt.Printf("%d is a multiple of %d\n", n, j)
   			found = true
   		}
   	}
   	if !found {
   		fmt.Printf("%d is not a multiple of 2, 3, or 5\n", n)
   	}
   }

Chapter 4: Data Structures
--------------------------

**Problem 1:** Write a program to record temperatures in different
locations and functionality to check whether its freezing or not.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   type Temperature float64

   func (t Temperature) Freezing() bool {
   	if t < Temperature(0.0) {
   		return true
   	}
   	return false
   }

   func main() {

   	temperatures := map[string]Temperature{
   		"New York":  9.3,
   		"London":    13.5,
   		"New Delhi": 31.5,
   		"Montreal":  -9.0,
   	}

   	location := "New Delhi"
   	fmt.Println(location, temperatures[location].Freezing())

   	location = "Montreal"
   	fmt.Println(location, temperatures[location].Freezing())

   }

**Problem 2:** Create a map of world nations and details. They key could
be the country name and value could be an object with details including
capital, currency, and population.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   type Country struct {
   	Capital    string
   	Currency   string
   	Popolation int
   }

   func main() {
   	countries := map[string]Country{}
   	countries["India"] = Country{Capital: "New Delhi",
   		Currency: "Indian Rupee", Popolation: 1428600000}
   	fmt.Printf("%#v\n", countries)
   }

Chapter 5: Functions & Methods
------------------------------

**Problem 1:** Write a program with a function to calculate perimeter of
a circle.

**Solution:**

::

   package main

   import "fmt"

   type Circle struct {
       Radius float64
   }

   // Area return the area of a circle for the given radius
   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }

   func main() {
       c := Circle{5.0}
       fmt.Println(c.Area())
   }

Chapter 6: Interfaces
---------------------

**Problem 1:** Implement the built-in ``error`` interface for a custom
data type. This is how the ``error`` interface is defined:

::

   type error interface {
       Error() string
   }

**Solution:**

::

   package main

   import "fmt"

   type UnauthorizedError struct {
       UserID string
   }

   func (e UnauthorizedError) Error() string {
       return "User not authorised: " + e.UserID
   }

   func SomeAction() error {
       return UnauthorizedError{"jack"}
   }

   func main() {
       err := SomeAction()
       if err != nil {
           fmt.Println(err)
       }
   }

Chapter 7: Concurrency
----------------------

**Problem 1:** Write a program to watch log files and detect any entry
with a particular word.

**Solution:**

::

   package main

   import (
       "bufio"
       "fmt"
       "os"
       "os/signal"
       "strings"
       "time"
   )

   func watch(word, fp string) error {

       f, err := os.Open(fp)
       if err != nil {
           return err
       }
       r := bufio.NewReader(f)
       defer f.Close()
       for {
           line, err := r.ReadBytes('\n')
           if err != nil {
               if err.Error() == "EOF" {
                   time.Sleep(2 * time.Second)
                   continue
               }
               fmt.Printf("Error: %s\n%v\n", line, err)
           }
           if strings.Contains(string(line), word) {
               fmt.Printf("%s: Matched: %s\n", fp, line)
           }
           time.Sleep(2 * time.Second)
       }
   }

   func main() {
       word := os.Args[1]
       files := []string{}
       for _, f := range os.Args[2:len(os.Args)] {
           files = append(files, f)
           go watch(word, f)
       }
       sig := make(chan os.Signal, 1)
       done := make(chan bool)
       signal.Notify(sig, os.Interrupt)
       go func() {
           for _ = range sig {
               done <- true
           }
       }()
       <-done
   }

Chapter 8: Packages
-------------------

**Problem 1:** Create a package with 3 source files and another *doc.go*
for documentation. The package should provide functions to calculate
areas for circle, rectangle, and triangle.

**Solution:**

circle.go:

::

   package shape

   // Circle represents a circle shape
   type Circle struct {
       Radius float64
   }

   // Area return the area of a circle
   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }

rectangle.go:

::

   package shape

   // Rectangle represents a rectangle shape
   type Rectangle struct {
       Length float64
       Width float64
   }

   // Area return the area of a rectangle
   func (r Rectangle) Area() float64 {
       return r.Length * r.Width
   }

triangle.go:

::

   package shape

   // Triangle represents a rectangle shape
   type Triangle struct {
       Breadth float64
       Height float64
   }

   // Area return the area of a triangle
   func (t Triangle) Area() float64 {
       return (t.Breadth * t.Height)/2
   }

doc.go:

::

   // Package shape provides areas for different shapes
   // This includes circle, rectangle, and triangle.

Chapter 9: Input/Output
-----------------------

**Problem 1:** Write a program to format a complex number as used in
mathematics. Example: ``2 + 5i``

Use a struct like this to define the complex number:

::

   type Complex struct {
       Real float64
       Imaginary float64
   }

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   type Complex struct {
   	Real      float64
   	Imaginary float64
   }

   func (c Complex) String() string {
   	return fmt.Sprintf("%.02f + %.02fi", c.Real, c.Imaginary)
   }

   func main() {
   	c1 := Complex{Real: 2.3, Imaginary: 5}
   	fmt.Println(c1)
   }

Chapter 10: Testing
-------------------

**Problem 1:** Write a program to fail test and not continue with
remaining tests.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "testing"

   func TestHelloWorld(t *testing.T) {
   	t.Errorf("First error and continue")
   	t.Fatalf("Second error and not continue")
   	t.Errorf("Third error does not display")
   }

Chapter 11: Tooling
-------------------

**Problem 1:** Write a program with exported type and methods with
documentation strings. Then print the documentation using ``go doc``
command.

**Solution:**

Here is the package definition for a circle object:

::

   // Package defines a circle object
   package circle

   // Circle represents a circle shape
   type Circle struct {
       Radius float64
   }

   // Area return the area of a circle
   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }

The docs can be accessed like this:

::

   $ go doc
   package circle // import "."

   Package defines a circle object

   type Circle struct{ ... }

   $ go doc  Circle
   type Circle struct {
           Radius float64
   }
       Circle represents a circle shape


   func (c Circle) Area() float64

   $ go doc  Circle.Area
   func (c Circle) Area() float64
       Area return the area of a circle
