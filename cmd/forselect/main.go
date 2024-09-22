package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"forselect/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
