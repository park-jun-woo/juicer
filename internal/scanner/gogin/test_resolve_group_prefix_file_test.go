//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefixFile 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefixFile(t *testing.T) {
	fset := token.NewFileSet()

	file := &ast.File{
		Name:    &ast.Ident{Name: "main"},
		Imports: nil,
	}
	pkg := &packages.Package{Fset: fset}
	epIndex := map[struct {
		file string
		line int
	}]int{}
	resolveGroupPrefixFile(file, pkg, nil, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndex)
}
