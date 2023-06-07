Input/Output
============

   *for out of the abundance of the heart the mouth speaketh.* – Bible

Users interact with software systems through various input/output
mechanisms. Some of the commonly used mechanisms are these:

-  web browser for web applications using various controllers/widgets

-  mobile for mobile applications

-  shell for command line applications

-  desktop application with native controllers/widgets

To provide these kind of user interactions, you will be required to use
specialized libraries. If the standard library doesn’t provide what you
are looking for, you may need to use third party libraries. This chapter
cover basic mechanisms provided at the language level which you can use
for input/output.

We have already seen some of the basic input/output techniques in the
last few chapters. This chapter will go though more input/output
mechanisms available in Go.

Command Line Arguments
----------------------

The command line arguments as user input and console for output is the
way command line programs are designed. Sometimes output will be other
files and devices. You can access all the command line arguments using
the ``Args`` array/slice attribute available in ``os`` package. Here is
an example:

::

   package main

   import (
       "fmt"
       "os"
   )

   func main() {
       fmt.Println(os.Args)
       fmt.Println(os.Args[0])
       fmt.Println(os.Args[1])
       fmt.Println(os.Args[2])
   }

You can run this program with minimum two additional arguments:

::

   $ go build cmdline.go
   $ ./cmdline one -two
   [cmdline one -two]
   ./cmdline
   one
   -two

As you can see, it’s difficult to parse command line arguments like
this. Go standard library has a package named *flag* which helps to
easily parse command line arguments. This chapter has a section to
explain the *flag* package.

Files
-----

Reading and writing data to files is a common I/O operation in computer
programming. You can manipulate files using the *os* and *io* packages.
It works with both text files and binary files. For the simplicity of
this section, all the examples given here works with text files.

Consider there exists a file named ``poem.txt`` with this text:

.. code:: text

   I wandered lonely as a cloud
   That floats on high o'er vales and hills,
   When all at once I saw a crowd,
   A host, of golden daffodils;
   Beside the lake, beneath the trees,
   Fluttering and dancing in the breeze.

Here is a program to read the the whole file and print:

.. code-block:: go
   :linenos:

   package main

   import (
   	"fmt"
   	"io"
   	"os"
   )

   func main() {
   	fd, err := os.Open("poem.txt")
   	if err != nil {
   		fmt.Println("Error reading file:", err)
   	}
   	for {
   		chars := make([]byte, 50)
   		n, err := fd.Read(chars)
   		if n == 0 && err == io.EOF {
   			break
   		}
   		fmt.Print(string(chars))
   	}
   	fd.Close()
   }

When you run this program, you will get the whole text as output. In the
line number 10, the ``Open`` function from the ``os`` package is called
to open the file for reading. It returns a file descriptor [1]_ and
error. In the line number 14, an infinite loop is stared to read the
content. Line 15 initialize a slice of bytes of length 50. The
``fd.Read`` method reads the given length of characters and writes to
the given slice. It returns the number of characters read and error. The
``io.EOF`` error is returned when end of file is reached. This is used
as the condition to break the loop.

Here is a program to write some text to a file:

.. code-block:: go
   :linenos:

   package main

   import (
   	"fmt"
   	"os"
   )

   func main() {
   	fd, err := os.Create("hello.txt")
   	if err != nil {
   		fmt.Println("Cannot write file:", err)
   		os.Exit(1)
   	}
   	fd.Write([]byte("Hello, World!\n"))
   	fd.Close()
   }

In th line number 9, the *Create* function from the *os* package is
called open the file for writing. It returns a file descriptor and
error. In the line number 14, the *Write* method is give a slice of
bytes to write. After running the program you can see the text in the
``hello.txt`` file.

::

   $ go run writefile.go
   $ cat hello.txt
   Hello, World!

Standard Streams
----------------

Standard streams [2]_ are input and output communication channels
between a computer program and its environment. The three input/output
connections are called standard input (stdin), standard output (stdout)
and standard error (stderr).

Stdin, Stdout, and Stderr are open files pointing to the standard input,
standard output, and standard error file descriptors.

The *fmt* package has functions to read values interactively.

Here is an example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       var name string
       fmt.Printf("Enter your name: ")
       fmt.Scanf("%s", &name)
       fmt.Println("Your name:", name)
   }

