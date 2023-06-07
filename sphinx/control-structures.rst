Control Structures
==================

   *Science is what we understand well enough to explain to a computer.
   Art is everything else we do.* — Donald E. Knuth

Control structure determines how the code block will get executed for
the given conditional expressions and parameters. Go provides a number
of control structures including *for*, *if*, *switch*, *defer*, and
*goto*. The quickstart chapter has already introduced control structures
like *if* and *for*. This chapter will elaborate more about these topics
and introduce some other related topics.

If
--

Basic If
~~~~~~~~

Programming involves lots of decision making based on the input
parameters. The If control structure allows to perform a particular
action on certain condition. A conditional expressions is what used for
making decisions in the code using the If control structure. So, the If
control structure will be always associated with a conditional
expression which evaluates to a boolean value. If the boolean value is
true, the statements within the block will be executed. Consider this
example::

::

   package main

   import "fmt"

   func main() {
       if 1 < 2 {
           fmt.Println("1 is less than 2")
        }
   }

The first line starts with the ``if`` keyword followed by a conditional
expression and the line ends with an opening curly bracket. If the
conditional expression is getting evaluated as true, the statements
given within the curly brace will get evaluated. In the above example,
the conditional expression ``1 < 2`` will be evaluated as true so the
print statement given below that will get executed. In fact, you can add
any number of statements within the braces.

If and Else
~~~~~~~~~~~

Sometimes you need to perform different set of action if the condition
is not true. Go provides a variation of the ``if`` syntax with the
``else`` block for that.

Consider this example with else block::

::

   package main

   import "fmt"

   func main() {
       if 3 > 4 {
           fmt.Println("3 is greater than 4")
       } else {
           fmt.Println("3 is not greater than 4")
       }
   }

In the above code, the code will be evaluated as false. So, the
statements within ``else`` block will get executed.

Example
~~~~~~~

Now we will go through a complete example, the task is to identify the
given number is even or odd. The input can be given as a command line
argument.

::

   package main

   import (
        "fmt"
        "os"
        "strconv"
   )

   func main() {
        i := os.Args[1]
        n, err := strconv.Atoi(i)
        if err != nil {
                fmt.Println("Not a number:", i)
                os.Exit(1)
        }
        if n%2 == 0 {
                fmt.Printf("%d is even\n", n)
        } else {
                fmt.Printf("%d is odd\n", n)
        }
   }

The ``os`` package has an attribute named ``Args``. The value of
``Args`` will be a slice of strings which contains all command line
arguments passed while running the program. As we have learned from the
Quickstart chapter, the values can be accessed using the index syntax.
The value at zero index will be the program name itself and the value at
1st index the first argument and the value at 2nd index the second
argument and so on. Since we are expecting only one argument, you can
access it using the 1st index (``os.Args[1]``).

The ``strconv`` package provides a function named ``Atoi`` to convert
strings to integers. This function return two values, the first one is
the integer value and the second one is the error value. If there is no
error during convertion, the error value will be ``nil``. If it’s a
non-nil value, that indicates there is an error occured during
conversion.

The ``nil`` is an identifier in Go which represents the “zero value” for
certain built-in types. The ``nil`` is used as the zero for these types:
interfaces, functions, pointers, maps, slices, and channels.

In the above example, the second expected value is an object conforming
to built-in ``error`` interface. We will discuss more about errors and
interfaces in later chapters. The zero value for interface, that is
``nil`` is considered as there is no error.

The ``Exit`` function within ``os`` package helps to exit the program
prematurely. The argument passed will be exit status code. Normally exit
code ``0`` is treated as success and non-zero value as error.

The conditional expression use the modulus operator to get remainder and
checking it is zero. If the remainder against 2 is zero, the value will
be even otherwise the value will odd.

Else If
~~~~~~~

There is a third alternative syntax available for the If control
structure, that is ``else if`` block. The Else If block get executed if
the conditional expression gives true value and previous conditions are
false. It is possible to add any number of Else If blocks based on the
requirements.

