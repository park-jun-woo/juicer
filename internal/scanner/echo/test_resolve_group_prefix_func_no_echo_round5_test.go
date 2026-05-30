//ff:func feature=scan type=test control=sequence topic=echo
//ff:what TestResolveGroupPrefixFunc_NoEcho_Round5 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefixFunc_NoEcho_Round5(t *testing.T) {
	src := `package m
func setup() {}
`
	fset := token.NewFileSet()
	file, _ := parser.ParseFile(fset, "m.go", src, parser.SkipObjectResolution)
	fn := file.Decls[0].(*ast.FuncDecl)

	resolveGroupPrefixFunc(fn, "", &packages.Package{Fset: fset}, nil, "/root", buildFuncIndex(nil), nil, map[int][]ast.Expr{}, map[struct {
		file string
		line int
	}]int{})
}
