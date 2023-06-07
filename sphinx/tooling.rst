Tooling
=======

   *We become what we behold. We shape our tools, and thereafter our
   tools shape us.* — Marshall McLuhan

Good support for lots of useful tools is another strength of Go. Apart
from the built-in tools, there any many other community-built tools.
This chapter covers the built-in Go tools and few other external tools.

The built-in Go tools can access through the ``go`` command. When you
install the Go compiler (``gc``); the ``go`` tool is available in the
path. The ``go`` tool has many commands. You can use the ``go`` tool to
compile Go programs, run test cases, and format source files among other
things.

Getting Help
------------

The go tool is self-documented. You can get help about any commands
easily. To see the list of all commands, you can run the ``"help"``
command. For example, to see help for ``build`` command, you can run
like this:

::

   go help build

The help command also provides help for specific topics like
"buildmode", "cache", "filetype", and "environment" among other topics.
To see help for a specific topic, you can run the command like this:

::

   go help environment

Basic Information
-----------------

Version
~~~~~~~

When reporting bugs, it is essential to specify the Go version number
and environment details. The Go tool gives access to this information
through the following commands.

To get version information, run this command:

::

   go version

The output should look something like this:

::

   go version go1.20.4 linux/amd64

As you can see, it shows the version number followed by operating system
and CPU architecture.

Environment
~~~~~~~~~~~

To get environment variables, you can run this command:

::

   go env

The output should display all the environment variables used by the Go
tool when running different commands.

A typical output will look like this:

::

   GO111MODULE=""
   GOARCH="amd64"
   GOBIN=""
   GOCACHE="/home/baiju/.cache/go-build"
   GOENV="/home/baiju/.config/go/env"
   GOEXE=""
   GOEXPERIMENT=""
   GOFLAGS=""
   GOHOSTARCH="amd64"
   GOHOSTOS="linux"
   GOINSECURE=""
   GOMODCACHE="/home/baiju/go/pkg/mod"
   GONOPROXY=""
   GONOSUMDB=""
   GOOS="linux"
   GOPATH="/home/baiju/go"
   GOPRIVATE=""
   GOPROXY="https://proxy.golang.org,direct"
   GOROOT="/usr/local/go"
   GOSUMDB="sum.golang.org"
   GOTMPDIR=""
   GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
   GOVCS=""
   GOVERSION="go1.20.4"
   GCCGO="gccgo"
   GOAMD64="v1"
   AR="ar"
   CC="gcc"
   CXX="g++"
   CGO_ENABLED="1"
   GOMOD="/dev/null"
   GOWORK=""
   CGO_CFLAGS="-O2 -g"
   CGO_CPPFLAGS=""
   CGO_CXXFLAGS="-O2 -g"
   CGO_FFLAGS="-O2 -g"
   CGO_LDFLAGS="-O2 -g"
   PKG_CONFIG="pkg-config"
   GOGCCFLAGS="-fPIC -m64 -pthread -Wl,--no-gc-sections -fmessage-length=0
     -fdebug-prefix-map=/tmp/go-build1378738152=/tmp/go-build
     -gno-record-gcc-switches"

List
~~~~

The *list* command provides meta information about packages. Running
without any arguments shows the current packages import path. The *-f*
helps to extract more information, and it can specify a format. The
*text/template* package syntax can be used to specify the format.

The struct used to format has many attributes, here is a subset:

-  *Dir* – directory containing package sources

-  *ImportPath* – import path of package in dir

-  *ImportComment* – path in import comment on package statement

-  *Name* – package name

-  *Doc* – package documentation string

-  *Target* – install path

-  *GoFiles* – list of ``.go`` source files

Here is an example usage:

::

   $ go list -f '{{.GoFiles}}' text/template
   [doc.go exec.go funcs.go helper.go option.go template.go]

Building and Running
--------------------

To compile a program, you can use the ``build`` command. To compile a
package, first change to the directory where the program is located and
run the ``build`` command:

::

   go build

You can also compile Go programs without changing the directory. To do
that, you are required to specify the package location in the command
line. For example, to compile ``github.com/baijum/introduction`` package
run the command given below:

::

   go build github.com/baijum/introduction

