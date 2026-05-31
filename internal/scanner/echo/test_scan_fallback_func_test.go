//ff:func feature=scan type=test topic=echo control=sequence
//ff:what scanFallbackFunc echo.Context 핸들러 본문에서 응답 수집 및 비-핸들러 무시 테스트
package echo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	"github.com/park-jun-woo/codistill/internal/scanner"
)

func TestScanFallbackFunc(t *testing.T) {
	src := `package p
import "github.com/labstack/echo/v4"
func H(c echo.Context) error { return c.JSON(200, map[string]any{"x": 1}) }
func NotHandler(x int) {}`
	f, err := parser.ParseFile(token.NewFileSet(), "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	idx := &funcIndex{
		byName:     map[string]*ast.FuncDecl{},
		byPos:      map[token.Pos]*ast.FuncDecl{},
		astStructs: map[string]*ast.StructType{},
	}
	indexFileDecls(f, nil, idx)

	h := idx.byName["H"]
	ep := &scanner.Endpoint{}
	scanFallbackFunc(ep, h.Type, h.Body, idx)
	if len(ep.Responses) == 0 {
		t.Errorf("handler should produce responses")
	}

	// non-echo-context function -> no ctx name -> no-op
	nh := idx.byName["NotHandler"]
	ep2 := &scanner.Endpoint{}
	scanFallbackFunc(ep2, nh.Type, nh.Body, idx)
	if len(ep2.Responses) != 0 {
		t.Errorf("non-handler should be no-op: %+v", ep2.Responses)
	}
}
