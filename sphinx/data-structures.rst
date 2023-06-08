Data Structures
===============

   *Bad programmers worry about the code. Good programmers worry about
   data structures and their relationships.* — Linus Torvalds

In the last last few chapters we have seen some of the primitive data
data types. We also introduced few other advanced data types without
going to the details. In this chapter, we are going to look into more
data structures.

Primitive Data Types
--------------------

Advanced data structures are built on top of primitive data types. This
section is going to cover the primitive data types in Go.

Zero Value
~~~~~~~~~~

In the Quick Start chapter, you have learned various ways to declare a
variable. When you declare a variable using the ``var`` statement
without assigning a value, a default Zero value will be assigned for
certain types. The Zero value is 0 for integers and floats, empty string
for strings, and false for boolean. To demonstrate this, here is an
example:

::

   package main

   import "fmt"

   func main() {
       var name string
       var age int
       var tall bool
       var weight float64
       fmt.Printf("%#v, %#v, %#v, %#v\n", name, age, tall, weight)
   }

This is the output:

::

   "", 0, false, 0

Variable
~~~~~~~~

In the quick start chapter, we have discussed about variables and its
usage. The variable declared outside the function (package level) can
access anywhere within the same package.

Here is an example:

::

   package main

   import (
       "fmt"
   )

   var name string
   var country string = "India"

   func main() {
       name = "Jack"
       fmt.Println("Name:", name)
       fmt.Println("Country:", country)
   }

In the above example, the ``name`` and ``country`` are two package level
variables. As we have seen above the ``name`` gets zero value, where as
value for ``country`` variable is explicitly initialized.

If the variable has been defined using the ``:=`` syntax, and the user
wants to change the value of that variable, they need to use ``=``
instead of ``:=`` syntax.

If you run the below program, it’s going to throw an error:

::

   package main

   import (
       "fmt"
   )

   func main() {
       age := 25
       age := 35
       fmt.Println(age)
   }

::

   $ go run update.go
   # command-line-arguments
   ./update.go:9:6: no new variables on left side of :=

The above can be fixed like this:

::

   package main

   import (
       "fmt"
   )

   func main() {
       age := 25
       age = 35
       fmt.Println(age)
   }

Now you should see the output:

::

   $ go run update.go
   35

Using the *reflect* package, you can identify the type of a variable:

::

   package main

   import (
          "fmt"
          "reflect"
   )

   func main() {
        var pi = 3.41
        fmt.Println("type:", reflect.TypeOf(pi))
   }

Using one or two letter variable names inside a function is common
practice. If the variable name is multi-word, use lower camelCase
(initial letter lower and subsequent words capitalized) for unexported
variables. If the variable is an exported one, use upper CamelCase (all
the words capitalized). If the variable name contains any abbreviations
like ID, use capital letters. Here are few examples: pi, w, r,
ErrorCode, nodeToDaemonPods, DB, InstanceID.

Unused variables and imports
^^^^^^^^^^^^^^^^^^^^^^^^^^^^

If you declare a variable inside a function, use that variable somewhere
in the same function where it is declared. Otherwise, you are going to
get a compile error. Whereas a global variable declared but unused is
not going to throw compile time error.

Any package that is getting imported should find a place to use. Unused
import also throws compile time error.

Boolean Type
~~~~~~~~~~~~

A boolean type represents a pair of truth values. The truth values are
denoted by the constants *true* and *false*. These are the three logical
operators that can be used with boolean values:

-  ``&&`` – Logical AND

-  ``||`` – Logical OR

-  ``!`` – Logical NOT

Here is an example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
   	yes := true
   	no := false
   	fmt.Println(yes && no)
   	fmt.Println(yes || no)
   	fmt.Println(!yes)
   	fmt.Println(!no)
   }

The output of the above logical operators are like this:

::

   $ go run logical.go
   false
   true
   false
   true

Numeric Types
~~~~~~~~~~~~~

The numeric type includes both integer types and floating-point types.
The allowed values of numeric types are same across all the CPU
architectures.

These are the unsigned integers:

-  uint8 – the set of all unsigned 8-bit integers (0 to 255)

