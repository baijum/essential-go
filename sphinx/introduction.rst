Introduction
============

   *I try to say everything at least three times: first, to introduce
   it; second, to show it in context; and third, to show it in a
   different context, or to review it.* — Robert J. Chassell (An
   Introduction to Programming in Emacs Lisp)

Having proficiency in computer programming empowers you to tackle a wide
range of real-world challenges. Programming encompasses a process that
begins with formulating a computing problem and culminates in the
creation of computer programs or software. It is an integral part of a
broader software development process.

Programming entails analyzing, designing, and implementing software
solutions. Coding is the active process of implementing the software. At
times, coding involves utilizing multiple programming languages and
incorporating other technologies. Acquiring proficiency in a programming
language is an essential aspect of the skill set required for computer
programming and software development as a whole.

By utilizing a programming language, we construct instructions for
computing machines such as desktop computers, laptops, and mobile
phones.

Today, there exists a multitude of programming languages with diverse
features. Choosing the appropriate programming language for a given
problem is crucial. Certain languages are better suited for specific
challenges. This book serves as an introduction to the Go programming
language, offering valuable insights to aid in making informed language
selection decisions for your projects. If you have already chosen to
work with Go and seek to enhance your understanding, this book provides
an excellent introduction to Go programming.

Go, an open-source language developed by Google with significant
contributions from the community, has emerged as a popular choice. The
project was initiated in 2007 by Robert Griesemer, Rob Pike, and Ken
Thompson and was subsequently released as open-source software by Google
in November 2009. Go has gained widespread adoption across diverse
problem domains, being embraced by numerous organizations.

The design of the Go programming language draws inspiration from various
programming languages, such as C, Python, and occam. Go programs can run
on many operating systems including GNU/Linux, Windows and Mac OS X.

Before diving into learning Go programming, there are certain
preparations that you need to make. The following section will provide
an explanation of these necessary preparations.

This book is tailored for beginners who are eager to learn programming.
While no prior programming experience is required, readers are expected
to have a basic understanding of computers. The book comprehensively
covers all the major topics in the Go programming language, with each
chapter focusing on one or two key areas. Additionally, minor topics are
introduced throughout the book at appropriate intervals.

Preparations
------------

Learning a programming language, such as Go, is akin to learning a
natural language like English, Chinese, or Spanish. It involves more
than just understanding the alphabet and grammar specific to that
language. To become a proficient programmer, there are several other
important concepts and skills that you need to acquire. While this book
primarily focuses on Go programming, we will also explore other
essential topics that are crucial for your overall programming journey.

In this section, we will delve into the installation of the Go compiler
and the setup of a suitable development environment. These steps are
necessary to ensure that you have the necessary tools in place to begin
your Go programming experience.

Text editors such as `Vim <https://vim-bootstrap.com>`__ and VS Code are
popular choices for writing Go programs. The files created using these
text editors are referred to as source files. To transform these source
files into executable programs, the Go compiler comes into play. Once
the executable program is generated, you can run it and obtain the
desired output. Therefore, it is essential to have a text editor and the
Go compiler properly installed on your computer system.

Utilizing a source code management system, such as Git, is highly
advantageous. It is strongly recommended to keep all your code under
version control. There are several public code hosting services
available, such as GitHub, Bitbucket, and GitLab, where you can store
your code examples securely. By leveraging these platforms, you can
easily manage your code, collaborate with others, and track changes
effectively.

To install the Go compiler, please refer to the instructions specific to
your operating system. If you encounter any difficulties during the
installation process, don’t hesitate to seek assistance from your
friends or fellow programmers. Once the Go compiler is installed, you
can proceed to write a simple Go program and run it to validate that the
installation was successful.

Linux Installation
~~~~~~~~~~~~~~~~~~

The Go project offers binaries for various major operating systems,
including GNU/Linux. If you are using a 64-bit GNU/Linux system, you can
find the appropriate binaries for your platform by visiting the
following link: https://go.dev/dl.

To download and install the Go compiler on a 64-bit GNU/Linux system,
please follow the instructions below. Before proceeding, make sure that
the Go compiler is not already installed by running the ``go`` command.
If the command returns ``command not found...``, you can proceed with
the installation process.

To execute these commands, you need *root* access or you can use the
``sudo`` command. If you are unfamiliar with the process, you can seek
assistance from someone who has the necessary knowledge.

::

   cd /tmp
   wget https://go.dev/dl/go1.x.linux-amd64.tar.gz
   tar -C /usr/local -zxvf go1.x.linux-amd64.tar.gz

