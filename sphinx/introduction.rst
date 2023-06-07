Introduction
============

   *I try to say everything at least three times: first, to introduce
   it; second, to show it in context; and third, to show it in a
   different context, or to review it.* — Robert J. Chassell (An
   Introduction to Programming in Emacs Lisp)

Computer programming skill helps you to solve many real-world problems.
Programming is a process starting from the formulation of a computing
problem to producing computer programs (software). Computer programming
is part of a more extensive software development process.

Programming involves analysis, design, and implementation of the
software. Coding is the activity of implementing software. Sometimes
coding involves more than one programming language and use of other
technologies. Learning a programming language is a crucial part of the
skill required for computer programming or software development in
general.

Using a programming language, we are preparing instructions for a
computing machine. The computing machine includes a desktop computer,
laptop, and mobile phone.

There are many programming languages in use today with different feature
sets. You should learn to pick the right programming language for the
problem at hand. Some languages are more suitable for specific problems.
This book provides an introduction to the Go programming language.
Studying the Go programming should help you to make the right decision
about programming language choice for your projects. If you have already
decided to use Go, this book should give an excellent introduction to Go
programming.

Go, also commonly referred to as Golang, is a general-purpose
programming language. Go was initially developed at Google in 2007 by
Robert Griesemer, Rob Pike, and Ken Thompson. Go was publicly released
as an open source software in November 2009 by Google.

Many other programming languages including C, Python, and occam has
inspired the design of the Go programming language. Go programs can run
on many operating systems including GNU/Linux, Windows and Mac OS X.

You require some preparations to learn Go programming. The next section
explains the preparations required.

This book is for beginners who want to learn to programme. The readers
are expected to have a basic knowledge of computers. This book covers
all the major topics in the Go programming language. Each chapter in
this book covers one or two major topics. However many minor topics are
introduced intermittently.

This chapter provides an introduction to the language, preparations
required to start practicing the exercises in this book, organizing code
and walk through of remaining chapters. Few suggestions for learning Go
using this book is given towards the end of this chapter.

Preparations
------------

Learning a natural language like English, Chinese, and Spanish is not
just understanding the alphabet and grammar of that particular language.
Similarly, there are many things that you need to learn to become a good
programmer. Even though this book is going to focus on Go programming;
now we are going to see some other topics that you need to learn. This
section also explains the installation of the Go compiler and setting up
the necessary development environment.

Text editors like Notepad, Notepad++, Gedit, and Vim can be used to
write Go programs. The file that you create using the text file is
called source file. The source file text is UTF-8 encoded. The Go
compiler creates executable programs from the source file. You can run
the executable program and get the output. So, you need a text editor
and Go compiler installed in your system.

Depending on your operating system, follow the instruction given below
to install the Go compiler. If you have difficulty following this, you
may get some help from your friends to do this step. Later we write a
simple Go program and run it to validate the steps.

You can use any text editor to write code. If you are not familiar with
any text editor, consider using Vim. You can bootstrap Vim configuration
for Go programming language from this webste: https://vim-bootstrap.com.

Using a source code management system like Git would be helpful. Keeping
all your code under version control is highly recommended. You can use a
public code hosting service like GitHub, Bitbucket, and GitLab to store
your examples.

Linux Installation
~~~~~~~~~~~~~~~~~~

Go project provides binaries for major operating systems including
GNU/Linux. You can find 64 bit binaries for GNU/Linux here:
https://go.dev/dl

The following commands will download and install Go compiler in a 64 bit
GNU/Linux system. Before performing these steps, ensure Go compiler is
not installed by running ``go`` command. If it prints *command not
found...*, you can proceed with these steps.

These commands must be run as *root* or through *sudo*. If you do not
know how to do it, you can get help from somebody else.

::

   cd /tmp
   wget https://go.dev/dl/go1.x.linux-amd64.tar.gz
   tar -C /usr/local -zxvf go1.x.linux-amd64.tar.gz

The first line ensure that current working directory is the ``/tmp``
directory.

You should change the version number in the second line and it’s going
to download the 64 bit binary for GNU/Linux. The ``wget`` is a command
line download manager. Alternatively you can use ``curl`` or any other
download manager to download the tar ball.

The third line extract the downloaded tar ball into ``/usr/local/go``
directory.

Now you can exit from the *root* user or stop using *sudo*.

By default Go packages are installed under ``$HOME/go`` directory. This
directory can be overridden using ``GOPATH`` environment variable. Any
binaries installed using ``go install`` and ``go get`` commands goes
into ``$GOPATH/bin`` directory.

You can also update PATH environment variable to include new binary
locations. Open the ``$HOME/.bashrc`` file in a text editor and enter
this lines at the bottom.

::

   export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH

Windows Installation
~~~~~~~~~~~~~~~~~~~~

There are separate installers (MSI files) available for 32 bit and 64
bit versions of Windows. The 32 bit version MSI file will be named like
this: *go1.x.y.windows-386.msi* (Replace ``x.y`` with the current
version). Similarly for 64 bit version, the MSI file will be named like
this: *go1.x.y.windows-amd64.msi* (Replace ``x.y`` with the current
version).

You can download the installers (MSI files) from here: https://go.dev/dl

After downloading the installer file, you can open the MSI file by
double clicking on that file. This should prompts few things about the
installation of the Go compiler. The installer place the Go related
files in the ``C:\Go`` directory.

The installer also put the ``C:\Go\bin`` directory in the system
``PATH`` environment variable. You may need to restart any open command
prompts for the change to take effect.

