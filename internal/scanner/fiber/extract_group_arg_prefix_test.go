//ff:func feature=scan type=test control=sequence
//ff:what extractGroupArgPrefix — 그룹 인자 prefix 추출 테스트
package fiber

import (
	"go/ast"
	"go/parser"
	"testing"
)

func groupCtx() *groupArgCtx {
	return &groupArgCtx{
		routers: map[string]*routerInfo{
			"api":       {prefix: "/api"},
			"authGroup": {prefix: "/auth"},
		},
	}
}

func TestExtractGroupArgPrefix_Variable(t *testing.T) {
	arg := ast.NewIdent("authGroup")
	prefix, ri, ok := extractGroupArgPrefix(arg, groupCtx())
	if !ok || prefix != "/auth" || ri == nil {
		t.Fatalf("variable group: prefix=%q ok=%v", prefix, ok)
	}
}

func TestExtractGroupArgPrefix_GroupCall(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Group("/v1")`)
	prefix, ri, ok := extractGroupArgPrefix(e, groupCtx())
	if !ok || prefix != "/api/v1" || ri == nil {
		t.Fatalf("group call: prefix=%q ok=%v", prefix, ok)
	}
}

func TestExtractGroupArgPrefix_GroupCallNoArgs(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Group()`)
	prefix, _, ok := extractGroupArgPrefix(e, groupCtx())
	if !ok || prefix != "/api" {
		t.Fatalf("group call no args: prefix=%q ok=%v", prefix, ok)
	}
}

func TestExtractGroupArgPrefix_NotCall(t *testing.T) {
	// unknown ident, not a call -> false
	e, _ := parser.ParseExpr(`42`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("basic lit should be false")
	}
}

func TestExtractGroupArgPrefix_NotGroupSelector(t *testing.T) {
	e, _ := parser.ParseExpr(`api.Use(mw)`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("non-Group selector should be false")
	}
}

func TestExtractGroupArgPrefix_UnknownRecv(t *testing.T) {
	e, _ := parser.ParseExpr(`unknown.Group("/x")`)
	_, _, ok := extractGroupArgPrefix(e, groupCtx())
	if ok {
		t.Fatal("unknown receiver should be false")
	}
}
