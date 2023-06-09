Objects
=======

   *Program to an interface, not an implementation.* – Design Patterns
   by Gang of Four

In the chapter on functions, we have also learned about methods. A
method is a function associated with a type, more specifically a
concrete type. As you know, an object is an instance of a type. In
general, methods define the behavior of an object.

Interfaces in Go provide a formal way to specify the behavior of an
object. In layman’s terms, Interface is like a blueprint which describes
an object. So, Interface is considered as an abstract type, commonly
referred to as interface type.

Small interfaces with one or two methods are common in Go.

Here is a ``Geometry`` interface which defines two methods:

::

   type Geometry interface {
       Area() float64
       Perimeter() float64
   }

If any type satisfy this interface - that is define these two methods
which returns float64 - then, we can say that type is implementing this
interface. One difference with many other languages with interface
support and Go is that, in Go implementing an interface happens
implicitly. So, no need to explicitly declare a type is implementing a
particular interface.

To understand this idea, consider this ``Rectangle`` struct type:

::

   type Rectangle struct {
       Width float64
       Height float64
   }

   func (r Rectangle) Area() float64 {
       return r.Width * r.Height
   }

   func (r Rectangle) Perimeter() float64 {
       return 2*(r.width + r.height)
   }

As you can see above, the above Rectangle type has two methods named
Area and Perimeter which returns float64. So, we can say Rectangle is
implementing the ``Geometry`` interface. To elaborate the example
further, we will create one more implementation:

::

   type Circle struct {
       Radius float64
   }

   func (c Circle) Area() float64 {
       return 3.14 * c.Radius * c.Radius
   }

   func (c Circle) Perimeter() float64 {
       return 2 * 3.14 * c.Radius
   }

Now we have two separate implementations for the ``Geometry`` interface.
So, if anywhere the ``Geometry`` interface type is expected, you can use
any of these implementations.

Let’s define a function which accepts a ``Geometry`` type and prints
area and perimeter.

::

   func Measure(g Geometry) {
       fmt.Println("Area:", g.Area())
       fmt.Println("Perimeter:", g.Perimeter())
   }

When you call the above function, you can pass the argument as an object
of ``Geometry`` interface type. Since both ``Rectangle`` and ``Circle``
satisfy that interface, you can use either one of them.

Here is a code which call ``Measure`` function with ``Rectangle`` and
``Circle`` objects:

::

   r := Rectangle{Width: 2.5, Height: 4.0}
   c := Circle{Radius: 6.5}
   Measure(r)
   Measure(c)

Type with Multiple Interfaces
-----------------------------

In Go, a type can implement more than one interface. If a type has
methods that satisfy different interfaces, we can say that that type is
implementing those interfaces.

Consider this interface:

::

   type Stringer interface {
        String() string
   }

In previous section, there was a Rectangle type declared with two
methods. In the same package, if you declare one more method like below,
it makes that type implementing Stringer interface in addition to the
Geometry interface.

::

   func (r Rectangle) String() string {
       return fmt.Sprintf("Rectangle %vx%v", r.Width * r.Height)
   }

Now the ``Rectangle`` type conforms to both ``Geometry`` interface and
``Stringer`` interface.

Empty Interface
---------------

The empty interface is the interface type that has no methods. Normally
the empty interface will be used in the literal form: ``interface``. All
types satisfy empty interface. A function which accepts empty interface,
can receive any type as the argument. Here is an example:

::

   func blackHole(v interface{}) {
   }

   blackHole(1)
   blackHole("Hello")
   blackHole(struct{})

In the above code, the ``blackHole`` functions accepts an empty
interface. So, when you are calling the function, any type of argument
can be passed.

The ``Println`` function in the ``fmt`` package is variadic function
which accepts empty interfaces. This is how the function signature looks
like:

