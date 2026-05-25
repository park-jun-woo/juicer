//ff:func feature=scan type=extract control=sequence
//ff:what TestInferValueType_Unknown 테스트
package scanner

import (
	"go/ast"
	"go/types"
	"testing"
)

func TestInferValueType_Unknown(t *testing.T) {
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	got := inferValueType(&ast.Ident{Name: "x"}, info)
	if got != "unknown" {
		t.Fatalf("expected unknown, got %s", got)
	}
}