Look at this example where we have three choices based on the age group.

::

   package main

   import "fmt"

   func main() {
           age := 10
           if age < 10 {
           fmt.Println("Junior", age)
       } else if age < 20 {
           fmt.Println("Senior", age)
       } else {
           fmt.Println("Other", age)
       }
   }

In the above example, the value printed will be either ``Junior``,
``Senior`` or ``Other``. You can change age value and run the program
again and again to see the outputs. The Else If can be repeated here to
create more choices.

Inline Statement
~~~~~~~~~~~~~~~~

In the previous section, the variable *age* was only within the If, Else
If and Else blocks. And that variable was not used used afterwards in
the function. Go provides a syntax to define a variable along with the
If where the scope of that variable will be within the blocks. In fact,
the syntax can be used for any valid Go statement. However, it is mostly
used for declaring variables.

Here is an example where a variable named ``money`` is declared along
with the If control structure.

::

   package main

   import "fmt"

   func main() {
       if money := 20000; money > 15000 {
           fmt.Println("I am going to buy a car.", money)
       } else {
           fmt.Println("I am going to buy a bike.", money)
       }
       // can't use the variable `money` here
   }

As mentioned above, the variable declared by the inline statement is
available only within the scope of If, Else If and Else blocks. So, the
variable ``money`` cannot be used outside the blocks.

It is possible to make any valid Go statement as part of the If control
structure. For example, it is possible to call a function like this:

::

   if money := someFunction(); money > 15000 {

For
---

Basic For
~~~~~~~~~

As we have seen briefly in the Quickstart, the For control structure
helps to create loops to repeat certain actions. The For control
structure has few syntax variants.

Consider a program to print few names.

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
       for i := 0; i < len(names); i++  {
           fmt.Println(names[i])
       }
   }

You can save the above program in a file named ``names.go`` and run it
like this:

::

   $ go run name.go
   Tom
   Polly
   Huck
   Becky

In the above example, ``names`` variable hold a slice of strings. The
value of ``i`` is initialized to zero and incremented one by one. The
``i++`` statement increment the value of ``i``. The second part of
``for`` loop check if value of ``i`` is less than length of the slice.
The built-in ``len`` gives the length of slice.

Other programming languages offer many ways for iterations. Some of the
examples are *while* and *do...while*. But in Go using syntactic
variation of *for* loop meets all requirements. Functional languages
prefer to use recursion instead of iteration.

Break Loop Prematurely
~~~~~~~~~~~~~~~~~~~~~~

Sometimes the iteration should be stopped prematurely on certain
condition. This can be achieved using the If control structure and break
statement. We have already studied If control structure from the
previous major section. The ``break`` keyword allows to create a break
statement. The break statement end the loop immediately. Though any
other code followed by For loop will be executed.

Let’s alter the previous program to stop printing after the name
``Polly`` found.

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
       for i := 0; i < len(names); i++  {
           fmt.Println(names[i])
           if names[i] == "Polly" {
               break
           }
       }
   }

In the above example, we added an If control structure to check for the
value of name during each iteration. If the value matches ``Polly``,
break statement will be executed. The break statement makes the For loop
to end immediately.

As you can see in the above code, the break statement can stand alone
without any other input. There is alternate syntax with label similar to
how *goto* works, which we are going to see below. This is useful when
you have multiple loops and want to break a particular one, may be the
outer loop.

To understand this better, let’s consider an example. The problem is to
to change print the name given the slice until a word with ``u`` found.

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
   Outer:
       for i := 0; i < len(names); i++ {
           for j := 0; j < len(names[i]); j++ {
               if names[i][j] == 'u' {
                   break Outer
               }
           }
           fmt.Println(names[i])
       }
   }

In the above example, we are declaring a label statement just before the
first For loop. There is an inner loop to iterate through the name
string and check for the presence of character ``u``. If the character
``u`` is found, then it will break the outer loop. If the label
``Outer`` is not used in the break statement, then the inner loop will
be stopped.

