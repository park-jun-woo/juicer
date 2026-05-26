//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_FloatCov 테스트
package gogin

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestInferValueType_FloatCov(t *testing.T) {
	if got := inferValueType(&ast.BasicLit{Kind: token.FLOAT, Value: "3.14"}, nil); got != "number" {
		t.Fatalf("expected number, got %s", got)
	}
}
