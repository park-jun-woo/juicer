//ff:func feature=scan type=test control=sequence
//ff:what resolveGroupPrefixFile 전 분기 테스트
package gogin

import (
	"go/ast"
	"go/token"
	rgpfPars "go/parser"
	rgpfTyp "go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
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

func TestResolveGroupPrefixFile_WithGin(t *testing.T) {
	src := `package m
import "github.com/gin-gonic/gin"
func Setup(r *gin.Engine) { r.GET("/x", h) }
var X = 1
`
	fset := token.NewFileSet()
	file, _ := rgpfPars.ParseFile(fset, "/proj/m.go", src, 0)
	pkg := &packages.Package{Syntax: []*ast.File{file}, TypesInfo: &rgpfTyp.Info{}, Fset: fset}
	idx := &funcIndex{}
	ep := map[struct {
		file string
		line int
	}]int{}
	resolveGroupPrefixFile(file, pkg, []*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{}, ep)
}