Partially Execute Loop Statements
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

Sometimes statements within For loop should be executed on certain
iterations. Go has a ``continue`` statement to proceed loop without
executing further statements.

Let’s modify the previous problem to print all names except ``Polly``.

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
       for i := 0; i < len(names); i++ {
           if names[i] == "Polly" {
               continue
           }
           fmt.Println(names[i])
       }
   }

In the above code, the ``continue`` statement makes it proceed with next
iteration in the loop without printing ``Polly``.

Similar to ``break`` statement with label, continue also can be used
with a label. This is useful if there are multiple loops and want to
continue a particular loop, say the outer one.

Let’s consider an example where you need to print names which doesn’t
have character ``u`` in it.

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
   Outer:
       for i := 0; i < len(names); i++ {
           for j := 0; j < len(names[i]); j++ {
               if names[i][j] == 'u' {
                   continue Outer
               }
           }
           fmt.Println(names[i])
       }
   }

In the above code, just before the first loop a label is declared. Later
inside the inner loop to iterate through the name string and check for
the presence of character ``u``. If the character ``u`` is found, then
it will continue the outer loop. If the label ``Outer`` is not used in
the continue statement, then the inner loop will be proceed to execute.

For with Outside Initialization
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

The statement for value initialization and the last pat to increment
value can be removed from the For control structure. The value
initialization can be moved outside For and value increment can be moved
inside loop.

The previous example can be changed like this:

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
       i := 0
       for i < len(names) {
           i++
           fmt.Println(names[i])
       }
   }

In the above example, the scope of variable ``i`` is outside For loop
code block. Whereas in the previous section, when the variable declared
along with For loop, the scope of that variable was within the loop code
block.

Infinite Loop
~~~~~~~~~~~~~

For loop has yet another syntax variant to support infinite loop. You
can create a loop that never ends until explicitly stopped using break
or exiting the whole program. To create an infinite loop, you can use
the ``for`` keyword followed by the curly bracket.

If any variable initialization is required, that should be declared
outside the loop. Conditions can be added inside the loop.

The previous example can be changed like this:

::

   package main

   import "fmt"

   func main() {
       names := []string{"Tom", "Polly", "Huck", "Becky"}
       i := 0
       for {
           if i >= len(names) {
               break
           }
           fmt.Println(names[i])
           i++
       }
   }

Range Loops
~~~~~~~~~~~

The range clause form of the for loop iterates over a slice or map. When
looping over a slice using range, two values are returned for each
iteration. The first is the index, and the second is a copy of the
element at that index.

The previous example ``for`` loop can be simplified using the ``range``
clause like this:

::

   package main

   import "fmt"

   func main() {
       characters := []string{"Tom", "Polly", "Huck", "Becky"}
       for _, j := range characters {
           fmt.Println(j)
       }
   }

The underscore is called blank indentifier, the value assigned to that
variable will be ignored. In the above example, the index values will be
assigned to the underscore.

The range loop can be used with map. Here is an example:

::

   package main

   import "fmt"

   func main() {
       var characters = map[string]int{
                   "Tom": 8,
                   "Polly": 51,
                   "Huck": 9,
                   "Becky": 8,
       }
       for name, age := range characters {
           fmt.Println(name, age)
       }
   }

Switch Cases
------------

Basic Switch
~~~~~~~~~~~~

In addition to the ``if`` condition, Go provides ``switch case`` control
structure for branch instructions. The ``switch case`` is more
convenient if many cases need to be handled in the branch instructions.

The below program use a switch case to print number names based on the
value.

::

   package main

   import "fmt"

   func main() {
       v := 1
       switch v {
       case 0:
               fmt.Println("zero")
       case 1:
               fmt.Println("one")
       case 2:
               fmt.Println("two")
       default:
               fmt.Println("unknown")
       }
   }

In this case, the value of ``v`` is ``1``, so the case that is going to
execute is 2nd one. This will be the output.

