package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// Get basic metrics for source files
// lines of code in a file
// if func is commented

type FuncMetrics struct {
	Function    string
	LOC         int
	HasComments bool
}

type LoCMetrics struct {
	Functions []FuncMetrics
}

func main() {
	sourceFile := "test_source_file.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, sourceFile, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	var lm LoCMetrics
	ast.Inspect(f, func(n ast.Node) bool {
		if fd, ok := n.(*ast.FuncDecl); ok {
			if !ok {
				return true
			}
			var Commented bool
			if fd.Doc != nil {
				Commented = true
			}
			// fmt.Println(fd.Doc.Text())
			lm.Functions = append(lm.Functions, FuncMetrics{
				Function:    fd.Name.Name,
				LOC:         fset.Position(fd.Body.Rbrace).Line - fset.Position(fd.Body.Lbrace).Line + 1,
				HasComments: Commented,
			},
			)
		}
		return true
	})

	for _, v := range lm.Functions {
		fmt.Println(v.Function, v.HasComments, v.LOC)
	}
}