The first line ensures that the current working directory is set to the
``/tmp`` directory.

In the second line, replace the version number with the appropriate
version available on the `Go downloads <https://go.dev/dl>`__ website.
This command will download the 64-bit binary for GNU/Linux using the
*wget* command, which is a command line download manager. Alternatively,
you can use *curl* or any other download manager of your choice to
download the tar ball.

The third line extracts the downloaded tar ball into the
``/usr/local/go`` directory.

You can now exit the *root* user or stop using *sudo*, depending on the
method you chose.

By default, Go packages are installed under the ``$HOME/go`` directory.
However, you can override this directory by setting the ``GOPATH``
environment variable. Any binaries installed using the ``go install``
command will be placed in the ``$GOPATH/bin`` directory.

To include the new binary locations in the ``PATH`` environment
variable, you can open the ``$HOME/.bashrc`` file in a text editor and
add the following lines at the bottom:

::

   export PATH=$HOME/go/bin:/usr/local/go/bin:$PATH

Windows Installation
~~~~~~~~~~~~~~~~~~~~

For Windows users, there are separate installers (MSI files) available
for both 32-bit and 64-bit versions of the Go programming language. The
32-bit version MSI file follows a naming convention similar to:
*go1.x.y.windows-386.msi*, where *x.y* represents the current version
number. Similarly, the 64-bit version MSI file is named as
*go1.x.y.windows-amd64.msi*. It is important to replace *x.y* with the
actual version number when downloading the appropriate installer file.

You can download the installers (MSI files) for the Go programming
language from the following website: https://go.dev/dl

This website provides the official distribution of Go, where you can
find the installers for various operating systems, including Windows.
Simply navigate to the provided link and select the appropriate MSI file
for your Windows version (32-bit or 64-bit) to begin the download.

Once you have downloaded the Go installer file (MSI), you can proceed
with the installation by following these steps:

#. Locate the downloaded MSI file on your computer.

#. Double-click on the MSI file to open it.

#. The installer will prompt you with a setup wizard that guides you
   through the installation process.

#. Follow the instructions provided by the setup wizard.

#. During the installation, you will be asked to choose the destination
   directory for the Go compiler.

#. By default, the installer will suggest the ``C:\Go`` directory as the
   installation location.

#. You can either accept the default directory or choose a different
   location on your system.

#. Once you have selected the installation directory, click on the
   *Install* or *Next* button to proceed.

#. The installer will then copy the necessary files to the chosen
   directory.

#. After the installation is complete, you can close the installer.

At this point, the Go compiler should be successfully installed on your
system, with the relevant files located in the specified installation
directory (e.g., ``C:\Go`` by default).

Additionally, the installer automatically adds the ``C:\ Go\bin``
directory to the system PATH environment variable. However, in order for
this change to take effect, you may need to restart any open command
prompts or terminals on your system. This ensures that the Go
executables can be accessed from any command prompt or terminal window
without specifying the full path to the ``bin`` directory.

To download third-party packages from websites like GitHub.com, it is
recommended to create a directory where these packages will be stored.
You can create a directory named ``mygo`` at the ``C:\`` directory by
following these steps:

#. Open a command prompt or terminal.

#. Type the following command and press Enter:

::

   mkdir C:\mygo

This will create a new directory named ``mygo`` directly under the
``C:\`` directory. You can use this directory to store and manage the
third-party packages you download for your Go projects.

To set the *GOPATH* environment variable to point to the location where
you created the ``mygo`` directory (in this case, ``C:\mygo``), follow
these steps:

#. Right-click on the *This PC* or *My Computer* icon on your desktop
   and select *Properties*.

#. In the System window, click on *Advanced system settings* on the
   left-hand side.

#. In the System Properties window, click on the *Environment Variables*
   button.

#. In the Environment Variables window, under the *User variables*
   section, click on the *New* button.

#. Enter *GOPATH* as the variable name.

#. Enter the path to the ``mygo`` directory (in this case, ``C:\mygo``)
   as the variable value.

#. Click *OK* to save the changes.

Once you have set the GOPATH environment variable, it will point to the
specified location, allowing Go to find and use the packages stored in
the ``mygo`` directory.

Verifying Installation
~~~~~~~~~~~~~~~~~~~~~~

To verify that you have successfully installed Go, run the following
command:

::

   go version

This command will display the installed Go version if the installation
was successful. The output should looks like somthing like this:

::

   go version go1.20.4 linux/amd64

Please note that the version number and platform may vary depending on
the Go version and the operating system you are using.

Hello World!
~~~~~~~~~~~~

It is a common tradition in programming education to introduce a *Hello
World* program as the first program. This program typically prints the
message ``Hello World`` to the console when executed.

Here is a simple *Hello World* program. You can type the following
source code into your favorite text editor and save it as ``hello.go``:

::

   package main

   import "fmt"

   func main() {
       fmt.Println("Hello, World!")
   }

Make sure to save the file with the extension *.go* to indicate that it
is a Go source code file.

Once you have saved the above source code into a file, you can open your
command line program (such as Terminal or Command Prompt). Then,
navigate to the directory where you saved the program code using the
``cd`` command. For example:

::

   cd /path/to/directory

Replace ``/path/to/directory`` with the actual path to the directory
where you saved the ``hello.go`` file.

Once you are in the correct directory, you can run the program by typing
the following command:

::

   go run hello.go

Press Enter to execute the command. The program will be compiled and
executed, and you should see the output ``Hello, World!`` displayed in
the command line.

If you see the output as ``Hello, World!``, congratulations! You have
successfully installed the Go compiler. In fact, the ``go run`` command
compiled your code into an executable format and then executed the
program. The next chapter will provide more detailed explanations about
this example and delve further into the concepts of Go programming.

Using Git
~~~~~~~~~

It is important to be comfortable using a source code management system,
and Git is highly recommended. Creating an account on GitHub and
publishing your example code there would be beneficial. If you are new
to Git and have no prior experience, dedicating 2 to 3 days to learn Git
would be worthwhile.

Using Command Line
~~~~~~~~~~~~~~~~~~

It is essential to be comfortable using command line interfaces such as
GNU Bash or PowerShell. There are numerous online tutorials available on
the Internet to learn shell commands. If you are unfamiliar with command
line usage, it is recommended to allocate a few days (around 3 to 4
days) to learn and familiarize yourself with the command line
environment.

Organization of Chapters
------------------------

The book is structured into the following chapters, which can be read in
the suggested order. The first six chapters are designed to be read
sequentially, while the remaining chapters can be read in any order you
prefer.

Chapter 2: Quick Start
   | 
   | This chapter serves as a tutorial introduction to the Go
     programming language. It covers essential topics that form the
     foundation of Go programming, including data types, variables,
     comments, for loops, range clauses, if statements, functions,
     operators, slices, and maps. By studying these topics, readers will
     gain a solid understanding of the fundamentals of Go programming.

Chapter 3: Control Structures
   | 
   | This chapter covers the different control structures available in
     the Go programming language, including goto statements, if
     conditions, for loops, and switch cases. Each of these topics is
     explained in detail, providing a comprehensive understanding of how
     to effectively use these control structures in Go programming.

Chapter 4: Data Structures
   | 
   | This chapter explores data structures in the Go programming
     language. It begins by discussing arrays and then delves into
     slices, which are a more versatile data structure built on top of
     arrays. The chapter also explains how to define custom data types
     using existing primitive types and introduces the use of structs
     for creating more complex custom data types. Pointers are also
     covered, along with structs, slices, and maps in the language.

Chapter 5: Functions
   | 
   | This chapter provides a comprehensive explanation of functions in
     Go. It covers various aspects such as sending input parameters and
     returning values. The chapter also explores variadic functions and
     anonymous functions. Additionally, there is a brief introduction to
     methods in Go.

Chapter 6: Objects
   | 
   | This chapter delves into the concept of objects and interfaces in
     Go and their practical applications. Interfaces hold significant
     importance in Go as they contribute to robust design. The chapter
     covers the concept of empty interfaces and provides an overview of
     pointer receivers and their significance. Additionally, the chapter
     explores type assertions and type switches in Go.

Chapter 7: Concurrency
   | 
   | In this chapter, the concurrency features of Go are explained in
     detail. Depending on the problem at hand, you can choose between
     channels and other synchronization techniques. The chapter covers
     the usage of goroutines and channels, highlighting their importance
     in concurrent programming. It also explores topics such as
     Waitgroups, Select statements, buffered channels, and channel
     direction. Additionally, the chapter provides an introduction to
     the usage of the sync.Once function.

Chapter 8: Packages
   | 
   | In this chapter, the concept of Go packages is thoroughly
     explained. Packages serve as fundamental building blocks for
     creating reusable Go programs. The chapter covers various aspects
     such as creating packages, documenting packages, and the process of
     publishing packages. It also delves into the topic of modules and
     their usage in managing dependencies. Additionally, the chapter
     provides insights on moving types across packages during the
     refactoring process.