The *Scanf* function reads the standard input. The first argument is the
format and the second one is the pointer variable. The value read from
standard input cab be accessed using the given variable.

You can run the above program in different ways:

::

   $ go run code/io/readprint.go
   Enter your name: Baiju
   Your name: Baiju
   $ echo "Baiju" |go run code/io/readprint.go
   Enter your name: Your name: Baiju
   $ go run code/io/readprint.go << EOF
   > Baiju
   > EOF
   Enter your name: Your name: Baiju
   $ echo "Baiju" > /tmp/baiju.txt
   $ go run code/io/readprint.go < /tmp/baiju.txt
   Enter your name: Your name: Baiju

As you can see from this program, the *Printf* function writes to
standard output and the *Scanf* reads the standard input. Go can also
writes to standard error output stream.

The *io* package provides a set of interfaces and functions that allow
developers to work with different types of input and output streams.

Consider a use case to convert everything that comes to standard input
to convert to upper case. This can be achieved by reading all standard
input using ``io.ReadAll`` and converting to upper case. Here is code:

.. code-block:: go
   :linenos:

   package main

   import (
   	"fmt"
   	"io"
   	"os"
   	"strings"
   )

   func main() {
   	stdin, err := io.ReadAll(os.Stdin)
   	if err != nil {
   		panic(err)
   	}
   	str := string(stdin)
   	newStr := strings.TrimSuffix(str, "\n")
   	upper := strings.ToUpper(newStr)
   	fmt.Println(upper)
   }

You can run this program similar to how you did with the previous
program.

You can use *fmt.Fprintf* with *os.Stderr* as the first argument to
write to standard error.

::

   fmt.Fprintf(os.Stderr, "This goes to standard error: %s", "OK")

Alternatively, you can call *WriteString* method of *os.Stderr*:

::

   os.Stderr.WriteString("This goes to standard error")

Using flag Package
------------------

As you have noticed before ``os.Args`` attribute in the *os* package
provides access to all command line arguments. The *flag* package
provides an easy way to parse command line arguments.

You can define string, boolean, and integer flags among others using the
*flag* package..

Here is an integer flag declaration:

::

   var pageCount = flag.Int("count", 240, "number of pages")

The above code snippet defines an integer flag with name as ``count``
and it is stored in a variable with the name as ``pageCount``. The type
of the variable is *\*int*. Similar to this integer flag, you can
defines flags of other types.

Once all the flags are defined, you can parse it like this:

::

   flag.Parse()

The above ``Parse`` function call parse the command line arguments and
store the values in the given variables.

Once the flags are parsed, you can dereference it like this:

::

   fmt.Println("pageCount: ", *pageCount)

To access non-flag arguments:

::

   flag.Args()

The above call returns a the arguments as a slice of strings. It
contains arguments not parsed as flags.

Cobra is a third party package providing a simple interface to create
command line interfaces. Cobra also helps to generate applications and
command files. Many of the most widely used Go projects are built using
Cobra. This is the Cobra website: https://github.com/spf13/cobra

String Formatting
-----------------

Go supports many string format options. To get the default format of any
value, you can use ``%v`` as the format string. Here is an example which
print formatted values using ``%v``:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       fmt.Printf("Value: %v, Type: %T\n", "Baiju", "Baiju")
       fmt.Printf("Value: %v, Type: %T\n", 7, 7)
       fmt.Printf("Value: %v, Type: %T\n", uint(7), uint(7))
       fmt.Printf("Value: %v, Type: %T\n", int8(7), int8(7))
       fmt.Printf("Value: %v, Type: %T\n", true, true)
       fmt.Printf("Value: %v, Type: %T\n", 7.0, 7.0)
       fmt.Printf("Value: %v, Type: %T\n", (1 + 6i), (1 + 6i))
   }

The ``%T`` shows the type of the value. The output of the above program
will be like this.

::

   Value: Baiju, Type: string
   Value: 7, Type: int
   Value: 7, Type: uint
   Value: 7, Type: int8
   Value: true, Type: bool
   Value: 7, Type: float64
   Value: (1+6i), Type: complex128

If you use a ``%+v`` as the format string for struct it shows the field
names. See this example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   // Circle represents a circle
   type Circle struct {
       radius float64
       color  string
   }

   func main() {
       c := Circle{radius: 76.45, color: "blue"}
       fmt.Printf("Value: %#v\n", c)
       fmt.Printf("Value with fields: %+v\n", c)
       fmt.Printf("Type: %T\n", c)
   }