You also need to create a directory to download third party packages
from github.com or similar sites. The directory can be created at
``C:\mygo`` like this:

::

   C:\> mkdir C:\mygo

After this you can set ``GOPATH`` environment variable to point to this
location.

::

   C:\> go env -w GOPATH=C:\mygo

You can also append ``C:\mygo\bin`` into the ``PATH`` environment
variable.

If you do not know how to set environment variable, just do a Google
search for: *set windows environment variable*.

Hello World!
~~~~~~~~~~~~

It’s kind of a tradition in teaching programming to introduce a *Hello
World* program as the first program. This program normally prints a
*Hello World* to the console when running.

Here is our hello world program. You can type the source code given
below to your favorite text editor and save it as ``hello.go``.

::

   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }

Once you saved the above source code into a file. You can open your
command line program (bash or cmd.exe) then change to the directory
where you saved the program code and run the above program like this:

::

   $ go run hello.go
   Hello, World!

If you see the output as ``Hello, World!``, congratulations! Now you
have successfully installed Go compiler. In fact, the ``go run`` command
compiled your code to an executable format and then run that program.
The next chapter explains more about this example.

Using Git
~~~~~~~~~

You should be comfortable using a source code management system. As
mentioned above Git would be a good choice. You can create an account in
GitHub and publish your example code there. If you do not have any prior
experience, you can spend 2 to 3 days to learn Git.

Using Command Line
~~~~~~~~~~~~~~~~~~

You should be comfortable using command line interfaces like GNU Bash or
PowerShell. There are many online tutorials available in the Internet to
learn shell commands. If you do not have any prior experience, you can
spend few days (3 to 4 days) to learn command line usage.

Organization of Chapters
------------------------

The rest of the book is organized into the following chapters. It is
recommended to read the first six chapters in order – that’s until the
chapter on interfaces. The remaining chapters can be read in any order.

Chapter 2: Quickstart
   | 
   | This chapter provides a tutorial introduction to the language. It
     introduce few basic topics in Go programming language. The topics
     include Data Types, Variables, Comments, For Loop, Range Clause,
     If, Function, Operators, Slices, and Maps.

Chapter 3: Control Structures
   | 
   | This chapter cover the various control structures like *goto*, *if
     condition*, *for loop* and *switch case* in the language. It goes
     into details of each of these topics.

Chapter 4: Data Structures
   | 
   | This chapter cover data structures. The chapter starts with arrays.
     Then slices, the more useful data structure built on top of array
     is explained. Then we looked at how to define custom data types
     using existing primitive types. The struct is introduced which is
     more useful to create custom data types. Pointer is also covered.
     like *struct*, *slice* and *map* in the language.

Chapter 5: Functions
   | 
   | This chapter explained all the major aspects of functions in Go.
     The chapter covered how to send input parameters and return values.
     It also explained about variadic function and anonymous function.
     This chapter briefly also covered methods.

Chapter 6: Interfaces
   | 
   | This chapter explained the concept of interfaces and it’s uses.
     Interface is an important concept in Go. Understanding interfaces
     and properly using it makes the design robust. The chapter covered
     empty interface. Also, briefly explained about pointer receiver and
     its significance. Type assertions and type switches are also
     explained.

Chapter 7: Concurrency
   | 
   | This chapter explained concurrency features of Go. Based on your
     problem, you can choose channels or other synchronization
     techniques. This chapter covered goroutines and channels usage. It
     covered Waitgroups, Select statement. It also covered buffered
     channels, channel direction. The chapter also touched upon
     *sync.Once* function usage.

Chapter 8: Packages
   | 
   | This chapter explained the Go package in detail. Package is one of
     building block of a reusable Go program. This chapter explained
     about creating packages, documenting packages, and finally about
     publish packages. The chapter also covered modules and its usage.
     Finally it explained moving types across packages during
     refactoring.

Chapter 9: Input/Output
   | 
   | This chapter discussed about various input/output related
     functionalities in Go. The chapter explained using command line
     arguments and interactive input. The chaptered using *flag*
     package. It also explained about various string formatting
     techniques.

Chapter 10: Testing
   | 
   | This chapter explained writing tests using the *testing* package.
     It covered how to mark a test as a failure, logging, skipping, and
     parallel running. Also, it briefly touched upon sub-tests.

Chapter 11: Tooling
   | 
   | This chapter introduced the Go tool. All the Go commands were
     explained in detail. Practical example usage was also given for
     each command. The chapter coverd how to build and run programs,
     running tests, formatting code, and displaying documentation. It
     also touched upon few other handy tools.

In addition to the solved exercises, each chapter contains additional
problems. Answers to these additional problems are given in Appendix A.

And finally there is an index at the end of the book.

Suggestions to Use this Book
----------------------------

Make sure to setup your system with Go compiler and the environment as
explained in this chapter. If you are finding it very difficult, you may
get help from your friends to setup the environment. Use source code
management system like Git to manage your code. You can write exercises
and solve additional problems and keep it under version control.

I would suggest not to copy & paste code from the E-book. Rather you can
type every examples in this book. This will help you to familiarize the
syntax much quickly.

The first 6 chapters, that is from Introduction to Interfaces should be
read in order. The remaining chapters are based on the first 6 chapters.
So, the chapter 7 onward can be read in any order.

Summary
-------

This chapter provided an introduction to Go programming language. We
briefly discussed about topics required to become a good programmer.

Then we covered chapter organization in this book. And finally few
suggestions for readers are given. The next chapter provides a
quickstart to programming with Go language.
