package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: linter <file.go>")
	}

	filename := os.Args[1]
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, filename, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("Failed to parse file: %s", err)
	}

	checkSelectPattern(file, fset)
}
