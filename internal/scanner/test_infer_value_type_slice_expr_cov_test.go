//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_SliceExprCov 테스트
package scanner

import (
	"go/ast"
	"testing"
)

func TestInferValueType_SliceExprCov(t *testing.T) {
	if got := inferValueType(&ast.SliceExpr{X: &ast.Ident{Name: "s"}}, nil); got != "array" {
		t.Fatalf("expected array, got %s", got)
	}
}
