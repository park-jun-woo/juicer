//ff:func feature=scan type=test control=sequence
//ff:what identName — Ident 이름 추출 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestIdentName(t *testing.T) {
	if got := identName(ast.NewIdent("app")); got != "app" {
		t.Errorf("ident: got %q", got)
	}
	// non-ident -> ""
	if got := identName(&ast.CallExpr{}); got != "" {
		t.Errorf("non-ident: got %q, want empty", got)
	}
}