-  uint16 – the set of all unsigned 16-bit integers (0 to 65535)

-  uint32 – the set of all unsigned 32-bit integers (0 to 4294967295)

-  uint64 – the set of all unsigned 64-bit integers (0 to
   18446744073709551615)

These are the signed integers:

-  int8 – the set of all signed 8-bit integers (-128 to 127)

-  int16 – the set of all signed 16-bit integers (-32768 to 32767)

-  int32 – the set of all signed 32-bit integers (-2147483648 to
   2147483647)

-  int64 – the set of all signed 64-bit integers (-9223372036854775808
   to 9223372036854775807)

These are the two floating-point numbers:

-  float32 – the set of all IEEE-754 32-bit floating-point numbers

-  float64 – the set of all IEEE-754 64-bit floating-point numbers

These are the two complex numbers:

-  complex64 – the set of all complex numbers with float32 real and
   imaginary parts

-  complex128 – the set of all complex numbers with float64 real and
   imaginary parts

These are the two commonly used used aliases:

-  byte – alias for uint8

-  rune – alias for int32

String Type
~~~~~~~~~~~

A string type is another most import primitive data type. String type
represents string values.

Constants
---------

A constant is an unchanging value. Constants are declared like
variables, but with the *const* keyword. Constants can be character,
string, boolean, or numeric values. Constants cannot be declared using
the ``:=`` syntax.

In Go, *const* is a keyword introducing a name for a scalar value such
as 2 or 3.14159 or "scrumptious". Such values, named or otherwise, are
called constants in Go. Constants can also be created by expressions
built from constants, such as 2+3 or 2+3i or math.Pi/2 or ("go"+"pher").
Constants can be declared are at package level or function level.

This is how to declare constants:

::

   package main

   import (
       "fmt"
   )

   const Freezing = true
   const Pi = 3.14
   const Name = "Tom"

   func main() {
       fmt.Println(Pi, Freezing, Name)
   }

You can also use the factored style declaration:

::

   package main

   import (
       "fmt"
   )

   const (
         Freezing = true
         Pi = 3.14
         Name = "Tom"
   )

   func main() {
       fmt.Println(Pi, Freezing, Name)
   }

   const (
         Freezing = true
         Pi = 3.14
         Name = "Tom"
   )

Compiler throws an error if the constant is tried to assign a new value:

::

   package main

   import (
       "fmt"
   )

   func main() {
       const Pi = 3.14
       Pi = 6.86
       fmt.Println(Pi)
   }

The above program throws an error like this:

::

   $ go run constants.go
   constants:9:5: cannot assign to Pi

iota
~~~~

The *iota* keyword is used to define constants of incrementing numbers.
This simplify defining many constants. The values of iota is reset to
*0* whenever the reserved word const appears. The value increments by
one after each line.

Consider this example:

::

   // Token represents a lexical token.
   type Token int

   const (
       // Illegal represents an illegal/invalid character
       Illegal Token = iota

       // Whitespace represents a white space
       // (" ", \t, \r, \n) character
       Whitespace

       // EOF represents end of file
       EOF

       // MarkerID represents '\id' or '\id1' marker
       MarkerID

       // MarkerIde represents '\ide' marker
       MarkerIde
   )

In the above example, the ``Token`` is custom type defined using the
primitive *int* type. The constants are defined using the factored
syntax (many constants within parenthesis). There are comments for each
constant values. Each constant value is be incremented starting from
``0``. In the above example, ``Illegal`` is ``0``, Whitespace is ``1``,
``EOF`` is ``2`` and so on.

