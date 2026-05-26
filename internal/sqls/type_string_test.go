//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_Ident 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_Ident(t *testing.T) {
	got := typeString(&ast.Ident{Name: "string"})
	if got != "string" {
		t.Fatalf("expected string, got %s", got)
	}
}
