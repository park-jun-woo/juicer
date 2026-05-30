//ff:func feature=scan type=test control=sequence
//ff:what TestResolveGroupPrefix_ProcessesFile 테스트
package gogin

import (
	"github.com/park-jun-woo/codistill/internal/scanner"
	"go/ast"
	rgpPars "go/parser"
	rgpTok "go/token"
	rgpTyp "go/types"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestResolveGroupPrefix_ProcessesFile(t *testing.T) {
	src := `package m
import "github.com/gin-gonic/gin"
func Setup(r *gin.Engine) { r.GET("/x", h) }
`
	fset := rgpTok.NewFileSet()
	file, _ := rgpPars.ParseFile(fset, "/proj/m.go", src, 0)
	pkg := &packages.Package{
		Syntax:          []*ast.File{file},
		TypesInfo:       &rgpTyp.Info{},
		CompiledGoFiles: []string{"/proj/m.go"},
		Fset:            fset,
	}
	idx := &funcIndex{}
	resolveGroupPrefix([]*packages.Package{pkg}, "/proj", idx, []scanner.Endpoint{}, map[int][]ast.Expr{})
}
