package main

import (
	"go/ast"
	"go/token"
	"log"
)

func checkSelectPattern(file *ast.File, fset *token.FileSet) {
	ast.Inspect(file, func(n ast.Node) bool {
		// find for loop statement
		forStmt, ok := n.(*ast.ForStmt)
		if !ok {
			return true
		}

		// iterate over statements in loop
		for _, stmt := range forStmt.Body.List {
			// check if it's a select statement
			selectStmt, ok := stmt.(*ast.SelectStmt)
			if !ok {
				continue
			}

			// iterate through the cases in the statement
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
						log.Printf("Warning: use 'msg, closed := <-msgCh' instead at %s",
							fset.Position(assignStmt.Pos()))
					}
				}
			}
		}

		return true
	})
}
