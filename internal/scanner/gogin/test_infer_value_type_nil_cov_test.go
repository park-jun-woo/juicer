//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_NilCov 테스트
package gogin

import (
	"go/ast"
	"testing"
)

func TestInferValueType_NilCov(t *testing.T) {
	if got := inferValueType(&ast.Ident{Name: "nil"}, nil); got != "null" {
		t.Fatalf("expected null, got %s", got)
	}
}