The ``iota`` can be used with expressions. The expression will be
repeated. Here is a good example taken from *Effective Go*
(https://go.dev/doc/effective_go#constants):

::

   type ByteSize float64

   const (
       // ignore first value (0) by assigning to blank identifier
       _           = iota
       KB ByteSize = 1 << (10 * iota)
       MB
       GB
       TB
       PB
       EB
       ZB
       YB
   )

Using ``_`` (blank identifier) you can ignore a value, but *iota*
increments the value. This can be used to skip certain values. As you
can see in the above example, you can use an expression with *iota*.

Iota is reset to ``0`` whenever the *const* keyword appears in the
source code. This means that if you have multiple *const* declarations
in a single file, *iota* will start at ``0`` for each declaration. Iota
can only be used in *const* declarations. It cannot be used in other
types of declarations, such as *var* declarations. The value of *iota*
is only available within the const declaration in which it is used. It
cannot be used outside of that declaration.

Blank Identifier
^^^^^^^^^^^^^^^^

Sometimes you may need to ignore the value returned by a function. Go
provides a special identifier called blank identifier to ignore any
types of values. In Go, underscore ``_`` is the blank identifier.

Here is an example usage of blank identifier where the second value
returned by the function is discarded.

::

   x, _ := someFunc()

Blank identifier can be used as import alias to invoke init function
without using the package.

::

   import (
          "database/sql"

          _ "github.com/lib/pq"
   )

In the above example, the ``pq`` package has some code which need to be
invoked to initialize the database driver provided by that package. And
the exported functions within the above package is supposed to be not
used.

We have already seen another example where blank identifier if used with
*iota* to ignore certain constants.

.. _`sec:arrays`:

Arrays
------

An array is an ordered container type with a fixed number of data. In
fact, the arrays are the foundation where slice is built. We will study
about slices in the next section. Most of the time, you can use slice
instead of an array.

The number of values in the array is called the length of that array.
The array type ``[n]T`` is an array of ``n`` values of type ``T``. Here
are two example arrays:

::

   colors := [3]string{"Red", "Green", "Blue"}
   heights := [4]int{153, 146, 167, 170}

In the above example, the length of first array is ``3`` and the array
values are string data. The second array contains ``int`` values. An
array’s length is part of its type, so arrays cannot be re-sized. So, if
the length is different for two arrays, those are distinct incompatible
types. The built-in ``len`` function gives the length of array.

Array values can be accessed using the index syntax, so the expression
``s[n]`` accesses the nth element, starting from zero.

An array values can be read like this:

::

   colors := [3]string{"Red", "Green", "Blue"}
   i := colors[1]
   fmt.Println(i)

Similarly array values can be set using index syntax. Here is an
example:

::

   colors := [3]string{"Red", "Green", "Blue"}
   colors[1] = "Yellow"

Arrays need not be initialized explicitly. The zero value of an array is
a usable array with all elements zeroed.

::

   var colors [3]string
   colors[1] = "Yellow"

In this example, the values of colors will be empty strings (zero
value). Later we can assign values using the index syntax.

There is a way to declare array literal without specifying the length.
When using this syntax variant, the compiler will count and set the
array length.

::

   colors := [...]string{"Red", "Green", "Blue"}

In the chapter on control structures, we have seen how to use For loop
for iterating over slices. In the same way, you can iterate over array.

Consider this complete example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       colors := [3]string{"Red", "Green", "Blue"}
       fmt.Println("Length:", len(colors))
       for i, v := range colors {
           fmt.Println(i, v)
       }
   }

If you save the above program in a file named ``colors.go`` and run it,
you will get output like this:

::

   $ go run colors.go
   Length: 3
   0 Red
   1 Green
   2 Blue

In the above program, a string array is declared and initialized with
three string values. In the 7th line, the length is printed and it gives
3. The ``range`` clause gives index and value, where the index starts
from zero.

Slices
------

Slice is one of most important data structure in Go. Slice is more
flexible than an array. It is possible to add and remove values from a
slice. There will be a length for slice at any time. Though the length
vary dynamically as the content value increase or decrease.

The number of values in the slice is called the length of that slice.
The slice type ``[]T`` is a slice of type ``T``. Here are two example
slices:

::

   colors := []string{"Red", "Green", "Blue"}
   heights := []int{153, 146, 167, 170}

The first one is a slice of strings and the second slice is a slice of
integers. The syntax is similar to array except the length of slice is
not explicitly specified. You can use built-in ``len`` function to see
the length of slice.

Slice values can be accessed using the index syntax, so the expression
``s[n]`` accesses the nth element, starting from zero.

A slice values can be read like this:

::

   colors := []string{"Red", "Green", "Blue"}
   i := colors[1]
   fmt.Println(i)

