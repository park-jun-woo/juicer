//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_SelectorCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_SelectorCov(t *testing.T) {
	if typeString(&ast.SelectorExpr{X: &ast.Ident{Name: "context"}, Sel: &ast.Ident{Name: "Context"}}) != "context.Context" {
		t.Fatal("expected context.Context")
	}
}
