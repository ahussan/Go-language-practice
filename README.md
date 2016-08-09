# Go-language-practice


##The Go Programming Language

The Go Programming Language (Go) is an open-source programming language sponsored by Google and created by Robert Griesemer, Rob Pike, and Ken Thompson that aims to simplify systems programming and to make programming fun again. 
Go is a relatively new language that is in constant development but it already has a considerable amount of industry support and is used in real systems at Google.

The notion of yet another programming language may seem weird to some people. But Go does bring something new to the table and its fundamental design principles make it different enough (and arguably, better enough) to justify its existence.

#Hello World!
No introduction to any language is complete without the canonical Hello World program:
```
package main

import "fmt"

func main() {

    fmt.Println("Hello world!")
}
```
The program defines a new package (main is always the package that contains the main function) then imports the input/output formatting package (fmt), defines its main function (which is the program’s entry point), and uses the Println function from the fmt package to print the string Hello world!.


For documentation about how to install and use Go, visit https://golang.org/ or load doc/install-source.html in your web browser.

