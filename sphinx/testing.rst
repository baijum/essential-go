Testing
=======

   *Test as You Fly, Fly as You Test.* — NASA

Writing automated tests helps you to improve the quality and reliability
of software. This chapter is about writing automated tests in Go. The
standard library contains a package named *testing* to write tests.
Also, the built-in Go tool has a test runner to run tests.

Consider this test for a ``Hello`` function which takes a string as
input and returns another string. The expected output string should
start with "Hello," and ends with the input value followed by an
exclamation.

.. code-block:: go
   :linenos:

   package hello

   import "testing"

   func TestHello(t *testing.T) {
       out := Hello("Tom")
       if out != "Hello, Tom!" {
           t.Fail()
       }
   }

In the first line, the package name is ``hello`` which is the same as
the package where the function is going to be defined. Since both test
and actual code is in the same package, the test can access any name
within that package irrespective whether it is an exported name or not.
At the same, when the actual problem is compiled using the ``go build``
command, the test files are ignored. The build ignores all files with
the name ending ``_test.go``. Sometimes these kinds of tests are called
white-box test as it is accessing both exported and unexported names
within the package. If you only want to access exported names within the
test, you can declare the name of the test package with ``_test``
suffix. In this case, that would be ``hello_test``, which should work in
this case as we are only testing the exported function directly.
However, to access those exported names – in this case, a function – the
package should import explicitly.

The line no. 5 starts the test function declaration. As per the test
framework, the function name should start with ``Test`` prefix. The
prefix is what helps the test runner to identify the tests that should
be running.

The input parameter ``t`` is a pointer of type ``testing.T``. The
``testing.T`` type provides functions to make the test pass or fail. It
also gives functions to log messages.

In the line no. 6 the ``Hello`` function is called with "Tom" as the
input string. The return value is assigning to a variable named ``out``.

In line no. 7 the actual output value is checked against the expected
output value. If the values are not matching, the ``Fail`` function is
getting called. The particular function state is going to be marked as a
failure when the ``Fail`` function is getting called.

The test is going to pass or fail based on the implementation of the
``Hello`` function. Here is an implementation of the function to satisfy
the test:

.. code-block:: go
   :linenos:

   package hello

   import "fmt"

   // Hello says "Hello" with name
   func Hello(name string) string {
       return fmt.Sprintf("Hello, %s!", name)
   }

As you can see in the above function definition, it takes a string as an
input argument. A new string is getting constructed as per the
requirement, and it returns that value.

Now you can run the test like this:

::

   $ go test
   PASS
   ok      _/home/baiju/hello     0.001s

If you want to see the verbose output, use the ``-v`` option:

::

   $ go test -v
   === RUN   TestHello
   --- PASS: TestHello (0.00s)
   PASS
   ok      _/home/baiju/hello     0.001s

Failing a Test
--------------

To fail a test, you need to explicitly call ``Fail`` function provided
by the value of ``testing.T`` type. As we have seen before, every test
function has access to a ``testing.T`` object. Usually, the name of that
value is going to write as ``t``. To fail a test, you can call the
``Fail`` function like this:

::

   t.Fail()

The test is going be a failure when the ``Fail`` function is getting
called. However, the remaining code in the same test function continue
to execute. If you want to stop executing the further lines immediately,
you can call ``FailNow`` function.

::

   t.FailNow()

Alternatively, there are other convenient functions which give similar
behavior along with logging message. The next section discusses logging
messages.

Logging Message
---------------

The ``testing.T`` has two functions for logging, one with default
formatting and the other with the user-specified format. Both functions
accept an arbitrary number of arguments.

The ``Log`` function formats its arguments using the default formats
available for any type. The behavior is similar to ``fmt.Println``
function. So, you can change the formatted value by implementing the
``fmt.Stringer`` interface:

::

   type Stringer interface {
           String() string
   }

You need to create a method named ``String`` which returns a string for
your custom types.

Here is an example calling ``Log`` with two arguments:

::

   t.Log("Some message", someValue)

In the above function call, there are only two arguments given, but you
can pass any number of arguments.

The log message going to print if the test is failing. The verbose mode,
the ``-v`` command-line option, also log the message irrespective of
whether a test fails or not.

The ``Logf`` function takes a string format followed by arguments
expected by the given string format. Here is an example:

::

   t.Logf("%d no. of lines: %s", 34, "More number of lines")

