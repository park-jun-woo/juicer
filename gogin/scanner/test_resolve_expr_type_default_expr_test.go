//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveExprType_DefaultExpr 테스트
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

func TestResolveExprType_DefaultExpr(t *testing.T) {
	// Use an expr type that falls into default case
	src := `package test

func f() int {
	return 1 + 2
}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "test.go", src, 0)
	conf := types.Config{}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}
	conf.Check("test", fset, []*ast.File{file}, info)

	// Find the binary expr (1 + 2)
	var binExpr *ast.BinaryExpr
	ast.Inspect(file, func(n ast.Node) bool {
		if be, ok := n.(*ast.BinaryExpr); ok {
			binExpr = be
			return false
		}
		return true
	})

	if binExpr != nil {
		typeName, _ := resolveExprType(binExpr, info)
		_ = typeName
	}
}