Similary slice values can be set using index syntax. Here is an example:

::

   colors := []string{"Red", "Green", "Blue"}
   colors[1] = "Yellow"

Slices should be initialized with a length more than zero to access or
set values. In the above examples, we used slice literal syntax for
that. If you define a slice using ``var`` statement without providing
default values, the slice will be have a special zero value called
``nil``.

Consider this complete example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       var v []string
       fmt.Printf("%#v, %#v\n", v, v == nil)
       // Output: []string(nil), true
   }

In the above example, the value of slice ``v`` is ``nil``. Since the
slice is nil, values cannot be accessed or set using the index. These
operations are going to raise runtime error (index out of range).

Sometimes it may not be possible to initialize a slice with some value
using the literal slice syntax given above. Go provides a built-in
function named ``make`` to initialize a slice with a given length and
zero values for all items. For example, if you want a slice with 3
items, the syntax is like this:

::

   colors := make([]string, 3)

In the above example, a slice will be initialized with 3 empty strings
as the items. Now it is possible to set and get values using the index
as given below:

::

   colors[0] = "Red"
   colors[1] = "Green"
   colors[2] = "Blue"
   i := colors[1]
   fmt.Println(i)

If you try to set value at 3rd index (``colors[3]``), it’s going to
raise runtime error with a message like this: "index out of range". Go
has a built-in function named ``append`` to add additional values. The
append function will increase the length of the slice.

Consider this example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       v := make([]string, 3)
       fmt.Printf("%v\n", len(v))
       v = append(v, "Yellow")
       fmt.Printf("%v\n", len(v))
   }

In the above example, the slice length is increased by one after append.
It is possible to add more values using ``append``. See this example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       v := make([]string, 3)
       fmt.Printf("%v\n", len(v))
       v = append(v, "Yellow", "Black")
       fmt.Printf("%v\n", len(v))
   }

The above example append two values. Though you can provide any number
of values to append.

You can use the "..." operator to expand a slice. This can be used to
append one slice to another slice. See this example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       v := make([]string, 3)
       fmt.Printf("%v\n", len(v))
       a := []string{"Yellow", "Black"}
       v = append(v, a...)
       fmt.Printf("%v\n", len(v))
   }

In the above example, the first slice is appended by all items in
another slice.

Slice Append Optimization
~~~~~~~~~~~~~~~~~~~~~~~~~

If you append too many values to a slice using a for loop, there is one
optimization related that you need to be aware.

Consider this example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       v := make([]string, 0)
       for i := 0; i < 9000000; i++ {
           v = append(v, "Yellow")
       }
       fmt.Printf("%v\n", len(v))
   }

If you run the above program, it’s going to take few seconds to execute.
To explain this, some understanding of internal structure of slice is
required. Slice is implemented as a struct and an array within. The
elements in the slice will be stored in the underlying array. As you
know, the length of array is part of the array type. So, when appending
an item to a slice the a new array will be created. To optimize, the
``append`` function actually created an array with double length.

In the above example, the underlying array must be changed many times.
This is the reason why it’s taking few seconds to execute. The length of
underlying array is called the capacity of the slice. Go provides a way
to initialize the underlying array with a particular length. The
``make`` function has a fourth argument to specify the capacity.

In the above example, you can specify the capacity like this:

::

   v := make([]string, 0, 9000000)

If you make this change and run the program again, you can see that it
run much faster than the earlier code. The reason for faster code is
that the slice capacity had already set with maximum required length.

Maps
----

Map is another important data structure in Go. We have briefly discussed
about maps in the Quick Start chapter. As you know, map is an
implementation of hash table. The hash table is available in many very
high level languages. The data in map is organized like key value pairs.

A variable of map can be declared like this:

::

   var fruits map[string]int

To make use that variable, it needs to be initialized using *make*
function.

::

   fruits = make(map[string]int)

You can also initialize using the *:=* syntax:

::

   fruits := map[string]int{}

or with *var* keywod:

::

   var fruits = map[string]int{}

You can initialize map with values like this:

::

   var fruits = map[string]int{
         "Apple":  45,
         "Mango":  24,
         "Orange": 34,
     }

After initializing, you can add new key value pairs like this:

::

   fruits["Grape"] = 15

If you try to add values to maps without initializing, you will get an
error like this:

::

   panic: assignment to entry in nil map

Here is an example that’s going to produce panic error:

::

   package main

   func main() {
       var m map[string]int
       m["k"] = 7
   }

To access a value corresponding to a key, you can use this syntax:

::

   mangoCount := fruits["Mango"]

Here is an example:

::

   package main

   import "fmt"

   func main() {
     var fruits = map[string]int{
         "Apple":  45,
         "Mango":  24,
         "Orange": 34,
     }
     fmt.Println(fruits["Apple"])
   }

If the key doesn’t exist, a zero value will be returned. For example, in
the below example, value of ``pineappleCount`` is going be ``0``.

::

   package main

   import "fmt"

   func main() {
     var fruits = map[string]int{
         "Apple":  45,
         "Mango":  24,
         "Orange": 34,
     }
     pineappleCount := fruits["Pineapple"]
     fmt.Println(pineappleCount)
   }

If you need to check if the key exist, the above syntax can be modified
to return two values. The first one would be the actual value or zero
value and the second one would be a boolean indicating if the key
exists.

::

   package main

   import "fmt"

   func main() {
     var fruits = map[string]int{
         "Apple":  45,
         "Mango":  24,
         "Orange": 34,
     }
     _, ok := fruits["Pineapple"]
     fmt.Println(ok)
   }

In the above program, the first returned value is ignored using a blank
identifier. And the value of ``ok`` variable is ``false``. Normally, you
can use if condition to check if the key exists like this:

::

   package main

   import "fmt"

   func main() {
       var fruits = map[string]int{
           "Apple":  45,
           "Mango":  24,
           "Orange": 34,
       }
       if _, ok := fruits["Pineapple"]; ok {
           fmt.Println("Key exists.")
       } else {
           fmt.Println("Key doesn't exist.")
       }
   }

To see the number of key/value pairs, you can use the built-in *len*
function.

::

   package main

   import "fmt"

   func main() {
       var fruits = map[string]int{
           "Apple":  45,
           "Mango":  24,
           "Orange": 34,
       }
       fmt.Println(len(fruits))
   }

The above program should print ``3`` as the number of items in the map.

To remove an item from the map, use the built-in *delete* function.

::

   package main

   import "fmt"

   func main() {
       var fruits = map[string]int{
           "Apple":  45,
           "Mango":  24,
           "Orange": 34,
       }
       delete(fruits, "Orange")
       fmt.Println(len(fruits))
   }

The above program should print ``2`` as the number of items in the map
after deleting one item.

Custom Data Types
-----------------

Apart from the built-in data types, you can create your own custom data
types. The type keyword can be used to create custom types. Here is an
example.

::

   package main

   import "fmt"

   type age int

   func main() {
        a := age(2)
        fmt.Println(a)
        fmt.Printf("Type: %T\n", a)
   }

If you run the above program, the output will be like this:

::

   $ go run age.go
   2
   Type: age

Structs
~~~~~~~

Struct is a composite type with multiple fields of different types
within the struct. For example, if you want to represent a person with
name and age, the ``struct`` type will be helpful. The ``Person`` struct
definition will look like this:

::

   type Person struct {
       Name string
       Age  int
   }

As you can see above, the ``Person`` struct is defined using ``type``
and ``struct`` keywords. Within the curly brace, attributes with other
types are defined. If you avoid attributes, it will become an empty
struct.

Here is an example empty struct:

::

   type Empty struct {
   }

Alternatively, the curly brace can be in the same line.

::

   type Empty struct {}

A struct can be initialized various ways. Using a ``var`` statement:

::

   var p1 Person
   p1.Name = "Tom"
   p1.Age = 10

You can give a literal form with all attribute values:

::

   p2 := Person{"Polly", 50}

You can also use named attributes. In the case of named attributes, if
you miss any values, the default zero value will be initialized.

::

   p3 := Person{Name: "Huck"}
   p4 := Person{Age: 10}

