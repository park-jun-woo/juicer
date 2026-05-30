//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefixFile_WithGin 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	rgpfPars "go/parser"
	"go/token"
	rgpfTyp "go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

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
