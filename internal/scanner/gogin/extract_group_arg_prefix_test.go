//ff:func feature=scan type=test control=sequence
//ff:what extractGroupArgPrefix 전 분기 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExtractGroupArgPrefix(t *testing.T) {
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{
			"r": {prefix: "/api"},
		},
	}

	// non-call expr
	_, _, ok := extractGroupArgPrefix(&ast.Ident{Name: "x"}, ctx)
	if ok {
		t.Fatal("expected false for non-call")
	}

	// call but not Group selector
	call := &ast.CallExpr{Fun: &ast.Ident{Name: "foo"}}
	_, _, ok = extractGroupArgPrefix(call, ctx)
	if ok {
		t.Fatal("expected false for non-selector call")
	}

	// Group call but unknown receiver
	groupCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "unknown"},
			Sel: &ast.Ident{Name: "Group"},
		},
	}
	_, _, ok = extractGroupArgPrefix(groupCall, ctx)
	if ok {
		t.Fatal("expected false for unknown receiver")
	}

	// Group call with known receiver and string arg
	groupCallOK := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "r"},
			Sel: &ast.Ident{Name: "Group"},
		},
		Args: []ast.Expr{&ast.BasicLit{Kind: 9, Value: `"/v1"`}},
	}
	prefix, parent, ok := extractGroupArgPrefix(groupCallOK, ctx)
	if !ok {
		t.Fatal("expected true for valid Group call")
	}
	_ = prefix
	_ = parent
}
