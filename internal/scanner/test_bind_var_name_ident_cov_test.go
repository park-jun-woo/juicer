//ff:func feature=scan type=test control=sequence
//ff:what TestBindVarName_IdentCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestBindVarName_IdentCov(t *testing.T) {
	expr := &ast.Ident{Name: "body"}
	got := bindVarName(expr)
	if got != "body" {
		t.Fatalf("got %q", got)
	}
}
