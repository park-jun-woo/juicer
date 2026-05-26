//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_IntCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestInferValueType_IntCov(t *testing.T) {
	if got := inferValueType(&ast.BasicLit{Kind: token.INT, Value: "42"}, nil); got != "integer" {
		t.Fatalf("expected integer, got %s", got)
	}
}