If you run the above program, the output is going to be like this:

::

   Value: {76.45 blue}
   Value with fields: {radius:76.45 color:blue}
   Type: main.Circle

As you can see from the output, in the first line ``%v`` doesn’t show
the fields. But in the second line, ``%+v`` shows the struct fields.

The ``%#v`` shows the representation of the value. Here is a modified
version of above program to print few values of primitive type.

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       fmt.Printf("Value: %#v, Type: %T\n", "Baiju", "Baiju")
       fmt.Printf("Value: %#v, Type: %T\n", 7, 7)
       fmt.Printf("Value: %#v, Type: %T\n", uint(7), uint(7))
       fmt.Printf("Value: %#v, Type: %T\n", int8(7), int8(7))
       fmt.Printf("Value: %#v, Type: %T\n", true, true)
       fmt.Printf("Value: %#v, Type: %T\n", 7.0, 7.0)
       fmt.Printf("Value: %#v, Type: %T\n", (1 + 6i), (1 + 6i))
   }

::

   Value: "Baiju", Type: string
   Value: 7, Type: int
   Value: 0x7, Type: uint
   Value: 7, Type: int8
   Value: true, Type: bool
   Value: 7, Type: float64
   Value: (1+6i), Type: complex128

As you can see in the representation, strings are written within quotes.
You can also see representation of few other primitive types.

If you want a literal ``%`` sign, use two ``%`` signs next to each
other. Here is a code snippet:

::

   fmt.Println("Tom scored 92%% marks")

The default string representation of custom types can be changed by
implementing ``fmt.Stringer`` interafce. The interface definition is
like this:

::

   type Stringer interface {
           String() string
   }

As per the ``Stringer`` interface, you need to create a ``String``
function which return a string. Now the value printed will be whatever
returned by that function. Here is an example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
       "strconv"
   )

   // Temperature repesent air temperature
   type Temperature struct {
       Value float64
       Unit  string
   }

   func (t Temperature) String() string {
       f := strconv.FormatFloat(t.Value, 'f', 2, 64)
       return f + " degree " + t.Unit
   }

   func main() {
       temp := Temperature{30.456, "Celsius"}
       fmt.Println(temp)
       fmt.Printf("%v\n", temp)
       fmt.Printf("%+v\n", temp)
       fmt.Printf("%#v\n", temp)
   }

The output of the above program will be like this:

::

   30.46 degree Celsius
   30.46 degree Celsius
   30.46 degree Celsius
   main.Temperature{Value:30.456, Unit:"Celsius"}

Exercises
---------

**Exercise 1:** Write a program to read length and width of a rectangle
through command line arguments and print the area. Use ``-length``
switch to get length and ``-width`` switch to get width. Represent the
rectangle using a struct.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import (
       "flag"
       "fmt"
       "log"
       "os"
   )

   // Rectangle represents a rectangle shape
   type Rectangle struct {
       Length float64
       Width  float64
   }

   // Area return the area of a rectangle
   func (r Rectangle) Area() float64 {
       return r.Length * r.Width
   }

   var length = flag.Float64("length", 0, "length of rectangle")
   var width = flag.Float64("width", 0, "width of rectangle")

   func main() {
       flag.Parse()
       if *length <= 0 {
           log.Println("Invalid length")
           os.Exit(1)
       }
       if *width <= 0 {
           log.Println("Invalid width")
           os.Exit(1)
       }
       r := Rectangle{Length: *length, Width: *width}
       a := r.Area()
       fmt.Println("Area: ", a)
   }

You can run the program like this:

::

   $ go run rectangle.go -length 2.5 -width 3.4
   Area:  8.5

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a program to format a complex number as used in
mathematics. Example: ``2 + 5i``

Use a struct like this to define the complex number:

::

   type Complex struct {
       Real float64
       Imaginary float64
   }

Summary
-------

This chapter discussed about various input/output related
functionalities in Go. The chapter explained using command line
arguments and interactive input. The chaptered using *flag* package. It
also explained about various string formatting techniques.

.. [1]
   https://en.wikipedia.org/wiki/File_descriptor

.. [2]
   https://en.wikipedia.org/wiki/Standard_streams
