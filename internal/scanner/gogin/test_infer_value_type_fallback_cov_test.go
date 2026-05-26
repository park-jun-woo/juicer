//ff:func feature=scan type=test control=sequence
//ff:what TestInferValueType_FallbackCov 테스트
package gogin

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_FallbackCov(t *testing.T) {
	info := &types.Info{Types: map[ast.Expr]types.TypeAndValue{}}
	if got := inferValueType(&ast.Ident{Name: "x"}, info); got != "unknown" {
		t.Fatalf("expected unknown, got %s", got)
	}
}
