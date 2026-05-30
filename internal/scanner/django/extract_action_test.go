//ff:func feature=scan type=test control=sequence topic=django
//ff:what extractActions — 클래스 body의 @action 메서드 추출 분기를 검증
package django

import (
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
)

func firstBlock(root *sitter.Node) *sitter.Node {
	var found *sitter.Node
	var walk func(n *sitter.Node)
	walk = func(n *sitter.Node) {
		if found != nil {
			return
		}
		if n.Type() == "block" {
			found = n
			return
		}
		for i := 0; i < int(n.ChildCount()); i++ {
			walk(n.Child(i))
		}
	}
	walk(root)
	return found
}

func TestExtractActions(t *testing.T) {
	src := `
class UserViewSet(ModelViewSet):
    @action(detail=True, methods=['post'])
    def set_password(self, request, pk=None):
        return Response()

    def plain(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if body == nil {
		t.Fatal("no class body block")
	}
	actions := extractActions(body, []byte(src), "views.py")
	if len(actions) != 1 {
		t.Fatalf("expected 1 @action method, got %d", len(actions))
	}
	if actions[0].name != "set_password" {
		t.Errorf("name = %q, want set_password", actions[0].name)
	}
}

func TestExtractActions_None(t *testing.T) {
	src := `
class C:
    def plain(self):
        pass
`
	root, err := parsePython([]byte(src))
	if err != nil {
		t.Fatal(err)
	}
	body := firstBlock(root)
	if a := extractActions(body, []byte(src), "v.py"); len(a) != 0 {
		t.Fatalf("expected no actions, got %d", len(a))
	}
}
