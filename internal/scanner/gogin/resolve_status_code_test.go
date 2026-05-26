//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_IntLit 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestResolveStatusCode_IntLit(t *testing.T) {
	expr := &ast.BasicLit{Kind: token.INT, Value: "200"}
	got := resolveStatusCode(expr, nil)
	if got != "200" {
		t.Fatalf("expected 200, got %s", got)
	}
}

