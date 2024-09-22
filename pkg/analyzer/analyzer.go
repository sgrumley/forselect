package analyzer

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "forselect",
	Doc:  "Checks that select statements in for loops check recieving channels are closed",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := func(node ast.Node) bool {
		// find for loop statement
		forStmt, ok := node.(*ast.ForStmt)
		if !ok {
			return true
		}

		// iterate over statements in for loop
		for _, stmt := range forStmt.Body.List {
			// check if it's a select statement
			selectStmt, ok := stmt.(*ast.SelectStmt)
			if !ok {
				continue
			}

			// iterate through the cases in the select statement
			for _, selectCase := range selectStmt.Body.List {
				commClause, ok := selectCase.(*ast.CommClause)
				if !ok {
					continue
				}

				// verify assignment statement
				assignStmt, ok := commClause.Comm.(*ast.AssignStmt)
				if !ok {
					continue
				}

				// check for walrus operator
				if assignStmt.Tok != token.DEFINE {
					continue
				}

				// check the the Right Hand Side is recieveing a channel
				if len(assignStmt.Rhs) == 1 {
					_, ok := assignStmt.Rhs[0].(*ast.UnaryExpr)
					if ok && len(assignStmt.Lhs) == 1 {
						pass.Reportf(node.Pos(), "Warning: use 'msg, closed := <-msgCh' instead")
					}
				}
			}
		}

		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}
	return nil, nil
}
