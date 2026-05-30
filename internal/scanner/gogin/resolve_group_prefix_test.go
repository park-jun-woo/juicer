//ff:func feature=scan type=test control=sequence
//ff:what resolveGroupPrefix 전 분기 테스트
package gogin

import (
	"go/ast"
	rgpTok "go/token"
	rgpPars "go/parser"
	rgpTyp "go/types"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
	"golang.org/x/tools/go/packages"
)

func TestResolveGroupPrefix(t *testing.T) {
	// empty pkgs
	resolveGroupPrefix(nil, "/tmp", &funcIndex{}, nil, nil)

	// pkg with nil TypesInfo
	pkg := &packages.Package{}
	resolveGroupPrefix([]*packages.Package{pkg}, "/tmp", &funcIndex{}, nil, nil)

	// pkg with TypesInfo but no syntax
	pkg2 := &packages.Package{
		TypesInfo: nil,
	}
	resolveGroupPrefix([]*packages.Package{pkg2}, "/tmp", &funcIndex{}, []scanner.Endpoint{}, map[int][]ast.Expr{})
}

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
