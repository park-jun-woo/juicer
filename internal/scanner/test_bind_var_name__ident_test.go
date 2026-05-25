//ff:func feature=scan type=extract control=sequence
//ff:what TestBindVarName_Ident 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestBindVarName_Ident(t *testing.T) {
	expr := &ast.Ident{Name: "req"}
	got := bindVarName(expr)
	if got != "req" {
		t.Fatalf("got %q", got)
	}
}
