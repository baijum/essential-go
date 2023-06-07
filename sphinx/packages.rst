Packages
========

   *A little copying is better than a little dependency.* — Go Proverb

Go encourages and provides mechanisms for code reusability. Packages are
one of the building block for code reusability. We have seen other code
reusability mechanisms like functions in the earlier chapters.

This chapter will explain everything you need to know about packages. In
the previous programs, there was a clause at the beginning of every
source files like this:

::

   package main

That clause was declaring the name of the package as ``main``. The
package name was ``main`` because the program was executable. For
non-executable programs, you can give a different meaningful short name
as the package name.

Most of the previous programs imported other packages like this:

::

   import "fmt"

Sometimes the import was within parenthesis (factored imports) for
multiple packages:

::

   import (
          "fmt"
          "time"
   )

Importing packages gives access to functionalities provided by those
packages.

Packages helps you to create modular reusable programs. Go programs are
organized using packages. Every source file will be associated with a
package. Basically, package is a collection of source files.

A package generally consists of elements like constants, types,
variables and functions which are accessible across all the source
files. The source files should should start with a statement declaring
the name of the package. The package name should be a meaningful short
name. This is important because normally the exported names within the
package will be accessed with the package name as a prefix.

The Go standard library has many packages. Each package will be related
to certain ideas.

Creating a Package
------------------

Consider this file named ``circle1.go`` which belongs to package
``main``:

::

   package main

   const pi float64 = 3.14

   type Circle struct {
       Radius float64
   }

Another file named ``circle2.go`` in the same directory:

::

   package main

   func (c Circle) Area() float64 {
       return pi * c.Radius * c.Radius
   }

Yet another file named ``circle3.go``:

::

   package main

   import "fmt"

   var c Circle

   func main() {
       c = Circle{3.5}
       fmt.Println(c.Area())
   }

As you can see above all the above source files belongs to ``main``
package.

You can build and run it like this:

::

   $ go build -o circle
   $ ./circle
   38.465

Alternatively, you can use ``go run``:

::

   $ go run *.go
   38.465

The final Go program is a created by joining multiple files. The package
is a container for your code.

Package Initialization
----------------------

When you run a Go program, all the associated packages initialize in a
particular order. The package level variables initialize as per the
order of declaration. However, if there is any dependency on
initialization, that is resolved first.

Consider this example:

::

   package varinit

   var a = b + 2
   var b = 1

In the above program, the variable ``a`` is depending on the value of
variable ``b``. Therefore, the variable ``b`` initialize first and the
variable ``a`` is the second. If you import the above package from
different other packages in the same program, the variable
initialization happens only once.

If the expression for variable initialization is complicated, you can
use a function to return the value. There is another approach possible
for initializing variables with complex expressions, that is by using
the special *init* function. The Go runtime executes the *init* function
only once. Here is an example:

.. code-block:: go
   :linenos:

   package config

   import (
       "log"

       "github.com/kelseyhightower/envconfig"
   )

   type Configuration struct {
       Address        string `default:":8080"`
       TokenSecretKey string `default:"secret" split_words:"true"`
   }

   var Config Configuration

   func init() {
       err := envconfig.Process("app", &Config)
       if err != nil {
           log.Fatal(err.Error())
       }
   }

The *init* function cannot be called or even referred from the program –
it’s going to produce a compile error. Also, the *init* function should
not return any value.

Documenting Packages
--------------------

You can write documentation for the package in the form of comment
before the declarion of package name. The comment should begin with the
word *Package* followed by the name of the package. Here is an example:

::

   // Package hello gives various greeting messages
   package hello

If the package is defined in multiple files, you may create a file named
*doc.go*. The *doc.go* file is just a naming convention. You can write
multi-line comment using ``/* ... */`` syntax. You can write source code
with tab-indentation which will be highlighted and displayed using
monospace font in HTML format.

Here is an example from the *http* package *doc.go* file.

::

   // Copyright 2011 The Go Authors. All rights reserved.
   // Use of this source code is governed by a BSD-style
   // license that can be found in the LICENSE file.

   /*
   Package http provides HTTP client and server implementations.

   Get, Head, Post, and PostForm make HTTP (or HTTPS) requests:

           resp, err := http.Get("http://example.com/")
           ...

   <...strip lines...>

   */
   package http

Many lines are stripped from the above example where its marked. As you
can see the above documentation file starts with copyright notice. But
the copyright notice is using single line comment multiple times. This
notice will be ignored when generating documentation. And the
documentation is written within multi-line comments. At the end of file,
the package name is also declared.

