//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_NilInfo 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_NilInfo(t *testing.T) {
	tn, f := resolveExprType(&ast.Ident{Name: "x"}, nil)
	if tn != "" || f != nil {
		t.Fatalf("nil info: %q %v", tn, f)
	}
}
