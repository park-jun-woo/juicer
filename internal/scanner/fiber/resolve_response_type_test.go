//ff:func feature=scan type=test control=sequence
//ff:what resolveResponseType — 응답 타입 추적 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveResponseType_NilInfo(t *testing.T) {
	tn, f, c := resolveResponseType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || f != nil || c != "" {
		t.Fatalf("nil info: %q %v %q", tn, f, c)
	}
}

func TestResolveResponseType_FullFields(t *testing.T) {
	src := `package m
type Resp struct { OK bool ` + "`json:\"ok\"`" + ` }
func use(interface{}) {}
func h() { var r Resp; use(r) }
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
	tn, fields, conf := resolveResponseType(arg, info)
	if tn != "Resp" || len(fields) != 1 || conf != "full" {
		t.Fatalf("full: %q %v %q", tn, fields, conf)
	}
}

func TestResolveResponseType_NoFields(t *testing.T) {
	// an int arg -> resolveExprType yields a type name but no fields
	src := `package m
func use(interface{}) {}
func h() { var n int; use(n) }
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
	_, fields, conf := resolveResponseType(arg, info)
	if len(fields) != 0 || conf != "" {
		t.Fatalf("no fields: %v %q", fields, conf)
	}
}
