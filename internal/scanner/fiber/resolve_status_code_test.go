//ff:func feature=scan type=test control=selection
//ff:what resolveStatusCode — 상태 코드 해석 테스트
package fiber

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestResolveStatusCode_IntLiteral(t *testing.T) {
	lit := &ast.BasicLit{Kind: token.INT, Value: "201"}
	if got := resolveStatusCode(lit, nil); got != "201" {
		t.Fatalf("int literal: got %q", got)
	}
}

func TestResolveStatusCode_NilInfoUnknown(t *testing.T) {
	if got := resolveStatusCode(&ast.Ident{Name: "x"}, nil); got != "(unknown)" {
		t.Fatalf("nil info: got %q", got)
	}
}

func TestResolveStatusCode_ConstSelector(t *testing.T) {
	src := `package m
import "net/http"
func use(int) {}
func h() { use(http.StatusCreated) }
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0]
			}
		}
		return true
	})
	if got := resolveStatusCode(arg, info); got != "201" {
		t.Fatalf("http.StatusCreated -> %q, want 201", got)
	}
}

func TestResolveStatusCode_ConstIdent(t *testing.T) {
	src := `package m
const Created = 201
func use(int) {}
func h() { use(Created) }
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0]
			}
		}
		return true
	})
	if got := resolveStatusCode(arg, info); got != "201" {
		t.Fatalf("Created -> %q, want 201", got)
	}
}

func TestResolveStatusCode_ConstExprValue(t *testing.T) {
	// a constant arithmetic expression -> resolved via info.Types[expr].Value
	src := `package m
func use(int) {}
func h() { use(200 + 1) }
`
	file, info := typedExprs(t, src)
	var arg ast.Expr
	ast.Inspect(file, func(n ast.Node) bool {
		if c, ok := n.(*ast.CallExpr); ok {
			if id, ok := c.Fun.(*ast.Ident); ok && id.Name == "use" {
				arg = c.Args[0]
			}
		}
		return true
	})
	if got := resolveStatusCode(arg, info); got != "201" {
		t.Fatalf("200+1 -> %q, want 201", got)
	}
}

func TestResolveStatusCode_Unknown(t *testing.T) {
	// a non-const ident with empty info maps -> "(unknown)"
	if got := resolveStatusCode(&ast.Ident{Name: "dynamic"}, newEmptyInfoFull()); got != "(unknown)" {
		t.Fatalf("unknown: got %q", got)
	}
}
