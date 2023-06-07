Concurrency
===========

   *Do not communicate by sharing memory; instead, share memory by
   communicating.* — Effective Go

If you observe, you could see many things happening around you at any
given time. This is how the world function - the train is gently moving,
passengers talking each other, farmers working in the field and many
other things are happening simultaneously. We can say, the world we live
in function concurrently.

Go has built-in concurrency features with syntax support. The Go
concurrency is inspired by a paper published in 1978 by Tony Hoare. The
paper title is *Communicating sequential processes*  [1]_.

Go has some new terminologies and keywords related to concurrent
programming. The two important words are *goroutine* and *channel*. This
chapter will go through these concepts and walk through some examples to
further explain concurrency in Go.

The Go runtime is part of the executable binary created when compiling
any Go code. The Go runtime contains a garbage collector and a scheduler
to manage lightweight threads called Goroutines. Goroutine is a
fundamental abstraction to support concurrency. Goroutine is an
independently executing part of the program. You can invoke any number
of goroutines and all of them could run concurrently.

Goroutines can communicate to each other via typed conduits called
channels. Channels can be used to send and receive data.

Goroutine
---------

Goroutine is like a process running in the background. A function with
*go* keyword as prefix starts the goroutine. Any function including
anonymous function can be invoked with *go* keyword. In fact, the *main*
function is a special goroutine invoked during the starup of any program
by the Go runtime.

To understand the Goroutine better let’s look at a simple program:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
       "time"
   )

   var msg string

   func setMessage() {
       msg = "Hello, World!"
   }

   func main() {
       go setMessage()
       time.Sleep(1 * time.Millisecond)
       fmt.Println(msg)
   }

In the above program, *setMessage* function is invoked as a goroutine in
line no 15 using the *go* keyword. If you run this program, you will get
the hello world message printed. If you change the sleep time to Zero,
the message will not be printed. This is because, the program exits when
main function completes execution. And in this case, since *setMessage*
is called as a goroutine, it goes to background and main goroutine
execution continues. In the earlier case when the time sleep was 1
second, the goroutine gets some time to execute before main completed.
That’s why the *msg* value is set and printed.

Channels
--------

Multiple goroutines can communicate using channels. Channels can be used
to send and receive any type of values. You can send and receive values
with this channel operator: ``<-``

This is how to declare a channel of *int* values:

::

   ch := make(chan int)

To send a value to *ch* channel:

::

   ch <- 4

To receive a value from ``ch`` channel and assign to a variable:

::

   v := <-ch

You can also receive value without really assigning:

::

   <-ch

Sending and receiving values from channels becomes a blocking operation.
So, if you try to receive value from a channel, there should be some
other part of the code which sends a value this channel. Until a value
sends to the channel, the receiving part of the code will block the
execution.

Here is an example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
   )

   var c = make(chan int)
   var msg string

   func setMessage() {
       msg = "Hello, World!"
       c <- 0
   }

   func main() {
       go setMessage()
       <-c
       fmt.Println(msg)
   }

In the above example, an int channel is assigned to a global variable
named ``c``. In line number 17, immediately after calling goroutines,
channel is trying to receive a value. This becomes a blocking operation
in the ``main`` goroutine. In line number 12, inside the ``setMessage``
function, after setting a value for ``msg``, a value is send to the
``c`` channel. This will make the operation to continue in the ``main``
goroutine.

Waitgroups
----------

Go standard library has a *sync* package which provides few
synchronization primitives. One of the mechanism is *Waitgroups* which
can be used to wait for multiple goroutines to complete. The ``Add``
function add the number of goroutines to wait for. At the end of these
goroutines call ``Done`` function to indicate the task has completed.
The ``Wait`` function call, block further operations until all
goroutines are completed.

Here is a modified version of the previous example using *Waitgroups*.

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
       "sync"
   )

   var msg string
   var wg sync.WaitGroup

   func setMessage() {
       msg = "Hello, World!"
       wg.Done()
   }

   func main() {
       wg.Add(1)
       go setMessage()
       wg.Wait()
       fmt.Println(msg)
   }

In the above example, the ``Add`` method at line number 17 make one item
to wait for. The next line invoke the goroutine. The line number 19, the
``Wait`` method call blocks any further operations until goroutines are
completed. The previous line made goroutine and inside the goroutine, at
the end of that goroutine, there is a ``Done`` call at line number 13.

Here is another example:

.. code-block:: go
   :linenos:

   package main

   import (
       "fmt"
       "sync"
       "time"
   )

   func someWork(i int) {
       time.Sleep(time.Millisecond * 10)
       fmt.Println(i)
   }

   func main() {
       var wg sync.WaitGroup
       for i := 0; i < 5; i++ {
           wg.Add(1)
           go func(j int) {
               defer wg.Done()
               someWork(j)
           }(i)
       }
       wg.Wait()
   }

