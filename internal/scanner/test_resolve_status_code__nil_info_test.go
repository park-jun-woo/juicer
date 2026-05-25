//ff:func feature=scan type=extract control=sequence
//ff:what TestResolveStatusCode_NilInfo 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_NilInfo(t *testing.T) {
	expr := &ast.Ident{Name: "StatusOK"}
	got := resolveStatusCode(expr, nil)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}
