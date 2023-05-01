package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Get basic metrics for source files
// lines of code in a file
// if public func is commented
// If public struct is commented

type FuncMetrics struct {
	Function    string // name of function decl
	LOC         int    // how many lines of code for this function
	Public      bool   // Is function decl public
	HasComments bool   // is function commented with // doc before func declaration
}

type LoCMetrics struct {
	Functions                 []FuncMetrics //slice of all Funcetric structs for a file
	CommentCoverageAsAPercent float32
}

type StructMetrics struct {
	Struct      string //name of struct
	HasComments bool   // is function commented with // doc before func declaration
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
		if _, ok := n.(*ast.StructType); ok {
			if !ok {
				return true
			}
		}
		return true
	})

	var noComments float32
	for _, v := range lm.Functions {
		if v.HasComments == false {
			noComments += 1
		}
	}
	lm.CommentCoverageAsAPercent = calculatePercentage(
		float32(noComments), float32(len(lm.Functions)),
	)
}

// Calculate what percentage argument a is of argument b
func calculatePercentage(a, b float32) float32 {
	return (a / b) * 100
}