The ``Logf`` formats the values based on the given format. The ``Logf``
is similar to ``fmt.Printf`` function.

Failing with Log Message
------------------------

Usually, logging and marking a test as failure happens simultaneously.
The ``testing.T`` has two functions for logging with failing, one with
default formatting and the other with the user-specified format. Both
functions accept an arbitrary number of arguments.

The ``Error`` function is equivalent to calling ``Log`` followed by
``Fail``. The function signature is similar to ``Log`` function.

Here is an example calling ``Error`` with two arguments:

::

   t.Error("Some message", someValue)

Similar to ``Error`` function, the ``Errorf`` function is equivalent to
calling ``Logf`` followed by ``Fail``. The function signature is similar
to ``Logf`` function.

The ``Errorf`` function takes a string format followed by arguments
expected by the given string format. Here is an example:

::

   t.Errorf("%d no. of lines: %s", 34, "More number of lines")

The ``Errorf`` formats the values based on the given format.

Skipping Test
-------------

When writing tests, there are situations where particular tests need not
run. Some tests might have written for a specific environment. The
criteria for running tests could be CPU architecture, operating system
or any other parameter. The *testing* package has functions to mark test
for skipping.

The ``SkipNow`` function call marks the test as having been skipped. It
stops the current test execution. If the test has marked as failed
before skipping, that particular test is yet considered to have failed.
The ``SkipNow`` function doesn’t accept any argument. Here is a simple
example:

.. code-block:: go
   :linenos:

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

If you run the above code on a Linux system, you can see the test has
skipped. The output is going to be something like this:

::

   $ go test . -v
   === RUN   TestHello
   --- SKIP: TestHello (0.00s)
   PASS
   ok      _/home/baiju/skip      0.001s

As you can see from the output, the test has skipped execution.

There are two more convenient functions similar to ``Error`` and
``Errorf``. Those functions are ``Skip`` and ``Skipf``. These functions
help you to log a message before skipping. The message could be the
reason for skipping the test.

Here is an example:

::

   t.Skip("Some reason message", someValue)

The ``Skipf`` function takes a string format followed by arguments
expected by the given string format. Here is an example:

::

   t.Skipf("%d no. of lines: %s", 34, "More number of lines")

The ``Skipf`` formats the values based on the given format.

Parallel Running
----------------

You can mark a test to run in parallel. To do so, you can call the
``t.Parallel`` function. The test is going to run in parallel to other
tests marked parallel.

Sub Tests
---------

The Go testing package allows you to group related tests together in a
hierarchical form. You can define multiple related tests under a
single-parent test using the ‘Run‘ method.

To create a subtest, you use the ``t.Run()`` method. The ``t.Run()``
method takes two arguments: the name of the subtest and the body of the
subtest. The body of the subtest is a regular Go function.

For example, the following code creates a subtest called ``foo``:

::

   func TestBar(t *testing.T) {
     t.Run("foo", func(t *testing.T) {
       // This is the body of the subtest.
     })
   }

Subtests are reported separately from each other. This means that if a
subtest fails, the test runner will report the failure for that subtest
only. The parent test will still be considered to have passed.

Subtests can be used to test different aspects of a function or method.
For example, you could use subtests to test different input values,
different output values, or different error conditions.

Subtests can also be used to test different implementations of a
function or method. For example, you could use subtests to test a
function that is implemented in Go and a function that is implemented in
C.

Subtests are a powerful feature of the Go testing package. They can be
used to organize your tests, make them easier to read and maintain, and
test different aspects of your code.

Exercises
---------

**Exercise 1:** Create a package with a function to return the sum of
two integers and write a test for the same.

**Solution:**

*sum.go*:

.. code-block:: go
   :linenos:

   package sum

   // Add adds to integers
   func Add(first, second int) int {
       return first + second
   }

*sum_test.go*:

.. code-block:: go
   :linenos:

   package sum

   import "testing"

   func TestAdd(t *testing.T) {
       out := Add(2, 3)
       if out != 5 {
           t.Error("Sum of 2 and 3:", out)
       }
   }

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a test case program to fail the test and not
continue with the remaining tests.

Summary
-------

This chapter explained writing tests using the *testing* package. It
covered how to mark a test as a failure, logging, skipping, and parallel
running. Also, it briefly touched upon sub-tests.
