//ff:func feature=sql type=test control=sequence
//ff:what TestTypeString_StarCov 테스트
package sqls

import (
	"go/ast"
	"testing"
)

func TestTypeString_StarCov(t *testing.T) {
	if typeString(&ast.StarExpr{X: &ast.Ident{Name: "T"}}) != "*T" {
		t.Fatal("expected *T")
	}
}
