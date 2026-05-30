//ff:func feature=scan type=test control=sequence
//ff:what TestResolveUsesConst_NotInUses 테스트
package fiber

import (
	"go/ast"
	"testing"
)

func TestResolveUsesConst_NotInUses(t *testing.T) {

	if got := resolveUsesConst(newEmptyInfoFull(), ast.NewIdent("ghost")); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}