::

   $ go run switchbasic.go
   one

If you change the value of ``v`` to ``0``, it’s going to print ``zero``
and for ``2`` it will print ``two``. If the value is any number other
than ``0``, ``1`` or ``2``, it’s going to print ``unknown``.

Fallthrough
~~~~~~~~~~~

The cases are evaluated top to bottom until a match is found. If a case
is matched, the statements within that case will be executed. And no
other case will be executed unless a ``fallthrough`` statement is used.
The ``fallthrough`` must be the last statement within the case.

Here is a modified version with ``fallthrough``

::

   package main

   import "fmt"

   func main() {
       v := 1
       switch v {
       case 0:
               fmt.Println("zero")
       case 1:
               fmt.Println("one")
               fallthrough
       case 2:
               fmt.Println("two")
       default:
               fmt.Println("unknown")
       }
   }

If you run this program, it will print ``one`` followed by ``two``.

::

   $ go run switchbasic.go
   one
   two

Break
~~~~~

As you can see from the above examples, the switch statements break
implicitly at the end of each cases. The ``fallthrough`` statement can
be used to passdown control to the next case. However, sometimes
execution should be stopped early without executing all statements. This
can can be achieved using ``break`` statements.

Here is an example:

::

   package main

   import (
       "fmt"
       "time"
   )

   func main() {
       v := "Becky"
       t := time.Now()
       switch v {
       case "Huck":
           if t.Hour() < 12 {
               fmt.Println("Good morning,", v)
               break
           }
           fmt.Println("Hello,", v)
       case "Becky":
           if t.Hour() < 12 {
               fmt.Println("Good morning,", v)
               break
           }
           fmt.Println("Hello,", v)
       default:
           fmt.Println("Hello")
       }
   }

In the above example, morning time greeting is different.

Multiple Cases
~~~~~~~~~~~~~~

Multple cases can be presented in comma-separated lists.

Here is an example.

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       o := "=="
       switch o {
       case "+", "-", "*", "/", "%", "&", "|", "^", "&^", "<<", ">>":
           fmt.Println("Arithmetic operator")
       case "==", "!=", "<", "<=", ">", ">=":
           fmt.Println("Comparison operators")
       case "&&", "||", "!":
           fmt.Println("Logical operators")
       default:
           fmt.Println("Unknown operator")
       }
   }

In this example, if any of the value is matched in the given list, that
case will be executed.

Without Expression
~~~~~~~~~~~~~~~~~~

If the switch has no expression it switches on true. This is useful to
write an if-else-if-else chain.

Let’s take the example program used earlier when Else If was introduced:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       age := 10
       switch {
       case age < 10:
           fmt.Println("Junior", age)
       case age < 20:
           fmt.Println("Senior", age)
       default:
           fmt.Println("Other", age)
       }
   }

Defer Statements
----------------

Sometimes it will require to force certain things to do before a
function returns. For example, closing an opened file descriptor. Go
provides the *defer* statements to do these kind of cleanup actions.

A defer statement add a function call into a stack. The stack of
function call executes at the end of the surrounding function in a
last-in-first-out (LIFO) order. Defer is commonly used to perform
various clean-up actions.

Here is a simple example:

::

   package main

   import (
       "fmt"
       "time"
   )

   func main() {
       defer fmt.Println("world")
       fmt.Println("hello")
   }

The above program is going to print ``hello`` followed by ``world``.

If there are multiple *defer* statements, it will execute in
last-in-first-out (LIFO) order.

Here is a simple example to demonstrate it:

::

   package main

   import "fmt"

   func main() {
       for i := 0; i < 5; i++ {
           defer fmt.Println(i)
       }
   }

The above program will print this output:

::

   4
   3
   2
   1
   0

The arguments passed the the deferred call are evaluated immediately.
But the deferred call itself is not executed until the function returns.
Here is a simple example to demonstrate it:

