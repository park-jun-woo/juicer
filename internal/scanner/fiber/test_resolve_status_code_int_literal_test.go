//ff:func feature=scan type=test control=sequence
//ff:what TestResolveStatusCode_IntLiteral 테스트
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