If you want to set the executable binary file name, use the ``-o``
option:

::

   go build -o myprog

If you want to build and run at once, you can use the ``"run"`` command:

::

   go run program.go

You can specify more that one Go source file in the command line
arguments:

::

   go run file1.go file2.go file3.go

Of course, when you specify more than one file names, only one "main"
function should be defined among all of the files.

Conditional Compilation
~~~~~~~~~~~~~~~~~~~~~~~

Sometimes you need to write code specific to a particular operating
system. In some other case the code for a particular CPU architecture.
It could be code optimized for that particular combination. The Go build
tool supports conditional compilation using build constraints. The Build
constraint is also known as build tag. There is another approach for
conditional compilation using a naming convention for files names. This
section is going to discuss both these approaches.

The build tag should be given as comments at the top of the source code.
The build tag comment should start like this:

::

   // +build

The comment should be before package documentation and there should be a
line in between.

The space is *OR* and comma is *AND*. The exclamation character is
stands for negation.

Here is an example:

::

   // +build linux,386

In the above example, the file will compile on 32-bit x86 Linux system.

::

   // +build linux darwin

The above one compiles on Linux or Darwin (Mac OS).

::

   // +build !linux

The above runs on anything that is not Linux.

The other approach uses file naming convention for conditional
compilation. The files are ignore if it doesn’t match the target OS and
CPU architecture, if any.

This compiles only on Linux:

::

   stat_linux.go

This one on 64 bit ARM linux:

::

   os_linux_arm64.go

Running Test
------------

The Go tool has a built-in test runner. To run tests for the current
package, run this command:

::

   go test

To demonstrate the remaining commands, consider packages organized like
this:

