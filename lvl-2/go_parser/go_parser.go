package go_parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func Run(fileName, funcName string) {
	var count int64
	set := token.NewFileSet()
	astFile, err := parser.ParseFile(set, fileName, nil, 0)
	if err != nil {
		fmt.Println("Failed to parse package:", err)
		os.Exit(1)
	}

	var function *ast.FuncDecl
	for _, d := range astFile.Decls {
		if fn, isFn := d.(*ast.FuncDecl); isFn {
			if fn.Name.String() != funcName {
				continue
			}
			function = fn
			count = countGoStmt(fn.Body.List)
			break
		}
	}

	if function == nil {
		log.Fatalf("function %s does not exist in file %s", funcName, fileName)
	} else {
		fmt.Printf("There are %d go statements in function %s in file %s", count, funcName, fileName)
	}

}

func countGoStmt(stmts []ast.Stmt) int64 {
	var count int64
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.GoStmt); ok {
			count++
		} else if ifStmt, ok := stmt.(*ast.IfStmt); ok {
			count += countGoStmt(ifStmt.Body.List)
		} else if forStmt, ok := stmt.(*ast.ForStmt); ok {
			count += countGoStmt(forStmt.Body.List)
		} else if switchStmt, ok := stmt.(*ast.SwitchStmt); ok {
			count += countGoStmt(switchStmt.Body.List)
		} else if caseClause, ok := stmt.(*ast.CaseClause); ok {
			count += countGoStmt(caseClause.Body)
		}
	}
	return count
}