Chapter 9: Input/Output
   | 
   | In this chapter, the various input/output functionalities in Go are
     explored. The chapter covers topics such as handling command line
     arguments and interactive input. It introduces the usage of the
     flag package for handling command line options and arguments.
     Additionally, the chapter provides insights into various string
     formatting techniques used in Go.

Chapter 10: Testing
   | 
   | In this chapter, the process of writing tests using the testing
     package in Go is explained. The chapter covers important concepts
     such as marking tests as failures, logging, skipping tests, and
     running tests in parallel. Additionally, the chapter briefly
     introduces the concept of sub-tests, providing a glimpse into its
     usage.

Chapter 11: Tooling
   | 
   | In this chapter, the Go tool is introduced and all its commands are
     explained in detail. Practical examples are provided for each
     command to illustrate their usage. The chapter covers important
     commands such as building and running programs, running tests,
     formatting code, and displaying documentation. Additionally, the
     chapter briefly mentions a few other useful tools that can enhance
     your development workflow.

In addition to the solved exercises, each chapter includes additional
problems for further practice. The answers to these additional problems
can be found in Appendix A.

Lastly, the book concludes with an index at the end, which serves as a
helpful reference for locating specific topics and concepts throughout
the book.

Suggestions to Use this Book
----------------------------

It is important to follow the instructions provided in this chapter to
set up your system with the Go compiler and the necessary environment.
If you encounter any difficulties during the setup process, don’t
hesitate to seek assistance from your friends or colleagues.
Additionally, utilizing a source code management system like Git will
greatly facilitate the management of your code. You can use it to write
exercises, solve additional problems, and keep your code under version
control for easy tracking and collaboration.

I would recommend against simply copying and pasting code from the book.
Instead, I encourage you to actively type out each example provided. By
manually typing the code, you will gain a better understanding of the
syntax and structure of the language, which will help you become more
familiar with it more quickly. Additionally, typing out the code will
improve your muscle memory and reinforce your learning process. So, take
the time to engage with the code actively and type it out yourself for a
more effective learning experience.

It is recommended to read the first six chapters of the book in order,
starting from the Introduction to Interfaces. These initial chapters lay
the foundation and cover important concepts that are built upon in the
later chapters. Reading them sequentially will provide a logical
progression of knowledge and understanding.

However, once you have completed the first six chapters, the remaining
chapters can be read in any order based on your specific interests or
needs. Each chapter in the later part of the book focuses on a specific
topic or aspect of Go programming, and they are designed to be
relatively independent of each other. Feel free to explore the chapters
that align with your interests or are relevant to your current
programming goals.

By following this approach, you will establish a solid understanding of
the fundamentals through the initial chapters and then have the
flexibility to delve into specific topics of your choice in the later
chapters.

Summary
-------

In this chapter, we provided an introduction to the Go programming
language. We covered essential topics that are crucial for becoming a
proficient programmer. Our discussion touched upon various aspects that
are important to understand in order to develop strong programming
skills.

We emphasized that learning Go programming involves more than just
understanding the language syntax. To become a good programmer, it is
necessary to grasp a wide range of concepts and skills. While our focus
in this book is on Go programming, we acknowledged that there are
several other topics and areas of knowledge that are valuable for
programmers to explore.

Furthermore, we highlighted the significance of installing the Go
compiler and setting up a suitable development environment to facilitate
the learning process. We also encouraged the use of source code
management systems like Git to manage and organize your code
effectively.

By introducing these essential elements in this chapter, we aim to
provide a solid foundation for your journey in learning Go programming.

Following the introduction to Go programming, we proceeded to discuss
the organization of chapters in this book. We outlined the structure and
sequence of topics covered, emphasizing that the first six chapters
should be read in order. These initial chapters serve as a foundation
for the subsequent content, forming the building blocks for a
comprehensive understanding of Go programming.

Additionally, we provided some helpful suggestions on how to effectively
utilize this book for learning Go. These recommendations are intended to
enhance your learning experience and maximize your understanding of the
language and its concepts.

Moving forward, the next chapter will provide a quick start to
programming with the Go language. It will offer practical examples and
hands-on exercises to help you get started with writing Go programs.
This chapter aims to provide a smooth transition from theoretical
concepts to practical application, enabling you to gain firsthand
experience with Go programming.