::

   .
   |-- hello.go
   |-- hello_test.go
   |-- sub1
   |   |-- sub1.go
   |   `-- sub1_test.go
   `-- sub2
       |-- sub2.go
       `-- sub2_test.go

If you run ``go test`` from the top-level directory, it’s going to run
tests in that directory, and not any sub directories. You can specify
directories as command line arguments to ``go test`` command to run
tests under multiple packages simultaneously. In the above listed case,
you can run all tests like this:

::

   go test . ./sub1 ./sub2

Instead of listing each packages separates, you can use the ellipsis
syntax:

::

   go test ./...

The above command run tests under current directory and its child
directories.

By default ``go test`` shows very few details about the tests.

::

   $ go test ./...
   ok      _/home/baiju/code/mypkg   0.001s
   ok      _/home/baiju/code/mypkg/sub1      0.001s
   --- FAIL: TestSub (0.00s)
   FAIL
   FAIL    _/home/baiju/code/mypkg/sub2      0.003s

In the above results, it shows the name of failed test. But details
about other passing tests are not available. If you want to see verbose
results, use the ``-v`` option.

::

   $ go test ./... -v
   === RUN   TestHello
   --- PASS: TestHello (0.00s)
   PASS
   ok      _/home/baiju/code/mypkg   0.001s
   === RUN   TestSub
   --- PASS: TestSub (0.00s)
   PASS
   ok      _/home/baiju/code/mypkg/sub1      0.001s
   === RUN   TestSub
   --- FAIL: TestSub (0.00s)
   FAIL
   FAIL    _/home/baiju/code/mypkg/sub2      0.002s

If you need to filter tests based on the name, you can use the ``-run``
option.

::

   $ go test ./... -v -run Sub
   testing: warning: no tests to run
   PASS
   ok      _/home/baiju/code/mypkg   0.001s [no tests to run]
   === RUN   TestSub
   --- PASS: TestSub (0.00s)
   PASS
   ok      _/home/baiju/code/mypkg/sub1      0.001s
   === RUN   TestSub
   --- FAIL: TestSub (0.00s)
   FAIL
   FAIL    _/home/baiju/code/mypkg/sub2      0.002s

As you can see above, the ``TestHello`` test was skipped as it doesn’t
match "Sub" pattern.

The chapter on testing has more details about writing test cases.

| *golangci-lint* is a handy program to run various lint tools and
  normalize their output. This program is useful to run through
  continuous integration. You can download the program from here:
  https://github.com/golangci/golangci-lint
| The supported lint tools include Vet, Golint, Varcheck, Errcheck,
  Deadcode, Gocyclo among others. *golangci-lint* allows to
  enable/disable the lint tools through a configuration file.

Formatting Code
---------------

Go has a recommended source code formatting. To format any Go source
file to conform to that format, it’s just a matter of running one
command. Normally you can integrate this command with your text editor
or IDE. But if you really want to invoke this program from command line,
this is how you do it:

::

   go fmt myprogram.go

In the above command, the source file name is explicitly specified. You
can also give package name:

::

   go fmt github.com/baijum/introduction

The command will format source files and write it back to the same file.
Also it will list the files that is formatted.

Generating Code
---------------

If you have use case to generate Go code from a grammar, you may
consider the *go generate*. In fact, you can add any command to be run
before compiling the code. You can add a special comment in your Go code
with this syntax:

::

   //go:generate command arguments

For example, if you want to use *peg*
(https://github.com/pointlander/peg), a Parsing Expression Grammar
implementation, you can add the command like this:

::

   //go:generate peg -output=parser.peg.go grammar.peg

When you build the program, the parser will be generated and will be
part of the code that compiles.

Embedding Code
--------------

Go programs are normally distributed as a single binary. What if your
program need some files to run. Go has a feature to embed files in the
binary. You can embed any type of files, including text files and binary
files. Some of the commonly embedded files are SQL, HTML, CSS,
JavaScript, and images. You can embed individual files or diectories
including nested sub-directories.

You need to import the *embed* package and use the ``//go:embed``
compiler directive to embed. Here is an example to embed an SQL file:

::

   import _ "embed"

   //go:embed database-schema.sql
   var dbSchema string

As you can see, the "embed" package is imported with a blank identifier
as it is not directly used in the code. This is required to initialize
the package to embed files. The variable must be at package level and
not at function or method level.

The variable could be slice of bytes. This is useful for binary files.
Here is an example:

::

   import _ "embed"

   //go:embed logo.jpg
   var logo []byte

If you need an entire directory, you can use the ``embed.FS`` as the
type:

::

   import "embed"

   //go:embed static
   var content embed.FS

Displaying Documentation
------------------------

Go has good support for writing documentation along with source code.
You can write documentation for packages, functions and custom defined
types. The Go tool can be used to display those documentation.

To see the documentation for the current packages, run this command:

::

   go doc

To see documentation for a specific package:

::

   go doc strings

The above command shows documentation for the "strings" package.

::

   go doc strings

If you want to see documentation for a particular function within that
package:

::

   go doc strings.ToLower

or a type:

::

   go doc strings.Reader

Or a method:

::

   go doc strings.Reader.Size

Find Suspicious Code Using Vet
------------------------------

There is a handy tool named *vet* to find suspicious code in your
program. Your program might compile and run. But some of the results may
not be desired output.

Consider this program:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   func main() {
       v := 1
       fmt.Printf("%#v %s\n", v)
   }

If you compile and run it. It’s going to be give some output. But if you
observe the code, there is an unnecessary ``%s`` format string.

If you run ``vet`` command, you can see the issue:

::

   $ go vet susp.go
   # command-line-arguments
   ./susp.go:9: Printf format %s reads arg #2,
   but call has only 1 arg

Note: The *vet* command is automatically run along with the *test*
command.

Exercises
---------

**Exercise 1:** Create a program with function to return "Hello, world!"
and write test and run it.

*hello.go*:

::

   package hello

   // SayHello returns a "Hello word!" message
   func SayHello() string {
        return "Hello, world!"
   }

*hello_test.go*:

::

   package hello

   import "testing"

   func TestSayHello(t *testing.T) {
        out := SayHello()
        if out != "Hello, world!" {
           t.Error("Incorrect message", out)
        }
   }

To run the test:

::

   go test . -v

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a program with exported type and methods with
documentation strings. Then print the documentation using ``go doc``
command.

Summary
-------

This chapter introduced the Go tool. All the Go commands were explained
in detail. Practical example usage was also given for each command. The
chapter coverd how to build and run programs, running tests, formatting
code, and displaying documentation. It also touched upon few other handy
tools.
