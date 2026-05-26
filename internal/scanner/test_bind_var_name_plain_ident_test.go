//ff:func feature=scan type=test control=sequence
//ff:what TestBindVarName_PlainIdent 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestBindVarName_PlainIdent(t *testing.T) {
	expr := &ast.Ident{Name: "data"}
	got := bindVarName(expr)
	if got != "data" {
		t.Fatalf("got %q", got)
	}
}
