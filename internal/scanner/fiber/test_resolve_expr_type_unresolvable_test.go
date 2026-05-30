//ff:func feature=scan type=test control=sequence
//ff:what TestResolveExprType_Unresolvable 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveExprType_Unresolvable(t *testing.T) {

	tn, f := resolveExprType(&ast.Ident{Name: "ghost"}, newEmptyInfoFull())
	if tn != "" || f != nil {
		t.Fatalf("unresolvable: %q %v", tn, f)
	}
}
