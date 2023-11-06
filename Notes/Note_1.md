Multithreading - allows multiple things to do at once. ex. google docs where multiple users are making changes at the same time on the same document. ex. watching youtube video, adding a comment etc.
Concurrency - while booking ticket no double booking should happeen when multiple users are booking at the same time. Python and node does not have built in concurrency. But cpp , java, go have concurrency.
Go is multi threaded , concurrent programming language.

We can use this for:
- building performant applications
- for building distributed applications
- simple and readable syntax
- efficieny and safety
- Go is used on server side - web applications, database services, microservices
- it is compiled language and its binary can be run on multiple paltforms.
- automation and cli based applications.


We need go compiler for executing.


================================

File structure of go file.

after creating a go file , initialze the go file with the command `go mod init <name-of-application/folder>`
everything in go is organized through packages. Even when we create own application we need to define it in package.
`package main`.
To do declaration : it is like an entry point of our application , we need our main starting point of execution.
*A program can have only 1 'main function' as only one entrypoint is possible.*
'we have to import the pckages in go implicitly.'

** diiference in print and println
In Go, both `print` and `println` are functions used to output text, but there are some differences between them:

1. **`print` function:**
   - The `print` function is used to print the arguments passed to it without any separator between the arguments.
   - Example:
     ```go
     package main
     
     import "fmt"
     
     func main() {
         print("Hello", "World") // Output: HelloWorld
     }
     ```

2. **`println` function:**
   - The `println` function is used to print the arguments passed to it with a space between the arguments and a newline character at the end.
   - Example:
     ```go
     package main
     
     import "fmt"
     
     func main() {
         println("Hello", "World")
         // Output:
         // Hello World
     }
     ```

In general, it is recommended to use `fmt.Print` and `fmt.Println` functions from the `fmt` package instead of `print` and `println` for more flexibility and formatting options. The `fmt` package provides functions for formatted I/O operations, allowing you to format strings, numbers, and other data types according to your needs.


=========================================

*Variables and Constants*
suppose there are names of users and we want to write their names like : thank you for booking ticket x , see you around x , here is tyour pass x..
updates everytime.
In golang , we get message saying that we defined a variabble but did not use it. ye hai kyunki dead code na ho aage chalke.
constant apna immutable hai , ek bar assign ki dobara change nhi hoga.

*Output ko format kaise krnma hai ?*  ---> using 'printf'


======================================

*Data types*
- strings
- integers
- float
- booleans
- map
- arrays
since go is statically types language so we need to tell go what type is the data of.
*But go can imply based on the value assigned to it.*










