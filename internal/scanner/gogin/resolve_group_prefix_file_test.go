//ff:func feature=scan type=test control=sequence
//ff:what resolveGroupPrefixFile 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/park-jun-woo/juicer/internal/scanner"
	"golang.org/x/tools/go/packages"
)

func TestResolveGroupPrefixFile(t *testing.T) {
	fset := token.NewFileSet()
	// file with no gin import -> early return
	file := &ast.File{
		Name:    &ast.Ident{Name: "main"},
		Imports: nil,
	}
	pkg := &packages.Package{Fset: fset}
	epIndex := map[struct{ file string; line int }]int{}
	resolveGroupPrefixFile(file, pkg, nil, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{}, epIndex)
}