Publishing Packages
-------------------

Unlike other mainstream languages, Go doesn’t have a central package
server. You can publish your code directly to any version control system
repositories. Git is most widely used VCS used to publish packages, but
you can also use Mercurial or Bazaar.

Here are few example packages:

-  https://github.com/auth0/go-jwt-middleware

-  https://github.com/blevesearch/bleve

-  https://github.com/dgrijalva/jwt-go

-  https://github.com/elazarl/go-bindata-assetfs

-  https://github.com/google/jsonapi

-  https://github.com/gorilla/mux

-  https://github.com/jpillora/backoff

-  https://github.com/kelseyhightower/envconfig

-  https://github.com/lib/pq

-  https://github.com/pkg/errors

-  https://github.com/thoas/stats

-  https://github.com/urfave/negroni

You can use ``go get`` command to get these packages locally. Go 1.11
release introduced the module support. The modules can be used to manage
external dependant packages.

Module
------

A *module* is a collection of Go packages with well defined name module
and dependency requirements. Also, the module has reproducible builds.

Modules use semantic versioning. The format of version should
vMAJOR.MINOR.PATCH. For example, v0.1.1, v1.3.1, or v2.0.0. Note that
the *v* prefix is mandatory.

There should be a file named *go.mod* at the top-level directory. This
file is used to declare the name of the module and list dependencies.

| The minimal version selection algorithm is used to select the versions
  of all modules used in a build. For each module in a build, the
  version selected by minimal version selection is always the
  semantically highest of the versions explicitly listed by a require
  directive in the main module or one of its dependencies.
| You can see an explanation of the algorithm here:
| https://research.swtch.com/vgo-mvs

Creating a module
~~~~~~~~~~~~~~~~~

The Go tool has support for creating modules. You can use the *mod*
command to manage module.

To initialize a new module, use *mod init* command with name of module
as argument. Normally the name will be same as the publishing location.
Here is an example:

::

   mkdir hello
   cd hello
   go mod init github.com/baijum/hello

In the above example, the name is given as ``github.com/baijum/hello``.
This command is going to create file named ``go.mod`` and the content of
that file will be like this:

::

   module github.com/baijum/hello

   go 1.20

As of now there are no dependencies. That is the reason the ``go.mod``
file doesn’t list any dependencies.

Let’s create a ``hello.go`` with ``github.com/lib/pq`` as a dependency.

::

   package main

   import (
       "database/sql"

       _ "github.com/lib/pq"
   )

   func main() {
       sql.Open("postgres",
           "host=lh port=5432 user=gt dbname=db password=secret")
   }

Note: The imported package is the driver used by the *sql* package.

You will see the ``mod.mod`` updated with dependency.

::

   module github.com/github.com/hello

   go 1.20

   require github.com/lib/pq v1.10.9

There will be another auto-generated file named ``go.sum`` which is used
for validation. You can commit both these files to your version control
system.

Moving Type Across Packages
---------------------------

When you refactoring a large package into smaller package or separating
certain features into another package, moving types will be required.
But moving types will be difficult as it may introduce backward
compatibility and it may affect existing consumer packages. Go has
support for type aliases to solve this problem. Type alias can be
declared as given in this example:

::

   type NewType = OldType

Type alias can be removed once all the dependant packages migrate to use
import path.

Exercises
---------

**Exercise 1:** Create a package named ``rectangle`` with exported
functions to calculate area and perimeter for the given rectangle.

**Solution:**

The name of the package could be ``rectangle``. Here is the program:

::

   // Package rectangle provides functions to calculate area
   // and perimeter for rectangle.
   package rectangle

   // Area calculate area of the given rectangle
   func Area(width, height float64) float64 {
        return width * height
   }

   // Perimeter calculate area of the given rectangle
   func Perimeter(width, height float64) float64 {
        return 2 * (width + height)
   }

As you can see above, the source file starts with package documentation.
Also you can see the documentation for all exported functions.

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Create a package with 3 source files and another *doc.go*
for documentation. The package should provide functions to calculate
areas for circle, rectangle, and triangle.

Summary
-------

This chapter explained the Go package in detail. Package is one of
building block of a reusable Go program. This chapter explained about
creating packages, documenting packages, and finally about publish
packages. The chapter also covered modules and its usage. Finally it
explained moving types across packages during refactoring.