::

   func Println(a ...interface{}) (n int, err error) {

Since the ``Println`` accepts empty interfaces, you could pass any type
arguments like this:

::

   fmt.Println(1, "Hello", struct{})

Pointer Receiver
----------------

In the chapter on Functions, you have seen that the methods can use a
pointer receiver. Also we understood that the pointer receivers are
required when the object attributes need be to modified or when passing
large size data.

Consider the implementation of ``Stringer`` interface here:

::

   type Temperature struct {
        Value float64
        Location string
   }

   func (t *Temperature) String() string {
        o := fmt.Sprintf("Temp: %.2f Loc: %s", t.Value, t.Location)
        return o
   }

In the above example, the ``String`` method is implemented using a
pointer receiver. Now if you define a function which accepts the
``fmt.Stringer`` interface, and want the ``Temperature`` object, it
should be a pointer to ``Temperature``.

::

   func cityTemperature(v fmt.Stringer) {
       fmt.Println(v.String())
   }

   func main() {
       v := Temperature{35.6, "Bangalore"}
       cityTemperature(&v)
   }

As you can see, the ``cityTemperature`` function is called with a
pointer. If you modify the above code and pass normal value, you will
get an error. The below code will produce an error as pointer is not
passed.

::

   func main() {
       v := Temperature{35.6, "Bangalore"}
       cityTemperature(v)
   }

The error message will be something like this:

::

   cannot use v (type Temperature) as type fmt.Stringer in argument to
   cityTemperature: Temperature does not implement fmt.Stringer (String
   method has pointer receiver)

Type Assertions
---------------

In some cases, you may want to access the underlying concrete value from
the interface value. Let’s say you define a function which accepts an
interface value and want access attribute of the concrete value.
Consider this example:

::

   type Geometry interface {
       Area() float64
       Perimeter() float64
   }

   type Rectangle struct {
       Width float64
       Height float64
   }

   func (r Rectangle) Area() float64 {
       return r.Width * r.Height
   }

   func (r Rectangle) Perimeter() float64 {
       return 2*(r.width + r.height)
   }

   func Measure(g Geometry) {
       fmt.Println("Area:", g.Area())
       fmt.Println("Perimeter:", g.Perimeter())
   }

In the above example, if you want to print the width and and height from
the ``Measure`` function, you can use type assertions.

Type assertion gives the underlying concrete value of an interface type.
In the above example, you can access the rectangle object like this:

::

   r := g.(Rectangle)
       fmt.Println("Width:", r.Width)
       fmt.Println("Height:", r.Height)

If the assertion fail, it will trigger a panic.

Type assertion has an alternate syntax where it will not panic if
assertion fail, but gives one more return value of boolean type. The
second return value will be ``true`` if assertion succeeds otherwise it
will give ``false``.

::

   r, ok := g.(Rectangle)
       if ok {
           fmt.Println("Width:", r.Width)
           fmt.Println("Height:", r.Height)
       }

If there are many types that need to be asserted like this, Go provides
a type switches which is explained in the next section.

Type Switches
-------------

As you have seen in the previous section, type assertions gives access
to the underlying value. But if there any many assertions need to be
made, there will be lots ``if`` blocks. To avoid this, Go provides type
switches.

::

   switch v := g.(type) {
       case Rectangle:
           fmt.Println("Width:", v.Width)
           fmt.Println("Height:", v.Height)
       case Circle:
           fmt.Println("Width:", v.Radius)
       case default:
           fmt.Println("Unknown:")
       }

In the above example, type assertion is used with switch cases. Based on
the type of ``g``, the case is executed.

Note that the *fallthrough* statement does not work in type switch.

Exercises
---------

**Exercise 1:** Using the ``Marshaller`` interface, make the marshalled
output of the ``Person`` object given here all in upper case.

::

   type Person struct {
       Name  string
       Place string
   }

**Solution:**

.. code-block:: go
   :linenos:

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

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Implement the built-in ``error`` interface for a custom
data type. This is how the ``error`` interface is defined:

::

   type error interface {
       Error() string
   }

Summary
-------

This chapter explained the concept of interfaces and their uses.
Interfaces are an important concept in Go. Understanding interfaces and
using them properly makes the design robust. The chapter covered the
empty interface, pointer receivers, and type assertions and type
switches.

Brief summary of key concepts introduced in this chapter:

-  An interface is a set of methods that a type must implement. A type
   that implements an interface can be used anywhere an interface is
   expected. This allows for greater flexibility and reusability in Go
   code.

-  A pointer receiver is a method that takes a pointer to a struct as
   its receiver. Pointer receivers are often used to modify the state of
   a struct.

-  A type assertion is a way of checking the type of a value at runtime.
   Type assertions can be used to ensure that a value is of a certain
   type before using it.

-  A type switch is a control flow statement that allows for different
   code to be executed based on the type of a value. Type switches can
   be used to make code more robust and easier to read.
