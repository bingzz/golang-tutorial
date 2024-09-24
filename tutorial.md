# GOLANG TUTORIAL

## 6 Main Points

1. Static typed language -> cannot change to another type
2. Strongly typed language -> cannot combine different types into a single variable
3. Compiled language -> can build into machine code
4. Fast compilation time -> convert to executable
5. Built in concurrency -> does not require external packages
6. Simplicity -> simple syntaxes and less lines of codes, less garbage

## Structure

* Modules -> Collection of packages
* Packages -> Collection of go files

## Commands

1. *go mod init* -> initialize a module

## Variable types

* int
* uint8 -> better used for 0-255 values
* uint -> unsigned (meaning positive values only)
* float32 ->
* float64 -> more accurate, but more memory

## Functions and Control Structures

## Arrays

* Fixed length
* Same type
* Indexable
* Contiguous in Memory

Ampersand (&) symbol -> memory location

Asterisk (*) symbol -> Pointer

## Maps

* Set of key value pairs to look up a value by its key

## Character Encoding

* ASCII (7 bits) -> 128 chars
* UTF-32(32 bits) -> 1,114,112 chars

## GoRoutines

* Calls multiple functions concurrently

CONCURRENCY != PARALLEL EXECUTION

Concurrency is to execute the function simultaneously on a single CPU while the parallel execution is executed simultaneously on different CPUs

The more CPU, the faster the execution

## Channels

* Work with Goroutines to pass around data
* Holds data
* Thread safe (prevent data erasure)
* Listen to Data
* Best combined with goroutines
