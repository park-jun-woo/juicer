//ff:func feature=scan type=test control=sequence
//ff:what extractGroupArgPrefix에서 *ast.Ident 인자가 routers 맵에 있으면 prefix와 middleware를 반환하는지 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestExtractGroupArgPrefix_Ident(t *testing.T) {
	ctx := &groupArgCtx{
		routers: map[string]*routerInfo{
			"authGroup": {prefix: "/api", middleware: []string{"middleware.Auth()"}},
		},
	}

	// known ident
	prefix, ri, ok := extractGroupArgPrefix(&ast.Ident{Name: "authGroup"}, ctx)
	if !ok {
		t.Fatal("expected true for known ident")
	}
	if prefix != "/api" {
		t.Errorf("expected prefix /api, got %q", prefix)
	}
	if len(ri.middleware) != 1 || ri.middleware[0] != "middleware.Auth()" {
		t.Errorf("expected middleware [middleware.Auth()], got %v", ri.middleware)
	}

	// unknown ident falls through to CallExpr check
	_, _, ok = extractGroupArgPrefix(&ast.Ident{Name: "unknown"}, ctx)
	if ok {
		t.Fatal("expected false for unknown ident")
	}
}