::

   package main

   import (
       "fmt"
       "time"
   )

   func main() {
       defer func(t time.Time) {
           fmt.Println(t, time.Now())
       }(time.Now())
   }

When you run the above program, you can see a small difference in time.
The *defer* can also be used to recover from *panic*, which will be
discussed in the next section.

Deffered Panic Recover
----------------------

We have discussed the commonly used control structures including if,
for, and switch. This section is going to discuss a less commonly used
set of control structures: *defer*, *panic*, and *recover*. We have
discussed the use of the defer statement in the previous section. In
this section, you are going to learn how to use the *defer* along with
*panic* and *recover*.

Few important points about defer, panic, and recover:

-  A panic causes the program stack to begin unwinding and recover can
   stop it

-  Deferred functions are still executed as the stack unwinds

-  If recover is called inside such a deferred function, the stack stops
   unwinding

-  The recover returns the value (as an *interface{}*) that was passed
   to panic

-  A panic cannot be recovered by a different goroutine

Here is an example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
       defer func() {
           if r := recover(); r != nil {
               fmt.Println("Recoverd", r)
           }
       }()
       panic("panic")
   }

Goto
----

The *goto* statement can be used to jump control to another statement.
The location to where the control should be passed is specified using
label statements. The *goto* statement and the corresponding label
should be within the same function. The *goto* cannot jump to a label
inside another code block.

::

   package main

   import "fmt"

   func main() {
       num := 10
       goto Marker
       num = 20
   Marker:
       fmt.Println("Value of num:", num)
   }

You can save the above program in a file named ``goto1.go`` and run it
like this:

::

   $ go run goto1.go
   Value of num: 10

In the above code ``Marker:`` is a label statement. A label statement is
a valid identifier followed by a colon. A label statement will be target
for *goto*, *break* or *continue* statement. We will look at *break* and
*continue* statement when we study the For control structure.

The *goto* statement is writen using the ``goto`` keyword followed by a
valid label name. In the above code, immediately after the *goto*
statement, there is a statement to assign a different value to ``num``.
But that statement is never getting executed as the *goto* makes the
program to jump to the label.

Exercises
---------

**Exercise 1:** Print whether the number given as the command line
argument is even or odd.

**Solution:**

You can store the program with a file named ``evenodd.go``. Later you
can compile this program and then you will get a binary executable with
name as ``evenodd``. You can execute this program like this:

::

   ./evenodd 3
   3 is odd
   ./evenodd 4
   4 is even

In the above program, the 3 and 4 are the command line arguments. The
command line arguments can be accessed from Go using the slice available
under ``os`` package. The arguments will be available with exported name
as ``Args`` and individual items can be accessed using the index. The
0th index contains the program itself, so it can be ignored. To access
the 1st command argument use ``os.Args[1]``. The values will be of type
string which can be converted later.

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
       "os"
       "strconv"
   )

   func main() {
       i := os.Args[1]
       n, err := strconv.Atoi(i)
       if err != nil {
           fmt.Println("Not a number:", i)
           os.Exit(1)
       }
       if n%2 == 0 {
           fmt.Printf("%d is even\n", n)
       } else {
           fmt.Printf("%d is odd\n", n)
       }
   }

**Exercise 2:** Write a program to print numbers below 20 which are
multiples of 3 or 5.

**Solution:**

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func main() {
   	for i := 1; i < 20; i++ {
   		if i%3 == 0 || i%5 == 0 {
   			fmt.Println(i)
   		}
   	}
   }

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a program to print greetings based on time.
Possible greetings are Good morning, Good afternoon and Good evening.

**Problem 2:** Write a program to check if a given number is a multiple
of 2, 3, or 5.

Summary
-------

This chapter introduced control structures available in Go except those
related to concurrency. The *if* control structure was covered first,
then *for* loop explained. The *switch* cases was discussed later. Then
*defer* statement and finally *goto* control structure was explained in
detail. This chapter also briefly explained about accessing command line
arguments from the program.