In the next, we are going to learn about funtions and methods. That
chapter expands the discussion about custom types bevior changes through
funtions associated with custom types called strcuts.

It is possible to embed structs inside other structs. Here is an
example:

::

   type Person struct {
        Name string
   }

   type Member struct {
        Person
        ID int
   }

Pointers
--------

When you are passing a variable as an argument to a function, Go creates
a copy of the value and send it. In some situations, creating a copy
will be expensive when the size of object is large. Another scenario
where pass by value is not feasible is when you need to modify the
original object inside the function. In the case of pass by value, you
can modify it as you are getting a new object every time. Go supports
another way to pass a reference to the original value using the memory
location or address of the object.

To get address of a variable, you can use ``&`` as a prefix for the
variable. Here in an example usage:

::

   a := 7
   fmt.Printf("%v\n", &a)

To get the value back from the address, you can use ``*`` as a prefix
for the variable. Here in an example usage:

::

   a := 7
   b := &a
   fmt.Printf("%v\n", *b)

Here is a complete example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func value(a int) {
       fmt.Printf("%v\n", &a)
   }

   func pointer(a *int) {
       fmt.Printf("%v\n", a)
   }

   func main() {
       a := 4
       fmt.Printf("%v\n", &a)
       value(a)
       pointer(&a)
   }

A typical output will be like this:

::

   0xc42000a340
   0xc42000a348
   0xc42000a340

As you can see above, the second output is different from the first and
third. This is because a value is passed instead of a pointer. And so
when we are printing the address, it’s printing the address of the new
variable.

In the functions chapter, the section about methods
(section `[sec:methods] <#sec:methods>`__) explains the pointer
receiver.

new
~~~

The built-in function *new* can be used to allocate memory. It allocates
zero values and returns the address of the given data type.

Here is an example:

::

   name := new(string)

In this example, a *string* pointer value is allocated with zero value,
in this case empty string, and assigned to a variable.

This above example is same as this:

::

   var name *string
   name = new(string)

In this one *string* pointer variable is declared, but it’s not
allocated with zeror value. It will have *nil* value and so it cannot be
dereferenced. If you try to reference, without allocating using the
*new* function, you will get an error like this:

::

   panic: runtime error: invalid memory address or nil pointer
   dereference

Here is another example using a custom type defined using a primitive
type:

::

   type Temperature float64
   name := new(Temperature)

Exercises
---------

**Exercise 1:** Create a custom type for circle using float64 and define
``Area`` and ``Perimeter``.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   type Circle float64

   func (r Circle) Area() float64 {
   	return float64(3.14 * r * r)
   }

   func (r Circle) Perimeter() float64 {
   	return float64(2 * 3.14 * r)
   }

   func main() {
   	c := Circle(4.0)
   	fmt.Println(c.Area())
   	fmt.Println(c.Perimeter())
   }

The custom ``Circle`` type is created using the built-in ``float64``
type. It would be better if the circle is defined using a struct. Using
struct helps to change the structure later with additional attributes.
The struct will look like this:

::

   type Circle struct {
       Radius float64
   }

**Exercise 2:** Create a slice of structs where the struct represents a
person with name and age.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   type Person struct {
   	Name string
   	Age int
   }

   func main() {
   	persons := []Person{
   		Person{Name: "Huck", Age: 11},
   		Person{Name: "Tom", Age: 10},
   		Person{Name: "Polly", Age: 52},
   	}
   	fmt.Println(persons)
   }

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a program to record temperatures in different
locations and functionality to check whether its freezing or not.

**Problem 2:** Create a map of world nations and details. They key could
be the country name and value could be an object with details including
capital, currency, and population.

Summary
-------

This chapter introduced various data structures in Go. These data
structures will be helpful for organizing data. The chapter started with
a section about zero values. Then constants explained in detail
including *iota* keyword for defining incrementing constants. Next, the
chapter briefly explained about arrays. Then slices, the more useful
data structure built on top of array is explained. Then we looked at how
to define custom data types using existing primitive types. The struct
was introduced which is more useful to create custom data types. Pointer
is also covered. The next chapter will explain more about functions and
methods.
