//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_NilInfoCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestResolveStatusCode_NilInfoCov(t *testing.T) {
	expr := &ast.Ident{Name: "code"}
	got := resolveStatusCode(expr, nil)
	if got != "(unknown)" {
		t.Fatalf("expected (unknown), got %s", got)
	}
}
