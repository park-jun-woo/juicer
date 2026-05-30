//ff:func feature=scan type=test control=sequence topic=fastapi
//ff:what extractDependsFromList: Depends 수집 / 비Depends 스킵
package fastapi

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstList(t *testing.T, src []byte) (*sitter.Node, []byte) {
	t.Helper()
	root, err := parsePython(src)
	if err != nil {
		t.Fatal(err)
	}
	lists := findAllByType(root, "list")
	if len(lists) == 0 {
		t.Fatal("no list")
	}
	return lists[0], src
}

func TestExtractDependsFromList_Collects(t *testing.T) {
	list, src := firstList(t, []byte("x = [Depends(auth), Depends(log), other()]\n"))
	deps := extractDependsFromList(list, src)
	if len(deps) != 2 || deps[0] != "auth" || deps[1] != "log" {
		t.Fatalf("got %v", deps)
	}
}

func TestExtractDependsFromList_None(t *testing.T) {
	list, src := firstList(t, []byte("x = [a, b, c]\n"))
	deps := extractDependsFromList(list, src)
	if len(deps) != 0 {
		t.Fatalf("expected none, got %v", deps)
	}
}
