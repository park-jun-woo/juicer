//ff:func feature=scan type=test control=sequence topic=express
//ff:what TestBuildUnresolvedSet_Round5 테스트
package express

import (
	sitter "github.com/smacker/go-tree-sitter"
	"testing"
)

func TestBuildUnresolvedSet_Round5(t *testing.T) {
	schemas := map[string]*sitter.Node{"A": nil}
	set := buildUnresolvedSet([]string{"A", "B", "C"}, schemas)
	if set["A"] {
		t.Error("A is resolved, should not be in set")
	}
	if !set["B"] || !set["C"] {
		t.Errorf("B,C should be unresolved: %v", set)
	}
}
