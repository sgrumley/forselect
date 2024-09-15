# Go Channel Select Linter
## Overview
This is a simple custom Go linter designed to detect specific usage of the select statement inside a for loop, where channels are received without checking if the channel is closed. A receive [operation](https://go.dev/ref/spec#Receive_operator) on a closed channel can always proceed immediately, yielding the element type's zero value after any previously sent values have been received. The linter raises a warning when the following pattern is found:

go
```go
for {
	select {
	case msg := <-msgCh:
		// Do something
	}
}
```

The linter suggests adding the second variable to handle the channel closing while the for statement is still running.

```go
for {
	select {
	case msg, closed := <-msgCh:
		// Handle closed channel properly
	}
}
```

## How It Works
The linter parses the provided Go source file.
It inspects the Abstract Syntax Tree (AST) of the file and looks for for loops that contain a select statement with a channel receive operation.
If the channel receive operation does not use the comma-ok pattern (msg, closed := <-msgCh), the linter issues a warning indicating that the preferred pattern is to use msg, closed := <-msgCh.
## Requirements
Go 1.18 or higher

## Installation
Clone the repository and build the linter:

```bash
git clone https://github.com/your-repo/go-select-linter.git
cd go-select-linter
go build -o linter linter.go
```

This will create an executable named linter.

## Usage
To use the linter, provide the path to the Go file you want to lint:

```bash
./linter yourfile.go
The linter will scan the file and print warnings in the following format if it detects any issues:
```


`Warning: use 'msg, closed := <-msgCh' instead at yourfile.go:line:column`

## Example
Given the following Go code:

```go
package main

func main() {
	msgCh := make(chan string)

	go func() {
		for {
			select {
			case msg := <-msgCh:
				// Handle message
			}
		}
	}()
}
```

When you run the linter:

```bash

./linter main.go
```
You will receive the following warning:

`Warning: use 'msg, closed := <-msgCh' instead at main.go:8:9`