Select
------

The *select* is a statement with some similarity to *switch*, but used
with channels and goroutines. The *select* statement lets a goroutine
wait on multiple communication operations through channels.

Under a *select* statement, you can add multiple cases. A select
statement blocks until one of its case is available for run – that is
the channel has some value. If multiple channels used in cases has value
readily avaibale, select chooses one at random.

Here is an example:

::

   package main

   import "time"
   import "fmt"

   func main() {

       c1 := make(chan string)
       c2 := make(chan string)

       go func() {
           time.Sleep(time.Second * 1)
           c1 <- "one"
       }()
       go func() {
           time.Sleep(time.Second * 2)
           c2 <- "two"
       }()

       for i := 0; i < 2; i++ {
           select {
           case msg1 := <-c1:
               fmt.Println("received", msg1)
           case msg2 := <-c2:
               fmt.Println("received", msg2)
           }
       }
   }

Buffered Channels
-----------------

Buffered channels are channels with a given capacity. The capacity is
the size of channel in terms of number of elements. If the capacity is
zero or absent, the channel is unbuffered. For a buffered channel
communication succeeds only when both a sender and receiver are ready.
Whereas for a buffered channel, communication succeeds without blocking
if the buffer is not full (sends) or not empty (receives).

The capacity can be given as the third argument to make function:

::

   make(chan int, 100)

Consider the below example:

::

   package main

   import "fmt"

   func main() {
       ch := make(chan string, 2)
       ch <- "Hello"
       ch <- "World"
       fmt.Println(<-ch)
       fmt.Println(<-ch)
   }

The *ch* channel is a buffered channel, this makes it possible to send
value without any receiver present.

Channel Direction
-----------------

When declaring a function with channels as input parameters, you can
also specify the direction of the channel. The direction of channel
declares whether it can only receive or only send values. The channel
direction helps to increases the type-safety of the program.

Here is an example:

.. code-block:: go
   :linenos:

   package main

   import "fmt"

   func sendOnly(name chan<- string) {
       name <- "Hi"
   }

   func receiveOnly(name <-chan string) {
       fmt.Println(<-name)
   }

   func main() {
       n := make(chan string)

       go func() {
           fmt.Println(<-n)
       }()

       sendOnly(n)

       go func() {
           n <- "Hello"
       }()

       receiveOnly(n)
   }

In the above example, the ``sendOnly`` function define a channel
variable which can be only used for sending data. If you tried to read
from that channel within that function, it’s going to be compile time
error. Similary the ``receiveOnly`` function define a channel variable
which can be only user for receive data. You cannot send any value to
that channel from that function.

Lazy Initialization Using sync.Once
-----------------------------------

The sync package provide another struct called Once which is useful for
lazy initialization.

Here is an example:

::

   import (
       "sync"
   )

   type DB struct{}

   var db *DB
   var once sync.Once

   func GetDB() *DB {
       once.Do(func() {
           db = &DB{}
       })
       return db
   }

If the above GetDB function is called multiple times, only once the DB
object will get constructed.

Exercises
---------

**Exercise 1:** Write a program to download a list of web pages
concurrently using Goroutines.

Hint: Use this tool for serving junk content for testing:
https://github.com/baijum/lipsum

**Solution:**

::

   package main

   import (
           "io/ioutil"
           "log"
           "net/http"
           "net/url"
           "sync"
   )

   func main() {
           urls := []string{
                   "http://localhost:9999/1.txt",
                   "http://localhost:9999/2.txt",
                   "http://localhost:9999/3.txt",
                   "http://localhost:9999/4.txt",
           }
           var wg sync.WaitGroup
           for _, u := range urls {
                   wg.Add(1)
                   go func(u string) {
                           defer wg.Done()
                           ul, err := url.Parse(u)
                           fn := ul.Path[1:len(ul.Path)]
                           res, err := http.Get(u)
                           if err != nil {
                                   log.Println(err, u)
                           }
                           content, _ := ioutil.ReadAll(res.Body)
                           ioutil.WriteFile(fn, content, 0644)
                           res.Body.Close()
                   }(u)
           }
           wg.Wait()
   }

Additional Exercises
~~~~~~~~~~~~~~~~~~~~

Answers to these additional exercises are given in the Appendix A.

**Problem 1:** Write a program to watch log files and detect any entry
with a particular word.

Summary
-------

This chapter explained concurrency features of Go. Based on your
problem, you can choose channels or other synchronization techniques.
This chapter covered goroutines and channels usage. It covered
Waitgroups, Select statement. It also covered buffered channels, channel
direction. The chapter also touched upon *sync.Once* function usage.

.. [1]
   http://usingcsp.com
