//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_SliceExprCase 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestInferValueType_SliceExprCase(t *testing.T) {
	got := inferValueType(&ast.SliceExpr{X: &ast.Ident{Name: "arr"}}, nil)
	if got != "array" {
		t.Fatalf("expected array, got %s", got)
	}
}
